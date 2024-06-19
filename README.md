<h1 align="center">
  RECIPE 🧑‍🍳<br>
  <small>where karma decides role</small>
</h1>


<> image here <>

<br>

## 🧑‍🍳 Recipe   
 **Recipe Web Application** Built a k8s controller that will allow the deployment of OPAL as part of the k8s platform. Developed a Recipe Webapp where users' karma and location affect their permissions to perform operations. Fetched the data dynamically from a MySQL server.

<br>

## 🪶 Features
- **User Karma and Location-Based Permissions** 📖:
    - Dynamic Permission Management: Users' permissions to perform various operations within the application are dynamically managed based on their karma and location. This ensures personalized and context-specific access control.
    - Karma System: Users accumulate karma points through their interactions and contributions within the app, which affects their permissions and access levels.
    - Location-Based Access: Users' locations are used to further refine their access permissions, allowing for location-specific content and operations.
- **Recipe Sharing Platform** ✒️:
    - User-Generated Content: Users can create, share, and browse recipes. This includes the ability to add detailed instructions, ingredients, and multimedia content such as images and videos.
    - Search and Filter Options: The app provides advanced search and filtering options to help users find recipes based on various criteria, including ingredients, cuisine, and user ratings.
- **Kubernetes Integration** 🚀:
    - OPAL Deployment as a Kubernetes Controller: The application includes a Kubernetes controller to manage the deployment of OPAL as part of the Kubernetes platform, ensuring scalability and reliability.
    - Containerized Services: The app's backend services are containerized using Docker and deployed on Kubernetes, providing a robust and scalable infrastructure.
- **Database Integration** 📰:
    - MySQL Backend: User data, recipes, and other relevant information are stored in a MySQL database. The app dynamically fetches data from the MySQL server as needed.
    - API Endpoints: The app provides RESTful API endpoints to interact with the MySQL database, enabling operations such as fetching user information and retrieving recipes based on permissions.
- **Authentication and Authorization** 🚆:
    - User Authentication: Secure user authentication mechanisms are implemented to ensure that only authorized users can access and interact with the application.
    - Role-Based Access Control (RBAC): The application supports role-based access control, where different user roles have varying levels of access and permissions.
- **Monitoring and Logging** 🚆:
    - Responsive Design: The application is designed to be fully responsive, ensuring a seamless experience across different devices and screen sizes.
    - User-Friendly Interface: Intuitive navigation and a user-friendly interface enhance the overall user experience, making it easy for users to interact with the app.

- **Next.js Frontend** 🚆:
    - React-Based User Interface: The frontend is built using Next.js, providing a fast and responsive user experience.
    - Server-Side Rendering (SSR): Next.js enables server-side rendering, which improves performance and SEO for the web application.
    - Real-Time Updates: The app fetches data dynamically and updates the user interface in real-time as users interact with the application.

#

## :books: Index

- [Demo](#movie_camera-Demo)
- [Screenshots](#screenshots)
- [Set Up](#outbox_tray-Set-up)
- [Contribute](#building_construction-Contribute)
- [Project Author](#people_holding_hands-Meet-the-Author)
- [Contact](#email-contact)

#

###  :movie_camera: Demo
- After a brief introduction, let's dive a little more inside the project.
- Here is the walk-through of **Recipe 🧑‍🍳**. If you want to witness a more a hd version, [click here]()

<> video here <>

<p align="center">Video Demonstration</p>

### Screenshots

<p align="center">
  <img src=""  />
  <p align="center"></p>
  <br>
  <p align="center">
  <img src=""  />
  <p align="center"></p>
  <br>
  <p align="center">
  <img src=""  />
  <p align="center"></p>
</p>

<br>


## Dependecies
1. **Next.js**
2. *React*
3. **Axios**
4. *Node.js*
5. **Express**
6. *MySQL*
7. **OPAL**
8. *Kubectl*
4. *Helm*
5. **Docker**

#

##  :outbox_tray: Installation Guide
- These are the steps required to install and run the Recipe 🧑‍🍳 project:


1. Clone the Repository: Open a terminal or command prompt and clone the Recipe 🧑‍🍳 repository from GitHub using the following command:

  ```bash
    git clone https://github.com/RS-labhub/OPAL-K8.git
  ```

2. Navigate to the Repository Directory: Change your current directory to the cloned Recipe 🧑‍🍳 repository:

  ```bash
    cd OPAL-K8
  ```
3. Kubernetes Controller to Deploy OPAL
    - Prerequisites
        - A running Kubernetes cluster (local or cloud-based).
        - Docker installed on your local machine.
        - kubectl configured to interact with your Kubernetes cluster.
        - Go programming language installed.

    - Step 1: Set Up a Kubernetes Cluster
        - You can use Minikube for local development or a cloud provider like GKE, EKS, or AKS for a production-ready cluster.

    - Step 2: Create the Kubernetes Controller
        - Define the Custom Resource Definitions (CRDs)
        - Create the Controller Code
    - Step 3: Build and Deploy the Controller
        1. Build the Controller Image
        ```docker
        docker build -t mycompany/opal-controller:latest .
        ``` 
        2. Push the Image to a Registry
        ```docker
        docker push mycompany/opal-controller:latest
        ```
        3. Deploy the Controller to the Cluster
        4. Apply the deployment
        ```kubectl
        kubectl apply -f opal_crd.yaml
        kubectl apply -f controller_deployment.yaml
        ```
4. Web Application with Dynamic Permissions
    - Prerequisites
        - Node.js and npm installed.
        - MySQL server running and configured.
    - Step 1: Set Up a Next.js Project
        1. Install Required Dependencies
        ```
        npm install axios mysql2
        ```
        2. Run the project
        ```sh
        npm run dev
        ```
<br>

> Note: The running port should look like this:
Open `http://localhost:3000` in your web browser to see the list of users and recipes filtered by karma and location.

<br>
<br>


## What to do after installation?
Once Recipe 🧑‍🍳 is installed and running, by following these guidelines, you will set up and install a Kubernetes controller to deploy OPAL and develop a web application where users' permissions are dynamically influenced by their karma and location, with data fetched from a MySQL server.


$~$

## Setup and Contributing Guidelines
    
**Set Up Your Environment**

1. `Fork` our repository to your GitHub account. 
2. `Clone` your fork to your local machine. 
    Use the command `git clone https://github.com/RS-labhub/OPAL-K8.git`.
3. Create a new branch for your work. 
    Use a descriptive name, like `fix-login-bug` or `add-user-profile-page`.
    
**Commit Your Changes**

- Commit your changes with a _clear commit message_. 
  e.g `git commit -m "Fix login bug by updating auth logic"`.

**Submit a Pull Request**

- Push your branch and changes to your fork on GitHub.
- Create a pull request, compare branches and submit.
- Provide a detailed description of what changes you've made and why. 
  Link the pull request to the issue it resolves. 🔗
    
**Review and Merge**

- I will review your pull request and provide feedback or request changes if necessary. 
- Once your pull request is approved, we will merge it into the main codebase 🥳

$~$

## :people_holding_hands: Meet the Author

<img  src="Preview/author.jpeg" alt="Author">


### :email: Contact 
- Email: rs4101976@gmail.com
- Head over to my github handle from [here](https://github.com/RS-labhub)

<br>

<p align="center">
<a href="https://twitter.com/rrs00179" target="blank"><img src="https://img.shields.io/badge/Twitter/X-000000?style=for-the-badge&logo=x&logoColor=white" alt="rrs00179" /></a>
<a href="https://www.linkedin.com/in/rohan-sharma-9386rs/" target="blank"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="rohan-sharma=9386" /></a>
</p>

<br>

<p align="center">
   Thank you for visting this Repo <br>If you like it, <a href="https://github.com/RS-labhub/I-Love-You/stargazers">star</a> ⭐ it
</p>
