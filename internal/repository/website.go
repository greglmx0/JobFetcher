package repository

import (
	"JobFetcher/internal/domain"
	"database/sql"
)

type WebsiteRepository struct {
	db *sql.DB
}

func NewWebsiteRepository(db *sql.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

// func (r *WebsiteRepository) GetWebsiteByID(id int) (*domain.Website, error) {
// 	var website domain.Website
// 	err := r.db.QueryRow("SELECT id, name, url FROM websites WHERE id = ?", id).Scan(&website.ID, &website.Name, &website.URL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &website, nil
// }

func (r *WebsiteRepository) GetAllWebsites() ([]domain.Website, error) {
	rows, err := r.db.Query("SELECT id, name, url, methode FROM websites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var websites []domain.Website
	for rows.Next() {
		var website domain.Website
		err := rows.Scan(&website.ID, &website.Name, &website.URL, &website.Methode)
		if err != nil {
			return nil, err
		}
		websites = append(websites, website)
	}

	return websites, nil
}

func (r *WebsiteRepository) CreateWebsite(website *domain.Website) (*domain.Website, error) {
	_, err := r.db.Exec("INSERT INTO websites (name, url, methode) VALUES (?, ?, ?)", website.Name, website.URL, website.Methode)
	if err != nil {
		return nil, err
	}
	return website, nil
}

func (r *WebsiteRepository) GetWebsiteByName(name string) (*domain.Website, error) {
	var website domain.Website
	err := r.db.QueryRow("SELECT id, name, url,	methode FROM websites WHERE name = ?", name).Scan(&website.ID, &website.Name, &website.URL, &website.Methode)
	if err != nil {
		return nil, err
	}
	return &website, nil
}
