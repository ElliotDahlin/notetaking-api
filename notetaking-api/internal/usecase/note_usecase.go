package usecase

import (
    "github.com/ElliotDahlin/notetaking-api/internal/domain"
    "github.com/ElliotDahlin/notetaking-api/internal/service"
)

type NoteUsecase struct {
    noteService *service.NoteService
}

func NewNoteUsecase(service *service.NoteService) *NoteUsecase {
    return &NoteUsecase{noteService: service}
}

func (u *NoteUsecase) CreateNote(note *domain.Note) error {
    return u.noteService.CreateNote(note)
}

func (u *NoteUsecase) GetNoteByID(id string) (*domain.Note, error) {
    return u.noteService.GetNoteByID(id)
}

func (u *NoteUsecase) GetAllNotes() ([]*domain.Note, error) {
    return u.noteService.GetAllNotes()
}

func (u *NoteUsecase) UpdateNote(note *domain.Note) error {
    return u.noteService.UpdateNote(note)
}

func (u *NoteUsecase) DeleteNote(id string) error {
    return u.noteService.DeleteNote(id)
}
