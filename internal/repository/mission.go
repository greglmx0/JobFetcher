package repository

import (
	"JobFetcher/internal/domain"
	"database/sql"
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

	var missions []*domain.Mission
	rows, err := r.db.Query("SELECT id, website_id, website_source, mission_title, mission_posted_date, organization_name, country_name, city_name, mission_duration, mission_start_date, view_counter, candidate_counter FROM missions WHERE website_source = ?", websiteSource)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Vérifier si des missions ont été trouvées sinon retourner une slice vide
	if len(missions) == 0 {
		return []*domain.Mission{}, nil
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
