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

func (u *UserUseCase) GetUserByID(id int) (*domain.User, error) {
    return u.userRepo.GetUserByID(uint(id)) // Conversion int -> uint
}

func (u *UserUseCase) GetAllUsers() ([]*domain.User, error) {
    return u.userRepo.GetAllUsers()
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, error) {
    return  u.userRepo.CreateUser(user)
}

func (u *UserUseCase) GetUserByEmail(email string) (*domain.User, error) {
    return u.userRepo.GetUserByEmail(email)
}