package repository

import (
	"JobFetcher/internal/domain"
	"gorm.io/gorm"
)

// WebsiteRepository gère les opérations CRUD sur les websites
type WebsiteRepository struct {
	db *gorm.DB
}

// NewWebsiteRepository crée une nouvelle instance de WebsiteRepository
func NewWebsiteRepository(db *gorm.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

// GetAllWebsites récupère tous les sites Web
func (r *WebsiteRepository) GetAllWebsites() ([]domain.Website, error) {
	var websites []domain.Website
	result := r.db.Find(&websites)
	if result.Error != nil {
		return nil, result.Error
	}
	return websites, nil
}

// CreateWebsite ajoute un nouveau site Web à la base de données
func (r *WebsiteRepository) CreateWebsite(website *domain.Website) (*domain.Website, error) {
	result := r.db.Create(website)
	if result.Error != nil {
		return nil, result.Error
	}
	return website, nil
}

// GetWebsiteByName récupère un site Web par son nom
func (r *WebsiteRepository) GetWebsiteByName(name string) (*domain.Website, error) {
	var website domain.Website
	result := r.db.Where("name = ?", name).First(&website)
	if result.Error != nil {
		return nil, result.Error
	}
	return &website, nil
}

// DeleteWebsiteByID supprime un site Web par son ID
func (r *WebsiteRepository) DeleteWebsiteByID(id uint) error {
	result := r.db.Delete(&domain.Website{}, id)
	return result.Error
}
