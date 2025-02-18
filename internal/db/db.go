package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initialise et retourne une connexion à la base de données SQLite
func InitDB(dbPath string) (*sql.DB, error) {
	ensureDBFolderExists(dbPath)

	db, err := sql.Open("sqlite3", dbPath+"?_cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	err = createUsersTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ensureDBFolderExists vérifie et crée le dossier de la base de données s'il n'existe pas
func ensureDBFolderExists(dbPath string) {
	dir := "/app/data"
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Erreur lors de la création du dossier de la base de données: %v", err)
	}
}

// createUsersTable crée la table utilisateurs si elle n'existe pas
func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	)`)
	return err
}
