package repository

import (
	"JobFetcher/internal/domain"
	"database/sql"
	"fmt"
	"log"
)

type MissionRepository struct {
	db *sql.DB
}

func NewMissionRepository(db *sql.DB) *MissionRepository {
	return &MissionRepository{db: db}
}

func (r *MissionRepository) GetMissionByID(id int) (*domain.Mission, error) {
	var mission domain.Mission
	err := r.db.QueryRow("SELECT id, website_id, website_source, mission_title, mission_posted_date, organization_name, country_name, city_name, mission_duration, mission_start_date, view_counter, candidate_counter FROM missions WHERE id = ?", id).Scan(&mission.ID, &mission.WebsiteId, &mission.WebsiteSource, &mission.MissionTitle, &mission.MissionPostedDate, &mission.OrganizationName, &mission.CountryName, &mission.CityName, &mission.MissionDuration, &mission.MissionStartDate, &mission.ViewCounter, &mission.CandidateCounter)
	if err != nil {
		return nil, err
	}
	return &mission, nil
}

func (r *MissionRepository) GetAllMissions() ([]domain.Mission, error) {
	rows, err := r.db.Query("SELECT id, website_id, website_source, mission_title, mission_posted_date, organization_name, country_name, city_name, mission_duration, mission_start_date, view_counter, candidate_counter FROM missions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []domain.Mission
	for rows.Next() {
		var mission domain.Mission
		err := rows.Scan(&mission.ID, &mission.WebsiteId, &mission.WebsiteSource, &mission.MissionTitle, &mission.MissionPostedDate, &mission.OrganizationName, &mission.CountryName, &mission.CityName, &mission.MissionDuration, &mission.MissionStartDate, &mission.ViewCounter, &mission.CandidateCounter)
		if err != nil {
			return nil, err
		}
		missions = append(missions, mission)
	}

	return missions, nil
}

func (r *MissionRepository) CreateMission(mission *domain.Mission) (*domain.Mission, error) {
	_, err := r.db.Exec("INSERT INTO missions (website_id, website_source, mission_title, mission_posted_date, organization_name, country_name, city_name, mission_duration, mission_start_date, view_counter, candidate_counter) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", mission.WebsiteId, mission.WebsiteSource, mission.MissionTitle, mission.MissionPostedDate, mission.OrganizationName, mission.CountryName, mission.CityName, mission.MissionDuration, mission.MissionStartDate, mission.ViewCounter, mission.CandidateCounter)
	if err != nil {
		return nil, err
	}
	return mission, nil
}

func (r *MissionRepository) GetMissionByWebsiteSource(websiteSource string) ([]*domain.Mission, error) {
	log.Printf("GetMissionByWebsiteSource: %v", websiteSource)

	var missions []*domain.Mission
	// Utilisation de LIKE pour rechercher les missions dont le champ website_source contient la sous-chaîne
	query := "SELECT id, website_id, website_source, mission_title, mission_posted_date, organization_name, country_name, city_name, mission_duration, mission_start_date, view_counter, candidate_counter FROM missions WHERE website_source LIKE ?"

	// Ajout des pourcentages (%) autour de la chaîne pour indiquer que la sous-chaîne peut être à n'importe quel endroit
	rows, err := r.db.Query(query, "%"+websiteSource+"%")
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de l'exécution de la requête: %v", err)
	}
	defer rows.Close()

	// Parcours des résultats de la requête et remplissage de la slice missions
	for rows.Next() {
		var mission domain.Mission
		err := rows.Scan(
			&mission.ID,
			&mission.WebsiteId,
			&mission.WebsiteSource,
			&mission.MissionTitle,
			&mission.MissionPostedDate,
			&mission.OrganizationName,
			&mission.CountryName,
			&mission.CityName,
			&mission.MissionDuration,
			&mission.MissionStartDate,
			&mission.ViewCounter,
			&mission.CandidateCounter)
		if err != nil {
			return nil, fmt.Errorf("Erreur lors de la lecture des résultats: %v", err)
		}

		// Ajouter la mission à la slice
		missions = append(missions, &mission)
	}

	// Vérification des erreurs de parcours des lignes
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Erreur lors du parcours des résultats: %v", err)
	}

	// Si aucune mission n'est trouvée, on retourne une slice vide
	if len(missions) == 0 {
		log.Printf("Aucune mission trouvée pour %v", websiteSource)
	}

	return missions, nil
}

func (r *MissionRepository) GetMissionsByWebsiteSourceAndWebsiteID(websiteSource string, websiteID int) ([]domain.Mission, error) {
	rows, err := r.db.Query("SELECT id, website_id, website_source, mission_title, mission_posted_date, organization_name, country_name, city_name, mission_duration, mission_start_date, view_counter, candidate_counter FROM missions WHERE website_source = ? AND website_id = ?", websiteSource, websiteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []domain.Mission
	for rows.Next() {
		var mission domain.Mission
		err := rows.Scan(&mission.ID, &mission.WebsiteId, &mission.WebsiteSource, &mission.MissionTitle, &mission.MissionPostedDate, &mission.OrganizationName, &mission.CountryName, &mission.CityName, &mission.MissionDuration, &mission.MissionStartDate, &mission.ViewCounter, &mission.CandidateCounter)
		if err != nil {
			return nil, err
		}
		missions = append(missions, mission)
	}

	return missions, nil
}
