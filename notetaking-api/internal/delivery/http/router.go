package http

import (
    "github.com/gorilla/mux"
    "github.com/ElliotDahlin/notetaking-api/internal/usecase"
)

func NewRouter(noteUsecase *usecase.NoteUsecase) *mux.Router {
    handler := NewHandler(noteUsecase)
    router := mux.NewRouter()

    router.HandleFunc("/notes", handler.CreateNote).Methods("POST")
    router.HandleFunc("/notes/{id}", handler.GetNoteByID).Methods("GET")
    router.HandleFunc("/notes", handler.GetAllNotes).Methods("GET")
    router.HandleFunc("/notes/{id}", handler.UpdateNote).Methods("PUT")
    router.HandleFunc("/notes/{id}", handler.DeleteNote).Methods("DELETE")

    return router
}
