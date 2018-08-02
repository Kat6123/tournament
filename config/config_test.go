package config

import (
	"testing"

	"github.com/kat6123/tournament/log"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tt := []struct {
		name      string
		c1        Configuration
		c2        Configuration
		expectedC Configuration
	}{
		{
			name: "merge with empty config",
			c1: Configuration{
				DB: dbConfig{
					URL:              "localhost:27017",
					DB:               "tours",
					TourCollection:   "tours",
					PlayerCollection: "players",
				},
				Port:  "3001",
				Debug: log.TraceLevel,
			},
			c2: Configuration{},
			expectedC: Configuration{
				DB: dbConfig{
					URL:              "localhost:27017",
					DB:               "tours",
					TourCollection:   "tours",
					PlayerCollection: "players",
				},
				Port:  "3001",
				Debug: log.TraceLevel,
			},
		},
		{
			name: "merge with empty config",
			c1: Configuration{
				DB: dbConfig{
					URL:              "localhost:27017",
					DB:               "tours",
					TourCollection:   "tours",
					PlayerCollection: "players",
				},
				Port:  "3001",
				Debug: log.ErrorLevel,
			},
			c2: Configuration{
				DB: dbConfig{
					URL:              "another",
					DB:               "another",
					TourCollection:   "another",
					PlayerCollection: "another",
				},
				Port:  "10000",
				Debug: log.TraceLevel,
			},
			expectedC: Configuration{
				DB: dbConfig{
					URL:              "another",
					DB:               "another",
					TourCollection:   "another",
					PlayerCollection: "another",
				},
				Port:  "10000",
				Debug: log.TraceLevel,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// How I can test that c1 and c2 don't change?
			gotC := merge(&tc.c1, &tc.c2)

			assert.Equal(t, tc.expectedC, gotC)
		})
	}
}

func TestGet(t *testing.T) {
	*yamlPath = "config-test.yaml"
	expectedConfig := Configuration{
		DB: dbConfig{
			URI:              "mongodb://localhost/tours",
			TourCollection:   "tours",
			PlayerCollection: "players",
		},
		Port:  "3001",
		Debug: log.TraceLevel,
	}

	assert.Equal(t, expectedConfig, Get())
}
