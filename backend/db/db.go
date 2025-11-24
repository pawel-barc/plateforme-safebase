package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"safebase/config"
)

var DB *sql.DB

// ConnecDB établit une connexion à la base des données PostgreSQL
func ConnectDB() {
	var err error

	// Chaîne de connexion
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.Cfg.DBUser,
		config.Cfg.DBPassword,
		config.Cfg.DBName,
		config.Cfg.DBHost,
		config.Cfg.DBPort,
		config.Cfg.DBSSLMode,
	)

	// Initialisation de la connexion
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base des données:", err)
	}

	// Verifie la connexion réelle à la base des données
	if err = DB.Ping(); err != nil {
		log.Fatal("Impossible de se connecter à la base des données", err)
	}

	log.Println("Connexion à la base des donnès réussie")
}