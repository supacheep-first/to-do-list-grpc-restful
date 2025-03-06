# To-Do List gRPC Service

This project is a simple To-Do list application implemented using gRPC in Go. It includes a gRPC server and client, and uses PostgreSQL as the database.

## Prerequisites

- Go 1.16 or later
- Docker and Docker Compose
- PostgreSQL

## Setup

### 1. Clone the repository

```sh
git clone https://github.com/yourusername/to-do-list-grpc.git
cd to-do-list-grpc
```

### 2. Start PostgreSQL using Docker Compose

```sh
docker-compose up -d
```

### 3. Install Go dependencies

```sh
go mod tidy
```

### 4. Generate gRPC code

```sh
protoc --go_out=. --go-grpc_out=. proto/todo.proto
```

### 5. Run the server

```sh
go run server/server.go
```

### 6. Run the client

```sh
go run client/client.go
```

## Project Structure

- `proto/todo.proto`: Protocol Buffers definition file.
- `server/server.go`: gRPC server implementation.
- `server/models/models.go`: Database models.
- `server/db/db.go`: Database initialization.
- `client/client.go`: gRPC client implementation.
- `docker-compose.yml`: Docker Compose configuration for PostgreSQL.

## gRPC Services

- `AddTask(TaskRequest)`: Adds a new task.
- `GetTasks(Empty)`: Retrieves all tasks.
- `CompleteTask(Task)`: Marks a task as completed.
- `DeleteTask(Task)`: Deletes a task.

## License

This project is licensed under the MIT License.
