package repository

import (
	"JobFetcher/internal/domain"
	"gorm.io/gorm"
)

// UserRepository gère les opérations CRUD sur les utilisateurs
type UserRepository struct {
    db *gorm.DB
}

// NewUserRepository crée une nouvelle instance de UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// GetUserByID récupère un utilisateur par son ID
func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
    var user domain.User
    result := r.db.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// GetAllUsers récupère tous les utilisateurs
func (r *UserRepository) GetAllUsers() ([]*domain.User, error) {
    var users []*domain.User
    result := r.db.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

// CreateUser crée un nouvel utilisateur et retourne l'objet mis à jour avec son ID
func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
    result := r.db.Create(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

// GetUserByEmail récupère un utilisateur par son email
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
    var user domain.User
    result := r.db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}
