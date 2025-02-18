package main

import (
	"JobFetcher/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	handlers "JobFetcher/internal/handler"
	"JobFetcher/internal/repository"
	"JobFetcher/internal/telegram"
	"JobFetcher/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func main() {

	// get env variables TELEGRAM_BOT_TOKEN
	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	telegramChatID, err := strconv.ParseInt(telegramChatIDStr, 10, 64)

	log.Println("Initialisation du bot Telegram")
	telegramBot, err := telegram.NewTelegramBot(telegramToken)

	// Initialiser le planificateur de tâches cron
	c := cron.New()
	i := 0
	c.AddFunc("*/5 * * * * *", func() {
		i++
		log.Printf("Hello, world %d!", i)

		message := "Hello, world " + fmt.Sprint(i) + " !"
		telegramBot.SendMessage(int64(telegramChatID), message)
	})
	c.Start()

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

	log.Println("Serveur en cours d'exécution sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	// Bloquer indéfiniment pour garder le planificateur cron en cours d'exécution
	select {}
}
