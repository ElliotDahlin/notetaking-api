package main

import (
    "log"
    "net/http"

    "github.com/ElliotDahlin/notetaking-api/internal/config"
    deliveryHTTP "github.com/ElliotDahlin/notetaking-api/internal/delivery/http"
    "github.com/ElliotDahlin/notetaking-api/internal/repository"
    "github.com/ElliotDahlin/notetaking-api/internal/service"
    "github.com/ElliotDahlin/notetaking-api/internal/usecase"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    noteRepo := repository.NewInMemoryNoteRepository()
    noteService := service.NewNoteService(noteRepo)
    noteUsecase := usecase.NewNoteUsecase(noteService)
    router := deliveryHTTP.NewRouter(noteUsecase)

    log.Println("Starting server on", cfg.ServerAddress)
    if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
        log.Fatal(err)
    }
}
