package app

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"route256/ws5/internal/model"
	pb "route256/ws5/pkg"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := i.note.Create(ctx, getRequest(req))
	if err != nil {
		if errors.Is(err, model.ErrInvalidRequest) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}

func getRequest(req *pb.CreateRequest) model.NoteInfo {
	return model.NoteInfo{
		Title:   req.GetInfo().GetTitle(),
		Content: req.GetInfo().GetContent(),
		UserID:  req.GetUserId(),
	}
}
