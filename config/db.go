package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	mapBoxKey string
	GGKey     string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ggMapKey := os.Getenv("GOOGLE_MAP_API_KEY")
	mapBoxKey := os.Getenv("MAPBOX_API_TOKEN")
	return Config{
		GGKey:     ggMapKey,
		mapBoxKey: mapBoxKey,
	}
}
