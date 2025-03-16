package usecase

import (
	"JobFetcher/internal/domain"
	"JobFetcher/internal/repository"
)

type MissionUseCase struct {
	missionRepo *repository.MissionRepository
}

func NewMissionUseCase(repo *repository.MissionRepository) *MissionUseCase {
	return &MissionUseCase{missionRepo: repo}
}

func (m *MissionUseCase) GetAllMissions() ([]domain.Mission, error) {
	return m.missionRepo.GetAllMissions()
}

func (m *MissionUseCase) CreateMission(mission *domain.Mission) (*domain.Mission, error) {
	return m.missionRepo.CreateMission(mission)
}

func (m *MissionUseCase) GetMissionByWebsiteSource(websiteSource string) ([]*domain.Mission, error) {
	return m.missionRepo.GetMissionByWebsiteSource(websiteSource)
}

func (m *MissionUseCase) GetMissionByID(id int) (*domain.Mission, error) {
    return m.missionRepo.GetMissionByID(uint(id)) // Conversion int -> uint
}

func (m *MissionUseCase) GetMissionsByWebsiteSourceAndWebsiteID(websiteSource string, id int) ([]*domain.Mission, error) {
    return m.missionRepo.GetMissionsByWebsiteSourceAndWebsiteID(websiteSource, uint(id)) // Conversion int -> uint
}