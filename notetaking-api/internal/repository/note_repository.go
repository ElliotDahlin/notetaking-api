package repository

import (
    "github.com/ElliotDahlin/notetaking-api/internal/domain"
)

type NoteRepository interface {
    Create(note *domain.Note) error
    GetByID(id string) (*domain.Note, error)
    GetAll() ([]*domain.Note, error)
    Update(note *domain.Note) error
    Delete(id string) error
}

type InMemoryNoteRepository struct {
    notes map[string]*domain.Note
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
    return &InMemoryNoteRepository{
        notes: make(map[string]*domain.Note),
    }
}

func (r *InMemoryNoteRepository) Create(note *domain.Note) error {
    r.notes[note.ID] = note
    return nil
}

func (r *InMemoryNoteRepository) GetByID(id string) (*domain.Note, error) {
    note, exists := r.notes[id]
    if !exists {
        return nil, nil
    }
    return note, nil
}

func (r *InMemoryNoteRepository) GetAll() ([]*domain.Note, error) {
    notes := make([]*domain.Note, 0, len(r.notes))
    for _, note := range r.notes {
        notes = append(notes, note)
    }
    return notes, nil
}

func (r *InMemoryNoteRepository) Update(note *domain.Note) error {
    r.notes[note.ID] = note
    return nil
}

func (r *InMemoryNoteRepository) Delete(id string) error {
    delete(r.notes, id)
    return nil
}
