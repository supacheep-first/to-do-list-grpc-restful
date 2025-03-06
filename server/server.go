package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-todo/proto"
	"grpc-todo/server/db"
	"grpc-todo/server/models"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
    proto.UnimplementedTodoServiceServer
}

func (s *server) AddTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResponse, error) {
    task := models.Task{Title: req.Title, Completed: false}
    db.DB.Create(&task)

    return &proto.TaskResponse{
        Task: &proto.Task{Id: int32(task.ID), Title: task.Title, Completed: task.Completed},
    }, nil
}

func (s *server) GetTasks(ctx context.Context, req *emptypb.Empty) (*proto.TaskListResponse, error) {
    var tasks []models.Task
    db.DB.Find(&tasks)

    var protoTasks []*proto.Task
    for _, t := range tasks {
        protoTasks = append(protoTasks, &proto.Task{Id: int32(t.ID), Title: t.Title, Completed: t.Completed})
    }

    return &proto.TaskListResponse{Tasks: protoTasks}, nil
}

func (s *server) CompleteTask(ctx context.Context, req *proto.Task) (*proto.TaskResponse, error) {
    var task models.Task
    if err := db.DB.First(&task, req.Id).Error; err != nil {
        return nil, err
    }

    task.Completed = true
    db.DB.Save(&task)

    return &proto.TaskResponse{
        Task: &proto.Task{Id: int32(task.ID), Title: task.Title, Completed: task.Completed},
    }, nil
}

func (s *server) DeleteTask(ctx context.Context, req *proto.Task) (*emptypb.Empty, error) {
    if err := db.DB.Delete(&models.Task{}, req.Id).Error; err != nil {
        return nil, err
    }
    return &emptypb.Empty{}, nil
}

func main() {
    db.InitDB()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal("Failed to listen:", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterTodoServiceServer(grpcServer, &server{})

    fmt.Println("gRPC Server running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
