package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Modèle User
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}

// Modèle Website
type Website struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	URL    string `gorm:"not null"`
	Source string `gorm:"not null"`
	Method string `gorm:"not null"`
	Body   string
}

// Modèle Mission
type Mission struct {
	ID                uint   `gorm:"primaryKey"`
	WebsiteID         uint   `gorm:"not null"`
	WebsiteSource     string `gorm:"not null"`
	MissionTitle      string `gorm:"not null"`
	MissionPostedDate string `gorm:"not null"`
	OrganizationName  string `gorm:"not null"`
	CountryName       string `gorm:"not null"`
	CityName          string `gorm:"not null"`
	MissionDuration   int    `gorm:"not null"`
	MissionStartDate  string `gorm:"not null"`
	ViewCounter       int    `gorm:"not null"`
	CandidateCounter  int    `gorm:"not null"`
}

// InitDB initialise et retourne une connexion à SQLite avec GORM
func InitDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migration automatique des tables
	err = db.AutoMigrate(&User{}, &Website{}, &Mission{})
	if err != nil {
		return nil, err
	}

	log.Println("Base de données initialisée avec succès 🚀")
	return db, nil
}