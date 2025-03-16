package repository

import (
	"JobFetcher/internal/domain"
	"gorm.io/gorm"
)

// MissionRepository gère les opérations CRUD sur les missions
type MissionRepository struct {
	db *gorm.DB
}

// NewMissionRepository crée une nouvelle instance de MissionRepository
func NewMissionRepository(db *gorm.DB) *MissionRepository {
	return &MissionRepository{db: db}
}

// GetMissionByID récupère une mission par son ID
func (r *MissionRepository) GetMissionByID(id uint) (*domain.Mission, error) {
	var mission domain.Mission
	result := r.db.First(&mission, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &mission, nil
}

// GetAllMissions récupère toutes les missions
func (r *MissionRepository) GetAllMissions() ([]domain.Mission, error) {
	var missions []domain.Mission
	result := r.db.Find(&missions)
	if result.Error != nil {
		return nil, result.Error
	}
	return missions, nil
}

// CreateMission crée une nouvelle mission et retourne l'objet mis à jour avec son ID
func (r *MissionRepository) CreateMission(mission *domain.Mission) (*domain.Mission, error) {
	result := r.db.Create(mission)
	if result.Error != nil {
		return nil, result.Error
	}
	return mission, nil
}

// GetMissionByWebsiteSource récupère les missions dont le champ website_source contient une sous-chaîne
func (r *MissionRepository) GetMissionByWebsiteSource(websiteSource string) ([]*domain.Mission, error) {
	var missions []*domain.Mission
	result := r.db.Where("website_source LIKE ?", "%"+websiteSource+"%").Find(&missions)
	if result.Error != nil {
		return nil, result.Error
	}
	return missions, nil
}


// GetMissionsByWebsiteSourceAndWebsiteID récupère les missions filtrées par website_source et website_id
func (r *MissionRepository) GetMissionsByWebsiteSourceAndWebsiteID(websiteSource string, websiteID uint) ([]*domain.Mission, error) {
	var missions []*domain.Mission
	result := r.db.Where("website_source = ? AND website_id = ?", websiteSource, websiteID).Find(&missions)

	if result.Error != nil {
		return nil, result.Error
	}

	return missions, nil
}