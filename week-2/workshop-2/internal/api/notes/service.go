package notes

import (
	"context"

	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/converter/server"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/service/notes"
	servicepb "gitlab.ozon.dev/go/classroom-8/students/workshop-2/pkg/api/notes/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ servicepb.NotesServer = (*Service)(nil)

// Service - уровень Delivery
type Service struct {
	servicepb.UnimplementedNotesServer                // Базовый класс
	impl                               *notes.Service // usecase
}

func (s *Service) SaveNote(ctx context.Context, req *servicepb.SaveNoteRequest) (*servicepb.SaveNoteResponse, error) {
	// Валидация и преобразование в модели уровня домена / DTO
	note, err := server.NoteFromReq(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// Uecase
	id, err := s.impl.SaveNote(ctx, note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Ответа
	return &servicepb.SaveNoteResponse{NoteId: uint64(id)}, nil
}

func (s *Service) ListNotes(ctx context.Context, _ *emptypb.Empty) (*servicepb.ListNotesResponse, error) {
	// Uecase
	ns, err := s.impl.ListNotes(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Ответа // из доменной модели делаем protobuf
	return server.NotesToResp(ns), nil
}

func NewNotesServer(impl *notes.Service) *Service {
	return &Service{impl: impl}
}
