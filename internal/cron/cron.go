package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"JobFetcher/internal/domain"
	"JobFetcher/internal/repository"
	"JobFetcher/internal/telegram"

	"github.com/robfig/cron"
)

type APIResponse struct {
	Result interface{} `json:"result"`
	Count  int         `json:"count"`
}

// InitCron initialise et démarre le planificateur de tâches cron
func InitCron(telegramBot *telegram.TelegramBot, telegramChatID int64, webSiteRepo *repository.WebsiteRepository, missionRepo *repository.MissionRepository) {
	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		websites, err := webSiteRepo.GetAllWebsites()
		if err != nil {
			log.Printf("Erreur lors de la récupération des sites web: %v", err)
			return
		}

		newMissions := fetchAndProcessMissions(websites, missionRepo)

		// Envoi des nouvelles missions par Telegram
		sendTelegramMessages(telegramBot, telegramChatID, newMissions)
	})
	c.Start()
}

func fetchAndProcessMissions(websites []domain.Website, missionRepo *repository.MissionRepository) []domain.Mission {
	var newMissions []domain.Mission

	for _, website := range websites {
		log.Printf("Traitement du site web: %v", website.Name)

		switch website.Source {
		case "VIE":
			rawMissions, err := PostRequest(website.URL, website.Body)
			if err != nil {
				log.Printf("Erreur lors de la requête POST pour %s: %v", website.Name, err)
				continue
			}

			missions, err := decodeVIEResponse(rawMissions)
			if err != nil {
				log.Printf("Erreur lors du décodage de la réponse pour %s: %v", website.Name, err)
				continue
			}

			newMissions = append(newMissions, processVIEMissions(missions, website.Name, missionRepo)...)
		}
	}
	return newMissions
}

func sendTelegramMessages(telegramBot *telegram.TelegramBot, telegramChatID int64, missions []domain.Mission) {
	for _, mission := range missions {
		message := fmt.Sprintf("🔹 Mission: %s \n Organisation: %s \n Pays: %s \n Ville: %s \n Durée: %d mois \n Vues: %d \n Candidats: %d \n Annonce postée le: %s",
			mission.MissionTitle, mission.OrganizationName, mission.CountryName, mission.CityName, mission.MissionDuration, mission.ViewCounter, mission.CandidateCounter, mission.MissionStartDate)

		telegramBot.SendMessage(telegramChatID, message)
	}
}

func PostRequest(url string, body string) (interface{}, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de la requête POST: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'envoi de la requête POST: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statut HTTP inattendu: %s", resp.Status)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage JSON: %v", err)
	}

	return apiResponse.Result, nil
}
