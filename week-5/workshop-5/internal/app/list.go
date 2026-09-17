package app

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"route256/ws5/internal/model"
	pb "route256/ws5/pkg"
)

func (i *Implementation) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	notes, err := i.note.List(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListResponse{
		Notes: getResponse(notes),
	}, nil
}

func getResponse(notes []model.Note) []*pb.ListResponse_Note {
	result := make([]*pb.ListResponse_Note, 0, len(notes))
	for _, note := range notes {
		result = append(result, &pb.ListResponse_Note{
			Id: note.ID,
			Info: &pb.ListResponse_Note_NoteInfo{
				Title:   note.Info.Title,
				Content: note.Info.Content,
			},
			CreatedAt: timestamppb.New(note.CreatedAt),
			UpdatedAt: timestamppb.New(note.UpdatedAt),
		})
	}
	return result
}
