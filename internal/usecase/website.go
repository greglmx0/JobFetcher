package usecase

import (
	"JobFetcher/internal/domain"
	"JobFetcher/internal/repository"
)

type WebsiteUseCase struct {
	websiteRepo *repository.WebsiteRepository
}

func NewWebsiteUseCase(repo *repository.WebsiteRepository) *WebsiteUseCase {
	return &WebsiteUseCase{websiteRepo: repo}
}

func (w *WebsiteUseCase) GetAllWebsites() ([]domain.Website, error) {
	return w.websiteRepo.GetAllWebsites()
}

func (w *WebsiteUseCase) CreateWebsite(website *domain.Website) (*domain.Website, error) {
	return w.websiteRepo.CreateWebsite(website)
}

func (w *WebsiteUseCase) GetWebsiteByName(name string) (*domain.Website, error) {
	return w.websiteRepo.GetWebsiteByName(name)
}
