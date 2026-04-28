# Order Service

A robust microservice for managing orders, built with Go, Gin, and MongoDB. This service operates independently to handle order creation and retrieval within the application ecosystem.

## Features

- **Health Monitoring**: Simple endpoint to quickly verify the service is up and running.
- **Order Management**: Retrieve the list of all created orders or place new orders with specified items and quantities.

## Technologies Used

- **Go**: Version 1.25.0
- **Gin**: High-performance HTTP web framework
- **MongoDB**: NoSQL database for flexible and fast data storage
- **Docker**: Containerization to streamline deployments
- **Kubernetes**: For container orchestration scaling and management (Configured via `k8s.yaml`)

## Prerequisites

To run this application locally, you will need:
- [Go 1.25.0+](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/) (Optional, for containerized run)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) or any Kubernetes cluster (Optional, for local K8s deployment)

## Environment Variables

This service relies on a `.env` file for configuration. Create one in the root of the project with the following contents:

```env
PORT=4000
SERVICE_NAME=order-service
# Add your MongoDB connection string if required by the models/services:
# MONGO_URI=mongodb://localhost:27017
```

## Running the Application Locally

1. **Clone the repository and enter the directory:**
   ```bash
   git clone <repository-url>
   cd order-service
   ```

2. **Install Go dependencies:**
   ```bash
   go mod download
   ```

3. **Start the server:**
   ```bash
   go run main.go
   ```
   *The service will start accepting connections on port 4000.*

## API Endpoints

### 1. Health Check
- **Endpoint**: `GET /health`
- **Description**: Returns the operational status of the service.
- **Response**: `200 OK`

### 2. Get All Orders
- **Endpoint**: `GET /orders`
- **Description**: Retrieves a list of all existing orders.
- **Response Example**:
  ```json
  [
    {
      "id": "101",
      "item": "Laptop",
      "quantity": 1
    }
  ]
  ```

### 3. Create an Order
- **Endpoint**: `POST /orders`
- **Description**: Places a new order.
- **Request Payload**:
  ```json
  {
    "item": "Laptop",
    "quantity": 1
  }
  ```

## Repository Structure

```text
order-service/
  ├── main.go       # Application entry point
  ├── handlers/     # HTTP request handlers and controllers
  ├── services/     # Core business logic
  ├── models/       # Database schemas and models
  ├── go.mod        # Go module definition and dependencies
  ├── Dockerfile    # Docker image configuration
  ├── Jenkinsfile   # Jenkins CI/CD pipeline configuration
  └── k8s.yaml      # Kubernetes Deployment and Service definitions
```

## Docker & Kubernetes Deployment

### Docker Setup

1. **Build the image:**
   ```bash
   docker build -t order-service:v1 .
   ```

2. **Run the container:**
   ```bash
   docker run -p 4000:4000 --env-file .env order-service:v1
   ```

### Kubernetes Setup

The provided `k8s.yaml` file configures a Deployment containing 5 replicas and a NodePort service on port 4000. 

> **Important**: If you are using Minikube locally, ensure you build the Docker image securely inside the Minikube environment (`eval $(minikube docker-env)`) before applying the deployment, or set `imagePullPolicy: Never` appropriately as already defined in the configuration.

1. **Deploy to the cluster:**
   ```bash
   kubectl apply -f k8s.yaml
   ```

2. **Verify the deployment:**
   ```bash
   kubectl get deployments
   kubectl get pods
   kubectl get svc order-service
   ```
