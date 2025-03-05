package main

import (
	"JobFetcher/internal/cron"
	"JobFetcher/internal/db"
	"log"
	"net/http"
	"os"
	"strconv"

	handlers "JobFetcher/internal/handler"
	"JobFetcher/internal/repository"
	"JobFetcher/internal/telegram"
	"JobFetcher/internal/usecase"

	"github.com/gorilla/mux"
)

func main() {
	// get env variables TELEGRAM_BOT_TOKEN
	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	telegramChatID, err := strconv.ParseInt(telegramChatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion de TELEGRAM_CHAT_ID: %v", err)
	}

	log.Println("Initialisation du bot Telegram")
	telegramBot, err := telegram.NewTelegramBot(telegramToken)
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation du bot Telegram: %v", err)
	}

	// Initialiser la base de données
	db, err := db.InitDB("/app/data/jobfetcher.db")
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}
	defer db.Close()

	// Initialiser les dépendances
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)
	webSiteRepo := repository.NewWebsiteRepository(db)
	webSiteUseCase := usecase.NewWebsiteUseCase(webSiteRepo)
	websiteHandler := handlers.NewWebsiteHandler(webSiteUseCase)

	// Configurer le routeur HTTP
	r := mux.NewRouter()
	r.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", userHandler.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/user", userHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/website", websiteHandler.CreateWebsiteHandler).Methods("POST")
	r.HandleFunc("/websites", websiteHandler.GetAllWebsitesHandler).Methods("GET")
	r.HandleFunc("/website/{id:[0-9]+}", websiteHandler.DeleteWebsiteHandler).Methods("DELETE")

	// Initialiser le planificateur de tâches cron
	cron.InitCron(telegramBot, telegramChatID, webSiteRepo)

	log.Println("Serveur en cours d'exécution sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
