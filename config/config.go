package main

import (
	"fmt"
	"io/ioutil"

	"os"

	"gopkg.in/yaml.v2"
)

const (
	trace logLevel = iota
	debug
	info
	warn
	// error
)

type (
	logLevel int

	dbConfig struct {
		URL              string `yaml:"url"`
		DB               string `yaml:"db"`
		TourCollection   string `yaml:"tour collection"`
		PlayerCollection string `yaml:"player collection"`
	}

	Configuration struct {
		DB    dbConfig
		Port  int      `yaml:"port"`
		Debug logLevel `yaml:"level"`
	}
)

func ConfigFromYAML(path string) (*Configuration, error) {
	c := new(Configuration)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c, fmt.Errorf("config from yaml: read file %q: %v", path, err)
	}

	err = yaml.Unmarshal([]byte(content), c)
	if err != nil {
		return c, fmt.Errorf("config from yaml: unmarshal: %v", err)
	}

	return c, nil
}

func ConfigFromEnv() (*Configuration, error) {
	c := new(Configuration)

	os.Getenv("DB")
	ENV := []string{"DBURL", "DB", "TOURS", "PLAYERS", "PORT", "DEBUG"}

	for i := range ENV {

	}
}

func Get() (*Configuration, error) {
	return nil, nil
}

func main() {
	c, _ := ConfigFromYAML("config.yaml")
	fmt.Println(c)
}
