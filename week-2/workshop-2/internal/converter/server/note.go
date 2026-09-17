package server

import (
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/model"
	desc "gitlab.ozon.dev/go/classroom-8/students/workshop-2/pkg/api/notes/v1"
)

func NoteFromReq(req *desc.SaveNoteRequest) (*model.Note, error) {
	err := req.ValidateAll() // автогенеренная валидация
	if err != nil {
		return nil, err
	}
	return &model.Note{
		Title:   req.GetInfo().GetTitle(),
		Content: req.GetInfo().GetContent(),
	}, nil
}

func NoteToResp(n *model.Note) *desc.Note {
	return &desc.Note{
		NoteId: uint64(n.Id),
		Info: &desc.NoteInfo{
			Title:   n.Title,
			Content: n.Content,
		},
	}
}

func NotesToResp(ns []*model.Note) *desc.ListNotesResponse {
	resp := &desc.ListNotesResponse{}
	for _, n := range ns {
		resp.Notes = append(resp.Notes, NoteToResp(n))
	}
	return resp
}
