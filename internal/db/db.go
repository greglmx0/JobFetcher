package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initialise et retourne une connexion à la base de données SQLite
func InitDB(dbPath string) (*sql.DB, error) {
	ensureDBFolderExists(dbPath)

	db, err := sql.Open("sqlite3", dbPath+"?_cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	// err = dropAllTables(db)
	// if err != nil {
	// 	return nil, err
	// }

	err = createUsersTable(db)
	if err != nil {
		return nil, err
	}

	err = createWebSitesTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ensureDBFolderExists vérifie et crée le dossier de la base de données s'il n'existe pas
func ensureDBFolderExists(dbPath string) {

	length := strings.LastIndex(dbPath, "/")
	dir := dbPath[0:length]

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Erreur lors de la création du dossier de la base de données: %v", err)
	}
	// change le propriétaire du dossier
	err := os.Chown(dir, 1000, 1000)
	if err != nil {
		log.Fatalf("Erreur lors du changement de propriétaire du dossier de la base de données: %v", err)
	}
}

func dropAllTables(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS users`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS websites`)
	if err != nil {
		return err
	}

	return nil
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	)`)
	return err
}

func createWebSitesTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS websites (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		url TEXT NOT NULL,
		methode TEXT NOT NULL,
		body	TEXT NULL
	)`)
	return err
}
