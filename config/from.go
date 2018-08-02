package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/kat6123/tournament/log"
	"gopkg.in/yaml.v2"
)

func fromYAML(path string) (*Configuration, error) {
	c := new(Configuration)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		//  TODO you don't need "config from yaml" here
 		return c, fmt.Errorf("config from yaml: read file %q: %v", path, err)
	}

	err = yaml.Unmarshal([]byte(content), c)
	if err != nil {
		return c, fmt.Errorf("unmarshal: %v", err)
	}

	return c, nil
}

func fromEnv() (*Configuration, error) {
	c := new(Configuration)

	c.DB.URL = os.Getenv("DBURL") // TODO bad name
	c.DB.DB = os.Getenv("DB")
	c.DB.TourCollection = os.Getenv("TOURS")
	c.DB.PlayerCollection = os.Getenv("PLAYERS")
	c.Port = os.Getenv("PORT")

	d, ok := os.LookupEnv("DEBUG")
	if ok {
		err := c.Debug.Set(d)
		if err != nil {
			return nil, fmt.Errorf("config DEBUG from env: %v", err)
		}
	}

	return c, nil
}

var (
	dbURL      = flag.String("dburl", defaultConfig.DB.URL, "database url for connection")
	db         = flag.String("db", defaultConfig.DB.DB, "database name")
	tours      = flag.String("tours", defaultConfig.DB.TourCollection, "tour collection name")
	players    = flag.String("players", defaultConfig.DB.PlayerCollection, "player collection name")
	port       = flag.String("port", defaultConfig.Port, "port number")
	debugLevel = log.Flag("debug", defaultConfig.Debug, "log level")
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
