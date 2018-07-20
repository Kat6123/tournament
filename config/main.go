package config

import "flag"

func Merge(c1, c2 Configuration) Configuration {
	c := c1

	if c2.DB.URL != "" {
		c.DB.URL = c2.DB.URL
	}

	if c2.DB.DB != "" {
		c.DB.DB = c2.DB.DB
	}

	if c2.DB.TourCollection != "" {
		c.DB.TourCollection = c2.DB.TourCollection
	}

	if c2.DB.PlayerCollection != "" {
		c.DB.PlayerCollection = c2.DB.PlayerCollection
	}

	if c2.Port != "" {
		c.Port = c2.Port
	}

	// What if set 0 level??
	if c2.Debug > c.Debug {
		c.Debug = c2.Debug
	}

	return c
}

var yamlPath = flag.String("yaml", "", "path to yaml file")

func Get() Configuration {
	conf := Default()
	flag.Parse()

	flags, err := fromFlags()
	if err != nil {
		// write to log
	} else {
		conf = Merge(conf, *flags)
	}

	// default value "" causes error in log each time?
	yaml, err := fromYAML(*yamlPath)
	if err != nil {
		// write to log
	} else {
		conf = Merge(conf, *yaml)
	}

	env, err := fromEnv()
	if err != nil {
		// write to log
	} else {
		conf = Merge(conf, *env)
	}

	return conf
}
