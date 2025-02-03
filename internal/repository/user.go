package repository

import (
	"JobFetcher/internal/domain"
	"database/sql"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (*domain.User, error) {
    var user domain.User
    err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
    rows, err := r.db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        var user domain.User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    _, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
    if err != nil {
        return err
    }
    return nil
}
