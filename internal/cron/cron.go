package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"JobFetcher/internal/telegram"
	// get all web	sites
	"JobFetcher/internal/domain"
	"JobFetcher/internal/repository"

	"github.com/robfig/cron"
)

type MissionResponce struct {
	ID                int    `json:"id"`
	MissionTitle      string `json:"missionTitle"`
	MissionPostedDate string `json:"creationDate"`
	OrganizationName  string `json:"organizationName"`
	CountryName       string `json:"countryName"`
	CityName          string `json:"cityName"`
	MissionDuration   int    `json:"missionDuration"`
	MissionStartDate  string `json:"missionStartDate"`
	ViewCounter       int    `json:"viewCounter"`
	CandidateCounter  int    `json:"candidateCounter"`
}
type APIResponse struct {
	Result []MissionResponce `json:"result"`
	Count  int               `json:"count"`
}

// InitCron initialise et démarre le planificateur de tâches cron
func InitCron(telegramBot *telegram.TelegramBot, telegramChatID int64, webSiteRepo *repository.WebsiteRepository, missionRepo *repository.MissionRepository) {
	c := cron.New()
	i := 0
	c.AddFunc("*/5 * * * * *", func() {
		i++
		// log.Printf("Hello, world %d!", i)

		// get all websites from the database
		websites, _ := webSiteRepo.GetAllWebsites()
		// if err != nil {
		// 	log.Fatalf("Erreur lors de la récupération des sites web: %v", err)
		// }

		for _, website := range websites {
			log.Printf("Site web: %v", website)
			saveMissions, err := missionRepo.GetMissionByWebsiteSource(website.Name)
			if err != nil {
				log.Fatalf("Erreur lors de la récupération des missions: %v", err)
			}
			url := website.URL
			methode := website.Methode
			body := website.Body
			missions := []MissionResponce{}

			if methode == "POST" {
				log.Printf("🔹 Appel API POST: %s", url)
				missions, err = PostRequest(url, body)

				if err != nil {
					log.Printf("Erreur lors de l'appel API: %v", err)
					return
				}
			}

			log.Printf("Nombre de missions sauvegardées: %d", len(saveMissions))
			log.Printf("Nombre de missions: %d", len(missions))

			for _, mission := range missions {

				convertedMission := ConvertMissionResponseToMission(mission, website.Name)

				mis, err := missionRepo.GetMissionsByWebsiteSourceAndWebsiteID(convertedMission.WebsiteSource, convertedMission.WebsiteId)

				if err != nil {
					log.Fatalf("Erreur lors de la récupération des missions: %v", err)
				}

				if len(mis) == 0 {
					log.Printf("Sauvegarde de la mission: %v", convertedMission)
					_, err := missionRepo.CreateMission(convertedMission)
					if err != nil {
						log.Printf("Erreur lors de la sauvegarde de la mission: %v", err)
					}
					// message := "Hello, world " + fmt.Sprint(i) + " !"
					message := fmt.Sprintf("🔹 Mission: %s | Organisation: %s | Pays: %s | Ville: %s | Durée: %d mois | Vues: %d | Candidats: %d",
						mission.MissionTitle, mission.OrganizationName, mission.CountryName, mission.CityName, mission.MissionDuration, mission.ViewCounter, mission.CandidateCounter)

					telegramBot.SendMessage(telegramChatID, message)
				}
			}

			// if methode == "GET" {
			// 	GetRequest(url)
			// }

		}

	})
	c.Start()
}

func PostRequest(url string, body string) ([]MissionResponce, error) {
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

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage JSON: %v", err)
	}

	return apiResponse.Result, nil
}

func ConvertMissionResponseToMission(mr MissionResponce, websiteSource string) *domain.Mission {
	return &domain.Mission{
		WebsiteId:         mr.ID,
		WebsiteSource:     websiteSource,
		MissionTitle:      mr.MissionTitle,
		MissionPostedDate: mr.MissionPostedDate,
		OrganizationName:  mr.OrganizationName,
		CountryName:       mr.CountryName,
		CityName:          mr.CityName,
		MissionDuration:   mr.MissionDuration,
		MissionStartDate:  mr.MissionStartDate,
		ViewCounter:       mr.ViewCounter,
		CandidateCounter:  mr.CandidateCounter,
	}
}
