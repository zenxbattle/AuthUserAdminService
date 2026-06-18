package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	Environment            string
	UserGRPCPort           string
	PostgresDSN            string
	JWTSecretKey           string
	APPURL                 string
	FRONTENDURL            string
	SMTPAppKey             string 
	SMTPHost               string 
	SMTPPort               string 
	SMTPUser               string 
	AdminPassword          string
	AdminUsername          string
	GoogleClientID         string
	GoogleClientSecret     string
	GoogleRedirectURL      string
	RedisURL               string
	ResendAPIKey           string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	pgUser := getEnv("POSTGRES_USER", "postgres")
	pgPassword := getEnv("POSTGRES_PASSWORD", "postgres")
	pgHost := getEnv("POSTGRES_HOST", "postgres.default.svc.cluster.local")
	pgPort := getEnv("POSTGRES_PORT", "5432")
	pgDB := getEnv("POSTGRES_DB", "postgres")
	pgDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUser, pgPassword, pgHost, pgPort, pgDB)

	config := Config{
		Environment:        getEnv("ENVIRONMENT", "development"),
		UserGRPCPort:       getEnv("USERGRPCPORT", "50051"),
		PostgresDSN:        pgDSN,
		JWTSecretKey:       getEnv("JWTSECRETKEY", "secretLeetcode"),
		APPURL:             getEnv("APPURL", "http://localhost:7000"),
		FRONTENDURL:        getEnv("FRONTENDURL", "http://localhost:8080"),
		SMTPAppKey:         getEnv("SMTPAPPKEY", ""),
		SMTPHost:           getEnv("SMTPHOST", "smtp.gmail.com"),
		SMTPPort:           getEnv("SMTPPORT", "587"),
		SMTPUser:           getEnv("SMTPUSER", "zenxbattle.space"),
		AdminPassword:      getEnv("ADMINPASSWORD", "admin"),
		AdminUsername:      getEnv("ADMINUSERNAME", "admin"),
		GoogleClientID:     getEnv("GOOGLECLIENTID", ""),
		GoogleClientSecret: getEnv("GOOGLECLIENTSECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLEREDIRECTURL", ""),
		RedisURL:           getEnv("REDISURL", "localhost:6379"),
		ResendAPIKey:       getEnv("RESENDAPIKEY", ""),
	}

	// fmt.Println(config)
	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	// log.Printf("Environment variable %s not set, using default: %s", key, defaultValue)
	return defaultValue
}
