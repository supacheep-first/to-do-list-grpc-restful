# To-Do List gRPC and RESTful API

This project provides a gRPC and RESTful API for managing a to-do list. It includes a gRPC server, a RESTful gateway, and a PostgreSQL database for storing tasks.

## Directory Structure

```
/Users/first/Workspaces/Codes/Go/to-do-list-grpc-restful/
  ├── proto/
  │   ├── google/
  │   │   ├── api/
  │   │   │   ├── annotations.proto
  │   │   │   ├── http.proto
  │   ├── todo.proto
  ├── server/
  │   ├── db/
  │   │   ├── db.go
  │   ├── models/
  │   │   ├── models.go
  │   ├── server.go
  ├── gateway/
  │   ├── gateway.go
  ├── README.md
```

## Prerequisites

- Go 1.16 or later
- PostgreSQL
- Protocol Buffers compiler (`protoc`)
- gRPC and gRPC-Gateway plugins for `protoc`

## Setup

### 1. Install Dependencies

Install the required Go packages:

```sh
go get -u google.golang.org/grpc
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/runtime
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### 2. Generate gRPC and gRPC-Gateway Code

Run the following command to generate the gRPC and gRPC-Gateway code:

```sh
protoc \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_out=. \
  --proto_path=./proto \
  proto/todo.proto
```

### 3. Set Up PostgreSQL

Create a PostgreSQL database and update the connection string in `server/db/db.go`:

```go
dsn := "host=localhost user=myuser password=mypassword dbname=myuser sslmode=disable"
```

### 4. Run the gRPC Server

Navigate to the `server` directory and run the gRPC server:

```sh
cd server
go run server.go
```

The gRPC server will be running on port `50051`.

### 5. Run the RESTful Gateway

Navigate to the `gateway` directory and run the RESTful gateway:

```sh
cd gateway
go run gateway.go
```

The RESTful API will be running on port `8080`.

## API Endpoints

### gRPC Endpoints

- `AddTask(TaskRequest) returns (TaskResponse)`
- `GetTasks(Empty) returns (TaskListResponse)`
- `CompleteTask(Task) returns (TaskResponse)`

### RESTful Endpoints

- `POST /tasks`: Add a new task
- `GET /tasks`: Get all tasks
- `PUT /tasks/{id}/complete`: Mark a task as completed

## License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.
