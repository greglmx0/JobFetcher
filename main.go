package main

import (
	"fmt"
	"os"

	"github.com/robfig/cron"

	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	handlers "JobFetcher/internal/handler"
	"JobFetcher/internal/repository"
	"JobFetcher/internal/usecase"
)




func main() {
    c := cron.New()
    i := 0
    c.AddFunc("*/5 * * * * *", func() {
        i++
        fmt.Println("Hello, world!" + fmt.Sprint(i))
    })
    c.Start()

    dbPath := "/app/data/jobfetcher.db"
    ensureDBFolderExists() // Assure que le dossier existe
    db, err := sql.Open("sqlite3", dbPath+"?_cache=shared&mode=rwc")
    // db, err := sql.Open("sqlite3", "/root/db/jobfetcher.db?_cache=shared&mode=rwc")
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

func ensureDBFolderExists() {
    err := os.MkdirAll("/app/data", 0755)
    if err != nil {
        log.Fatal(err)
    }
}
