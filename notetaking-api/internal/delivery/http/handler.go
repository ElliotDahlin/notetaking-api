package http

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ElliotDahlin/notetaking-api/internal/domain"
    "github.com/ElliotDahlin/notetaking-api/internal/usecase"
)

type Handler struct {
    noteUsecase *usecase.NoteUsecase
}

func NewHandler(usecase *usecase.NoteUsecase) *Handler {
    return &Handler{noteUsecase: usecase}
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
    var note domain.Note
    if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.noteUsecase.CreateNote(&note); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    note, err := h.noteUsecase.GetNoteByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if note == nil {
        http.NotFound(w, r)
        return
    }
    json.NewEncoder(w).Encode(note)
}

func (h *Handler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
    notes, err := h.noteUsecase.GetAllNotes()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(notes)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
    var note domain.Note
    if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.noteUsecase.UpdateNote(&note); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if err := h.noteUsecase.DeleteNote(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
