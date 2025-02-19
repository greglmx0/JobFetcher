package handlers

import (
	"JobFetcher/internal/domain"
	"JobFetcher/internal/usecase"
	"encoding/json"
	"net/http"
)

type WebsiteHandler struct {
	websiteUseCase *usecase.WebsiteUseCase
}

func NewWebsiteHandler(uc *usecase.WebsiteUseCase) *WebsiteHandler {
	return &WebsiteHandler{websiteUseCase: uc}
}

func (h *WebsiteHandler) GetAllWebsitesHandler(w http.ResponseWriter, r *http.Request) {
	websites, err := h.websiteUseCase.GetAllWebsites()
	if err != nil {
		http.Error(w, "Failed to fetch websites", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(websites)
}

func (h *WebsiteHandler) CreateWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	var website domain.Website
	if err := json.NewDecoder(r.Body).Decode(&website); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if _, err := h.websiteUseCase.GetWebsiteByName(website.Name); err == nil {
		http.Error(w, "Website already exists", http.StatusConflict)
		return
	}

	if _, err := h.websiteUseCase.CreateWebsite(&website); err != nil {
		http.Error(w, "Failed to create website", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(website)
}
