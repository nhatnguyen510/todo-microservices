# Todo Microservices Application

## Overview

The Todo Microservices Application is a scalable and efficient system built using Go, gRPC, and Docker. This application provides a set of microservices for managing todo items, allowing users to create, retrieve, update, and delete todos. Each service operates independently, promoting a microservices architecture.

## Features

- **Microservices Architecture**: Each service (Todo, Auth, User, Notification) is designed to function independently.
- **gRPC Communication**: Services communicate using gRPC, ensuring high performance and low latency.
- **Dockerized**: The application is fully containerized with Docker, simplifying deployment and management.
- **Protocol Buffers**: Utilizes Protocol Buffers for defining service interfaces and message types.

## Services

- **Todo Service**: Handles all operations related to todo items.
- **Auth Service**: Manages user authentication and authorization.
- **User Service**: Responsible for user profile management.
- **Notification Service**: Sends notifications related to todo activities.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.22.3 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/nhatnguyen0510/todo-microservices-app.git
cd todo-microservices-app
```

### Generate Go Files from Proto Definitions

Run the following script to generate Go files from the Protocol Buffers definitions:

```bash
make proto-gen
```

### Build and Run the Application

Use Docker Compose to build and run the application:

```bash
make build
```

### Access the Services

Once the services are running, you can access them via the following endpoints:

- **Todo Service**: `http://localhost:8082`
- **Auth Service**: `http://localhost:8083`
- **User Service**: `http://localhost:8085`
- **Gateway Service**: `http://localhost:8080`

### Stopping the Services

To stop all services, run:

```bash
make stop
```

## Directory Structure

todo-microservices-app/
├── cmd/ # Command-line applications for each service
├── api/ # API definitions and Protocol Buffers
├── configs/ # Configuration files for services
├── deployments/ # Deployment configurations (Docker, Kubernetes)
├── internal/ # Internal packages and logic
├── pkg/ # Shared packages and libraries
├── scripts/ # Utility scripts for building and managing the application
├── go.mod # Go module file
└── go.sum # Go module dependencies

## Contributing

Contributions are welcome! If you have suggestions or improvements, please submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go Kit](https://gokit.io/) for building microservices in Go.
- [gRPC](https://grpc.io/) for high-performance RPC framework.
- [Docker](https://www.docker.com/) for containerization.
