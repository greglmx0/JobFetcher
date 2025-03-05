package handlers

import (
	"JobFetcher/internal/domain"
	"JobFetcher/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// Vérifier si le site existe déjà
	if _, err := h.websiteUseCase.GetWebsiteByName(website.Name); err == nil {
		http.Error(w, "Website already exists", http.StatusConflict)
		return
	}

	// Créer le site
	if _, err := h.websiteUseCase.CreateWebsite(&website); err != nil {
		http.Error(w, "Failed to create website", http.StatusInternalServerError)
		return
	}

	// Récupérer le site avec son ID généré
	createdWebsite, err := h.websiteUseCase.GetWebsiteByName(website.Name)
	if err != nil {
		http.Error(w, "Failed to fetch created website", http.StatusInternalServerError)
		return
	}

	// Répondre avec l'objet JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWebsite)
}

func (h *WebsiteHandler) DeleteWebsiteHandler(w http.ResponseWriter, r *http.Request) {

	// Récupérer l'ID du site à supprimer
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid website ID", http.StatusBadRequest)
		return
	}

	// Supprimer le site par son ID
	if err := h.websiteUseCase.DeleteWebsiteByID(id); err != nil {
		http.Error(w, "Failed to delete website", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
