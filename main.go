package main

import (
    "context"
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

    "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/schema"

    "k8s.io/apimachinery/pkg/util/wait"
    "k8s.io/apimachinery/pkg/watch"
    "k8s.io/client-go/dynamic"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/util/workqueue"
    "k8s.io/klog/v2"
)

const (
    controllerName = "opal-controller"
)

type OpalDeployment struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec              OpalDeploymentSpec   `json:"spec,omitempty"`
    Status            OpalDeploymentStatus `json:"status,omitempty"`
}

type OpalDeploymentSpec struct {
    Replicas int32 `json:"replicas"`
}

type OpalDeploymentStatus struct {
    AvailableReplicas int32 `json:"availableReplicas"`
}

type Controller struct {
    kubeclientset kubernetes.Interface
    dynamicclient dynamic.Interface
    queue         workqueue.RateLimitingInterface
    informer      cache.SharedIndexInformer
}

func NewController(
    kubeclientset kubernetes.Interface,
    dynamicclient dynamic.Interface,
    informer cache.SharedIndexInformer) *Controller {

    queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

    informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(obj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(obj)
            if err == nil {
                queue.Add(key)
            }
        },
        UpdateFunc: func(old, new interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(new)
            if err == nil {
                queue.Add(key)
            }
        },
        DeleteFunc: func(obj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(obj)
            if err == nil {
                queue.Add(key)
            }
        },
    })

    return &Controller{
        kubeclientset: kubeclientset,
        dynamicclient: dynamicclient,
        informer:      informer,
        queue:         queue,
    }
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
    defer runtime.HandleCrash()
    defer c.queue.ShutDown()

    klog.Info("Starting Controller")

    go c.informer.Run(stopCh)

    if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
        return fmt.Errorf("Timed out waiting for caches to sync")
    }

    for i := 0; i < threadiness; i++ {
        go wait.Until(c.runWorker, time.Second, stopCh)
    }

    <-stopCh
    klog.Info("Stopping Controller")

    return nil
}

func (c *Controller) runWorker() {
    for c.processNextItem() {
    }
}

func (c *Controller) processNextItem() bool {
    key, quit := c.queue.Get()
    if quit {
        return false
    }
    defer c.queue.Done(key)

    err := c.syncHandler(key.(string))
    if err == nil {
        c.queue.Forget(key)
    } else {
        runtime.HandleError(fmt.Errorf("Error syncing '%s': %s", key, err.Error()))
        c.queue.AddRateLimited(key)
    }

    return true
}

func (c *Controller) syncHandler(key string) error {
    namespace, name, err := cache.SplitMetaNamespaceKey(key)
    if err != nil {
        runtime.HandleError(fmt.Errorf("Invalid resource key: %s", key))
        return nil
    }

    // Get the resource with this namespace/name
    gvr := schema.GroupVersionResource{Group: "example.com", Version: "v1alpha1", Resource: "opaldeployments"}
    opalDeployment, err := c.dynamicclient.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
    if err != nil {
        if errors.IsNotFound(err) {
            runtime.HandleError(fmt.Errorf("opalDeployment '%s' in work queue no longer exists", key))
            return nil
        }
        return err
    }

    // Reconcile the resource
    err = c.reconcileOpalDeployment(opalDeployment)
    if err != nil {
        return err
    }

    return nil
}

func (c *Controller) reconcileOpalDeployment(opalDeployment *unstructured.Unstructured) error {
    // Example: Check if the deployment exists, if not create it
    deploymentName := opalDeployment.GetName() + "-opal"
    namespace := opalDeployment.GetNamespace()

    gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
    _, err := c.dynamicclient.Resource(gvr).Namespace(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
    if errors.IsNotFound(err) {
        klog.Infof("Creating deployment %s/%s", namespace, deploymentName)
        deployment := &unstructured.Unstructured{
            Object: map[string]interface{}{
                "apiVersion": "apps/v1",
                "kind":       "Deployment",
                "metadata": map[string]interface{}{
                    "name":      deploymentName,
                    "namespace": namespace,
                },
                "spec": map[string]interface{}{
                    "replicas": opalDeployment.Object["spec"].(map[string]interface{})["replicas"],
                    "selector": map[string]interface{}{
                        "matchLabels": map[string]interface{}{
                            "app": deploymentName,
                        },
                    },
                    "template": map[string]interface{}{
                        "metadata": map[string]interface{}{
                            "labels": map[string]interface{}{
                                "app": deploymentName,
                            },
                        },
                        "spec": map[string]interface{}{
                            "containers": []map[string]interface{}{
                                {
                                    "name":  "opal",
                                    "image": "openpolicyagent/opa:latest",
                                    "ports": []map[string]interface{}{
                                        {
                                            "containerPort": 8181,
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        }

        _, err = c.dynamicclient.Resource(gvr).Namespace(namespace).Create(context.Background(), deployment, metav1.CreateOptions{})
        if err != nil {
            return fmt.Errorf("failed to create deployment: %v", err)
        }
    } else if err != nil {
        return fmt.Errorf("failed to get deployment: %v", err)
    }

    // Update the status of the OpalDeployment
    opalDeployment.Object["status"] = map[string]interface{}{
        "availableReplicas": 1,
    }
    _, err = c.dynamicclient.Resource(gvr).Namespace(namespace).UpdateStatus(context.Background(), opalDeployment, metav1.UpdateOptions{})
    if err != nil {
        return fmt.Errorf("failed to update opalDeployment status: %v", err)
    }

    return nil
}

func main() {
    klog.InitFlags(nil)
    flag.Set("logtostderr", "true")
    flag.Parse()

    config, err := rest.InClusterConfig()
    if err != nil {
        klog.Fatalf("Error building kubeconfig: %s", err.Error())
    }

    kubeclientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
    }

    dynamicclient, err := dynamic.NewForConfig(config)
    if err != nil {
        klog.Fatalf("Error building dynamic client: %s", err.Error())
    }

    gvr := schema.GroupVersionResource{Group: "example.com", Version: "v1alpha1", Resource: "opaldeployments"}
    informer := cache.NewSharedIndexInformer(
        &cache.ListWatch{
            ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
                return dynamicclient.Resource(gvr).Namespace(metav1.NamespaceAll).List(context.Background(), options)
            },
            WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
                return dynamicclient.Resource(gvr).Namespace(metav1.NamespaceAll).Watch(context.Background(), options)
            },
        },
        &unstructured.Unstructured{},
        0,
        cache.Indexers{},
    )

    controller := NewController(kubeclientset, dynamicclient, informer)

    stopCh := make(chan struct{})
    defer close(stopCh)

    go func() {
        sigCh := make(chan os.Signal, 1)
        signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
        <-sigCh
        close(stopCh)
    }()

    if err = controller.Run(2, stopCh); err != nil {
        klog.Fatalf("Error running controller: %s", err.Error())
    }
}