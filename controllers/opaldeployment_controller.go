/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
    "context"
    "fmt"

    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/apimachinery/pkg/api/meta"
    "k8s.io/apimachinery/pkg/runtime"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/controller"
    "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
    "sigs.k8s.io/controller-runtime/pkg/log"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

    opalv1alpha1 "github.com/example/opal-operator/api/v1alpha1"
)

// OPALDeploymentReconciler reconciles a OPALDeployment object
type OPALDeploymentReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=opal.example.com,resources=opaldeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=opal.example.com,resources=opaldeployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=opal.example.com,resources=opaldeployments/finalizers,verbs=update

func (r *OPALDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    _ = log.FromContext(ctx)

    // Fetch the OPALDeployment instance
    instance := &opalv1alpha1.OPALDeployment{}
    err := r.Get(ctx, req.NamespacedName, instance)
    if err != nil {
        if errors.IsNotFound(err) {
            return ctrl.Result{}, nil
        }
        return ctrl.Result{}, err
    }

    // Define the desired Deployment object
    dep := r.deploymentForOPAL(instance)

    // Set OPALDeployment instance as the owner and controller
    if err := controllerutil.SetControllerReference(instance, dep, r.Scheme); err != nil {
        return ctrl.Result{}, err
    }

    // Check if the Deployment already exists
    found := &appsv1.Deployment{}
    err = r.Get(ctx, req.NamespacedName, found)
    if err != nil && errors.IsNotFound(err) {
        log.Log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
        err = r.Create(ctx, dep)
        if err != nil {
            return ctrl.Result{}, err
        }
        return ctrl.Result{}, nil
    } else if err != nil {
        return ctrl.Result{}, err
    }

    // Update the Deployment if necessary
    if !meta.EqualObjectMeta(instance.ObjectMeta, found.ObjectMeta) {
        log.Log.Info("Updating Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
        err = r.Update(ctx, dep)
        if err != nil {
            return ctrl.Result{}, err
        }
    }

    return ctrl.Result{}, nil
}

// deploymentForOPAL returns a Deployment object
func (r *OPALDeploymentReconciler) deploymentForOPAL(m *opalv1alpha1.OPALDeployment) *appsv1.Deployment {
    labels := map[string]string{"app": "opal"}
    replicas := int32(m.Spec.Size)

    return &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name:      m.Name,
            Namespace: m.Namespace,
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: &replicas,
            Selector: &metav1.LabelSelector{
                MatchLabels: labels,
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: labels,
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{{
                        Name:  "opal",
                        Image: "authorizon/opal-server:latest",
                    }},
                },
            },
        },
    }
}

func (r *OPALDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&opalv1alpha1.OPALDeployment{}).
        Complete(r)
}
