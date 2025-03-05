package cron

import (
	"fmt"
	"log"

	"JobFetcher/internal/telegram"
	// get all web	sites
	"JobFetcher/internal/repository"

	"github.com/robfig/cron"
)

// InitCron initialise et démarre le planificateur de tâches cron
func InitCron(telegramBot *telegram.TelegramBot, telegramChatID int64, webSiteRepo *repository.WebsiteRepository) {
	c := cron.New()
	i := 0
	c.AddFunc("*/5 * * * * *", func() {
		i++
		// log.Printf("Hello, world %d!", i)

		// get all websites from the database
		websites, err := webSiteRepo.GetAllWebsites()
		if err != nil {
			log.Fatalf("Erreur lors de la récupération des sites web: %v", err)
		}

		for _, website := range websites {
			log.Printf("Site web: %v", website)

			// get the website content
		}

		message := "Hello, world " + fmt.Sprint(i) + " !"
		telegramBot.SendMessage(telegramChatID, message)
	})
	c.Start()
}
