# Note-taking API

This is a simple Note-taking API built with Golang using clean architecture principles.

## Project Structure

- `cmd/main.go` - Entry point of the application.
- `internal/config` - Configuration loading.
- `internal/domain` - Domain models.
- `internal/repository` - Data repository implementations.
- `internal/service` - Business logic services.
- `internal/usecase` - Use cases connecting services to handlers.
- `internal/delivery/http` - HTTP handlers and router.

## How to Run

1. Clone the repository:
    ```sh
    git clone https://github.com/ElliotDahlin/notetaking-api.git
    ```

2. Navigate to the project directory:
    ```sh
    cd notetaking-api
    ```

3. Run the application:
    ```sh
    go run cmd/main.go
    ```

## API Endpoints

- `POST /notes` - Create a new note.
- `GET /notes/{id}` - Get a note by ID.
- `GET /notes` - Get all notes.
- `PUT /notes/{id}` - Update a note by ID.
- `DELETE /notes/{id}` - Delete a note by ID.

## Postman
In this repo you can find postman collection file that you can use to test out the application


