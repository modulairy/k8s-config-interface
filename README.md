# ConfigMap HTTP API Server

This HTTP server provides an API for updating and retrieving Kubernetes ConfigMaps. It can be run within a Docker container and deployed in a Kubernetes cluster.

## Prerequisites

- [Docker](https://www.docker.com/) installed on your machine.

## Usage

1. **Pull the Docker Image:**
   ```shell
   docker pull modulairy/k8s-configmap-api-server:latest
   ```
2. **Run the Container:**
   ```shell
   docker run -p 8080:8080 modulairy/k8s-configmap-api-server:latest
   ```
3. **Access the FastAPI Application:**
   Open your web browser and navigate to [http://localhost:8080](http://localhost:8000/). 

## Build and Run the Docker Container

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```
2. **Build the Docker Image:**

   ```bash
   docker build -t configmap-api-server:latest .
   ```

   This command will build a Docker image named `configmap-api-server` using the provided Dockerfile.
3. **Run the Docker Container:**

   ```bash
   docker run -p 8080:8080 -e PERMIT_CONFIG_NAME=your_config_name -e PERMIT_NAMESPACE=your_namespace -e PERMIT_CONFIG_KEY=your_config_key configmap-api-server:latest
   ```

   Replace `your_config_name`, `your_namespace`, and `your_config_key` with the appropriate values.

   - `your_config_name`: The name of the ConfigMap to be updated or retrieved.
   - `your_namespace`: The namespace in which the ConfigMap resides.
   - `your_config_key`: The key within the ConfigMap where the data will be stored.

   The server will start and listen on port 8080.

## API Endpoints

### Health Check

- **Endpoint:** `/health`
- **Method:** GET
- **Description:** Performs a health check and returns "Ok" if the server is running.

### Update ConfigMap

- **Endpoint:** `/config`
- **Method:** PUT
- **Query Parameters:**
  - `subscriptionId` (optional): Namespace of the ConfigMap.
  - `configName` (optional): Name of the ConfigMap.
  - `configKey` (optional): Key within the ConfigMap where the data will be stored.
- **Request Body:** JSON data to be stored in the ConfigMap.
- **Description:** Updates a ConfigMap with the provided data.

### Get ConfigMap

- **Endpoint:** `/config`
- **Method:** GET
- **Query Parameters:**
  - `subscriptionId` (optional): Namespace of the ConfigMap.
  - `configName` (optional): Name of the ConfigMap.
  - `configKey` (optional): Key within the ConfigMap to retrieve data from.
- **Description:** Retrieves data from a ConfigMap based on the provided parameters.

## Usage

- To update a ConfigMap, send a PUT request to `/config` with the appropriate query parameters and JSON data in the request body.

  Example:

  ```bash
  curl -X PUT -H "Content-Type: application/json" -d '{"key": "value"}' http://localhost:8080/config?subscriptionId=default&configName=my-config&configKey=data
  ```
- To retrieve data from a ConfigMap, send a GET request to `/config` with the appropriate query parameters.

  Example:

  ```bash
  curl http://localhost:8080/config?subscriptionId=default&configName=my-config&configKey=data
  ```

Note: Replace the query parameters with your specific values.
