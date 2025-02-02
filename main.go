package main

import (
    "fmt"
    "github.com/robfig/cron"

    "database/sql"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"

    "JobFetcher/internal/usecase"
    "JobFetcher/internal/handler"
    "JobFetcher/internal/repository"

)

func main() {
    c := cron.New()
    c.AddFunc("*/5 * * * * *", func() {
        fmt.Println("Hello, world! 2")
    })
    c.Start()

    db, err := sql.Open("sqlite3", "./db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    userRepo := repository.NewUserRepository(db)
    userUseCase := usecase.NewUserUseCase(userRepo)
    userHandler := handlers.NewUserHandler(userUseCase)

    r := mux.NewRouter()
    r.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUserHandler).Methods("GET")
    r.HandleFunc("/users", userHandler.GetAllUsersHandler).Methods("GET")
    r.HandleFunc("/user", userHandler.CreateUserHandler).Methods("POST")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))

    select {}
}
