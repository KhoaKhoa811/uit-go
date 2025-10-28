package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBUrl        string
	DriverSvcUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load("pkg/config/envs/dev.env")

	if err != nil {
		log.Panicf("Could not load trip .env file: %v\n", err)
	}

	cfg := &Config{
		Port:         getEnv("PORT", ":50052"),
		DBUrl:        getEnv("DB_URL", "localhost"),
		DriverSvcUrl: getEnv("DRIVER_SVC_URl", "localhost:50053"),
	}

	log.Printf("Config loaded: %+v\n", cfg)
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
