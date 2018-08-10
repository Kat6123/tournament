package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kat6123/tournament/log"
	"gopkg.in/yaml.v2"
)

func fromYAML(path string) (*Configuration, error) {
	c := new(Configuration)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("read file: %v", err)
	}

	err = yaml.Unmarshal(content, c)
	if err != nil {
		return c, fmt.Errorf("unmarshal: %v", err)
	}

	return c, nil
}

func fromEnv() (*Configuration, error) {
	c := new(Configuration)

	c.DB.URI = os.Getenv("DB_URI")
	c.DB.TourCollection = os.Getenv("TOURS")
	c.DB.PlayerCollection = os.Getenv("PLAYERS")
	c.Port = os.Getenv("PORT")

	d, ok := os.LookupEnv("LOG_LEVEL")
	if ok {
		err := c.Debug.Set(d)
		if err != nil {
			return nil, fmt.Errorf("config LOG_LEVEL from env: %v", err)
		}
	}

	return c, nil
}

var (
	dbURI      = flag.String("uri", defaultConfig.DB.URI, "database uri for connection")
	tours      = flag.String("tours", defaultConfig.DB.TourCollection, "tour collection name")
	players    = flag.String("players", defaultConfig.DB.PlayerCollection, "player collection name")
	port       = flag.String("port", defaultConfig.Port, "port number")
	debugLevel = log.Flag("log-level", defaultConfig.Debug, "log level")

	yamlPath = flag.String("yaml", "", "path to yaml file")
)

func fromFlags() *Configuration {
	c := &Configuration{
		DB: dbConfig{
			URI:              *dbURI,
			TourCollection:   *tours,
			PlayerCollection: *players,
		},
		Port:  *port,
		Debug: *debugLevel,
	}

	return c
}
