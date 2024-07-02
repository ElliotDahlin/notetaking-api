package service

import (
    "github.com/ElliotDahlin/notetaking-api/internal/domain"
    "github.com/ElliotDahlin/notetaking-api/internal/repository"
)

type NoteService struct {
    noteRepo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
    return &NoteService{noteRepo: repo}
}

func (s *NoteService) CreateNote(note *domain.Note) error {
    return s.noteRepo.Create(note)
}

func (s *NoteService) GetNoteByID(id string) (*domain.Note, error) {
    return s.noteRepo.GetByID(id)
}

func (s *NoteService) GetAllNotes() ([]*domain.Note, error) {
    return s.noteRepo.GetAll()
}

func (s *NoteService) UpdateNote(note *domain.Note) error {
    return s.noteRepo.Update(note)
}

func (s *NoteService) DeleteNote(id string) error {
    return s.noteRepo.Delete(id)
}
