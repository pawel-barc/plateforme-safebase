package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort string
	DBSSLMode string
	JWTSecret []byte
}

var Cfg Config
// Gestion des données sensibles stockées dans le fichier .env
func init() {
	godotenv.Load();

	Cfg = Config{
		DBUser: getEnvOrFail("DB_USER"),
		DBPassword: getEnvOrFail("DB_PASSWORD"),
		DBName: getEnvOrFail("DB_NAME"),
		JWTSecret: []byte(getEnvOrFail("JWT_SECRET")),

		DBPort: os.Getenv("DB_PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBSSLMode: os.Getenv("DB_SSLMODE"),

	}

	if Cfg.DBHost == "" {
		Cfg.DBHost = "localhost"
	}
	if Cfg.DBPort == "" {
		Cfg.DBPort = "5432"
	}
	if Cfg.DBSSLMode == "" {
		Cfg.DBSSLMode = "disable"
	}
}

func getEnvOrFail(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Variable environnementale %s non trouvée", key)
	}
	return value
}