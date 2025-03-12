package cron

import (
	"encoding/json"
	"fmt"
	"log"

	"JobFetcher/internal/domain"
	"JobFetcher/internal/repository"
)

// MissionVIEResponce représente la structure de la réponse JSON pour les missions VIE
type MissionVIEResponce struct {
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

// ConvertVIEMissionResponseToMission convertit une réponse VIE en une mission
func ConvertVIEMissionResponseToMission(mr MissionVIEResponce, websiteSource string) *domain.Mission {
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

// decodeVIEResponse décode une réponse JSON en un slice de MissionVIEResponce
func decodeVIEResponse(data interface{}) ([]MissionVIEResponce, error) {
	var response []MissionVIEResponce
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la conversion de l'interface en []byte: %v", err)
	}

	if err := json.Unmarshal(dataBytes, &response); err != nil {
		return nil, fmt.Errorf("Erreur lors du décodage JSON: %v", err)
	}

	return response, nil
}

// processVIEMissions traite les missions VIE et retourne les nouvelles missions
func processVIEMissions(missions []MissionVIEResponce, websiteSource string, missionRepo *repository.MissionRepository) []domain.Mission {
	var newMissions []domain.Mission

	for _, mission := range missions {
		convertedMission := ConvertVIEMissionResponseToMission(mission, websiteSource)

		existingMissions, err := missionRepo.GetMissionsByWebsiteSourceAndWebsiteID(convertedMission.WebsiteSource, convertedMission.WebsiteId)
		if err != nil {
			log.Printf("Erreur lors de la récupération des missions existantes: %v", err)
			continue
		}

		if len(existingMissions) == 0 {
			log.Printf("Sauvegarde de la nouvelle mission: %v", convertedMission.MissionTitle)
			if _, err := missionRepo.CreateMission(convertedMission); err != nil {
				log.Printf("Erreur lors de la sauvegarde de la mission: %v", err)
				continue
			}
			newMissions = append(newMissions, *convertedMission)
		}
	}
	return newMissions
}
