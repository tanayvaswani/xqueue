package server

import (
	"context"

	"github.com/tanayvaswani/xqueue/pb"
)

type Server struct {
	pb.UnimplementedTaskServer
}

func (s *Server) CreateTask(
	ctx context.Context,
	in *pb.CreateTaskRequest,
) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{
		TaskId:   "1",
		TaskName: in.TaskName,
		Status:   "pending",
	}, nil
}

func (s *Server) GetTask(
	ctx context.Context,
	in *pb.GetTaskRequest,
) (*pb.GetTaskResponse, error) {
	return &pb.GetTaskResponse{
		TaskId:   in.TaskId,
		TaskName: "task1",
		Status:   "pending",
	}, nil
}
