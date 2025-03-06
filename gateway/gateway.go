package main

import (
	"context"
	"fmt"
	"grpc-todo/proto"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
    grpcConn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatal("Failed to connect to gRPC server:", err)
    }
    defer grpcConn.Close()

    mux := runtime.NewServeMux()
    err = proto.RegisterTodoServiceHandler(context.Background(), mux, grpcConn)
    if err != nil {
        log.Fatal("Failed to register gRPC-Gateway:", err)
    }

    ginRouter := gin.Default()
    ginRouter.Any("/tasks", gin.WrapH(mux))
    ginRouter.Any("/tasks/:id/complete", gin.WrapH(mux))

    fmt.Println("REST API running on port 8080...")
    if err := ginRouter.Run(":8080"); err != nil {
        log.Fatal("Failed to start Gin server:", err)
    }
}
