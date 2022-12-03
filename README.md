# A simple http server in GO with SWAGGER 🚀

## How to run

```bash
go run main.go
```

## How to build

```bash
go build main.go
```

## What does it do?

This code is a simple Go program that defines a struct to represent a user and uses the gin library to define a simple REST API for managing users. The code defines a GET route at /users that returns a list of all users. The code also uses the swaggo library to add Swagger documentation to the API. When the program is run, it listens for HTTP requests on port 8080 and logs messages to the console.## How to use

### Swagger UI

The Swagger UI is served at the /swagger/ URL path prefix. For example, if the server is running on localhost, you can access the Swagger UI at http://localhost:8080/swagger/.

### API

The API is served at the / URL path prefix. For example, if the server is running on localhost, you can access the API at http://localhost:8080/.

## How to test

### Swagger UI

The Swagger UI provides an interactive documentation and testing tool for the API. You can use the Swagger UI to test the API by entering values for the parameters and clicking the "Try it out!" button.

### API

You can also test the API by sending HTTP requests directly to the API endpoints. For example, you can use curl to send a GET request to the /api/hello endpoint:

```bash

curl http://localhost:8080/users

```

## How to deploy

### Docker

To build a Docker image from this Dockerfile, you can run the following command in the same directory as the Dockerfile:

```bash
docker build -t hello-world-api:latest .
```

This command will build a new Docker image called hello-world-api with the latest tag. The image will contain your program, as well as the necessary dependencies and runtime environment to run it.

Once the image is built, you can run it as a Docker container using the docker run command:

```bash
docker run -p 8080:8080 hello-world-api:latest
```

This command will run a new container from the hello-world-api image, and bind the container's port 8080 to the host's port 8080. This will allow you to access the API from your host machine by sending HTTP requests to http://localhost:8080.

You can also use the docker-compose tool to manage the Docker container, along with any other dependencies or services that your program may need. For example, you can use docker-compose to run a database or message queue alongside your program in separate containers, and configure them to communicate with each other.

Overall, using Docker allows you to easily package and deploy your program as a standalone, self-contained image that can be run on any host that has Docker installed. This makes it easier to share and run your program, and ensures that it will always run in a consistent environment, regardless of the host system's configuration.

### Kubernetes

To deploy your program to Kubernetes, you can use the kubectl tool to create a new Kubernetes deployment. For example, you can run the following command to create a new deployment called hello-world-api:

```bash
kubectl create deployment hello-world-api --image=hello-world-api:latest
```

This command will create a new deployment called hello-world-api, which will run a single pod containing a container based on the hello-world-api image. The deployment will also create a new Kubernetes service called hello-world-api, which will expose the deployment's pods on port 8080.

Once the deployment is created, you can use the kubectl expose command to create a new Kubernetes service that exposes the deployment on port 8080:

```bash
kubectl expose deployment hello-world-api --type=LoadBalancer --port=8080
```

This command will create a new Kubernetes service called hello-world-api, which will expose the deployment's pods on port 8080. The service will also create a new load balancer that will forward incoming traffic on port 8080 to the service's pods.

Once the service is created, you can use the kubectl get command to retrieve the public IP address of the load balancer:

```bash
kubectl get service hello-world-api
```

This command will print the public IP address of the load balancer to the terminal. You can then use this IP address to access the API from your host machine by sending HTTP requests to http://<public-ip>:8080.

Overall, using Kubernetes allows you to easily deploy your program to a cluster of machines, and manage the program's lifecycle. This makes it easier to scale your program, and ensures that it will always run in a consistent environment, regardless of the host system's configuration.

## How to contribute

### Code of Conduct

This project adheres to the Contributor Covenant [code of conduct](https://www.contributor-covenant.org/version/1/4/code-of-conduct.html). By participating, you are expected to uphold this code. Please report unacceptable behavior to [raphael.mansuy@gmail.com].# go_rest_api
