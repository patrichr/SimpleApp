package handler

import (
	"context"
	"net/http"

	pb "github.com/patrichr/SimpleApp/grpc"
	"github.com/patrichr/SimpleApp/server/model"
)

// integrate widget
func (s *Server) Register(ctx context.Context, in *pb.RegRequest) (*pb.RegResponse, error) {
	_, err := s.repo.Register(ctx, model.Register{
		TaskName:    in.GetTaskName(),
		TaskOwner:   in.GetTaskOwner(),
		Description: in.GetDescription(),
		Status:      in.GetStatus(),
	})

	if err != nil {
		return &pb.RegResponse{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}, err
	}
	return &pb.RegResponse{
		HttpStatus: http.StatusCreated,
		Message:    "successfully registered",
	}, err
}
