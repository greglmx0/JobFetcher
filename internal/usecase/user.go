package usecase

import (
    "JobFetcher/internal/domain"
    "JobFetcher/internal/repository"
)

type UserUseCase struct {
    userRepo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
    return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) GetUser(id int) (*domain.User, error) {
    return u.userRepo.GetUserByID(id)
}

func (u *UserUseCase) GetAllUsers() ([]domain.User, error) {
    return u.userRepo.GetAllUsers()
}

func (u *UserUseCase) CreateUser(user *domain.User) error {
    return u.userRepo.CreateUser(user)
}