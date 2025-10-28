package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DriverSvcUrl string
	TripSvcUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load("pkg/config/envs/dev.env")

	if err != nil {
		log.Panicf("Could not load api gateway .env file:%v\n", err)
	}

	cfg := &Config{
		Port:         getEnv("PORT", ":3000"),
		DriverSvcUrl: getEnv("DRIVER_SVC_URL", "localhost:50053"),
		TripSvcUrl:   getEnv("TRIP_SVC_URL", "localhost:50052"),
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
