package config

import "github.com/kat6123/tournament/log"

var defaultConfig = Configuration{
	DB: dbConfig{
		URI:              "mongodb://localhost/tours",
		TourCollection:   "tours",
		PlayerCollection: "players",
	},
	Port:  "3001",
	Debug: log.ErrorLevel,
}

type (
	dbConfig struct {
		URI              string `yaml:"uri"`
		TourCollection   string `yaml:"tour_collection"`
		PlayerCollection string `yaml:"player_collection"`
	}

	// Configuration type describes configuration for project.
	Configuration struct {
		DB    dbConfig
		Port  string    `yaml:"app_port"`
		Debug log.Level `yaml:"log_level"`
	}
)
