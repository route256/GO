package client

import (
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/model"
	desc "gitlab.ozon.dev/go/classroom-8/students/workshop-2/pkg/api/notes/v1"
)

func NoteToReq(n *model.Note) *desc.SaveNoteRequest {
	return &desc.SaveNoteRequest{Info: &desc.NoteInfo{
		Title:   n.Title,
		Content: n.Content,
	}}
}

func NoteFromResp(n *desc.Note) *model.Note {
	return &model.Note{
		Id:      int(n.GetNoteId()),
		Title:   n.GetInfo().GetTitle(),
		Content: n.GetInfo().GetContent(),
	}
}

func NotesFromResp(ns *desc.ListNotesResponse) []*model.Note {
	var resp []*model.Note
	for _, n := range ns.GetNotes() {
		resp = append(resp, NoteFromResp(n))
	}
	return resp
}
