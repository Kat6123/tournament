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
		URL              string `yaml:"url"`
		DB               string `yaml:"db"`
		TourCollection   string `yaml:"tour collection"`
		PlayerCollection string `yaml:"player collection"`
	}

	Configuration struct {
		DB    dbConfig
		Port  string    `yaml:"port"`
		Debug log.Level `yaml:"level"`
	}
)

// Not obvious???
func Default() Configuration {
	return defaultConfig
}
