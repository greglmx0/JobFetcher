package handlers

import (
	"JobFetcher/internal/domain"
	"JobFetcher/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MissionHandler struct {
	missionUseCase *usecase.MissionUseCase
}

func NewMissionHandler(uc *usecase.MissionUseCase) *MissionHandler {
	return &MissionHandler{missionUseCase: uc}
}

func (h *MissionHandler) GetAllMissionsHandler(w http.ResponseWriter, r *http.Request) {
	missions, err := h.missionUseCase.GetAllMissions()
	if err != nil {
		http.Error(w, "Failed to fetch missions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(missions)
}

func (h *MissionHandler) CreateMissionHandler(w http.ResponseWriter, r *http.Request) {
	var mission domain.Mission
	if err := json.NewDecoder(r.Body).Decode(&mission); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Vérifier si la mission existe déjà pour le site avec le même websiteSource et website id
	if _, err := h.missionUseCase.GetMissionsByWebsiteSourceAndWebsiteID(mission.WebsiteSource, mission.WebsiteId); err == nil {
		http.Error(w, "Mission already exists", http.StatusConflict)
		return
	}

	// Créer la mission
	if _, err := h.missionUseCase.CreateMission(&mission); err != nil {
		http.Error(w, "Failed to create mission", http.StatusInternalServerError)
		return
	}

	// Récupérer la mission avec son ID généré
	createdMission, err := h.missionUseCase.GetMissionByWebsiteSource(mission.WebsiteSource)
	if err != nil {
		http.Error(w, "Failed to fetch created mission", http.StatusInternalServerError)
		return
	}

	// Répondre avec l'objet JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMission)
}

// func (h *MissionHandler) DeleteMissionHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid mission ID", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.missionUseCase.DeleteMission(id)
// 	if err != nil {
// 		http.Error(w, "Failed to delete mission", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Mission deleted successfully"})
// }

func (h *MissionHandler) GetMissionsByWebsiteSourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	websiteSource := vars["websiteSource"]

	mission, err := h.missionUseCase.GetMissionByWebsiteSource(websiteSource)
	if err != nil {
		http.Error(w, "Mission not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mission)
}
