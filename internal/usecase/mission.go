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
	return m.missionRepo.GetMissionByID(id)
}

func (m *MissionUseCase) GetMissionsByWebsiteSourceAndWebsiteID(name string, id int) ([]domain.Mission, error) {
	return m.missionRepo.GetMissionsByWebsiteSourceAndWebsiteID(name, id)
}
