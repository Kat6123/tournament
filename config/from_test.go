package config

import (
	"testing"

	"os"

	"github.com/kat6123/tournament/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fromYAML(t *testing.T) {
	const (
		path           = "config-test.yaml"
		notExistedPath = "not_exist"
		badConfigPath  = "badconfig-test.yaml"
	)

	expectedConfig := &Configuration{
		DB: dbConfig{
			URI:              "mongodb://localhost/tours",
			TourCollection:   "tours",
			PlayerCollection: "players",
		},
		Port:  "3001",
		Debug: log.TraceLevel,
	}

	gotConfig, err := fromYAML(path)
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, gotConfig)

	_, err = fromYAML(notExistedPath)
	assert.EqualError(t, err, "read file: open not_exist: no such file or directory")

	_, err = fromYAML(badConfigPath)
	assert.EqualError(t, err, "unmarshal: yaml: line 5: could not find expected ':'")
}

func setEnv(env map[string]string) {
	for v := range env {
		os.Setenv(v, env[v])
	}
}

func Test_fromENV(t *testing.T) {
	tt := []struct {
		name           string
		env            map[string]string
		expectedConfig *Configuration
		expectedErr    string
	}{
		{
			name: "valid config",
			env: map[string]string{
				"DB_URI":    "mongodb://localhost/tours",
				"TOURS":     "tours",
				"PLAYERS":   "players",
				"PORT":      "8000",
				"LOG_LEVEL": "info",
			},
			expectedConfig: &Configuration{
				DB: dbConfig{
					URI:              "mongodb://localhost/tours",
					TourCollection:   "tours",
					PlayerCollection: "players",
				},
				Port:  "8000",
				Debug: log.InfoLevel,
			},
		},
		{
			name: "valid config",
			env: map[string]string{
				"DB_URI":    "mongodb://localhost/tours",
				"TOURS":     "tours",
				"PLAYERS":   "players",
				"PORT":      "8000",
				"LOG_LEVEL": "abra",
			},
			expectedErr: "config LOG_LEVEL from env: undefined log level: abra",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			setEnv(tc.env)
			gotConfig, err := fromEnv()

			if tc.expectedErr != "" {
				require.EqualError(t, err, tc.expectedErr)
			}

			assert.Equal(t, tc.expectedConfig, gotConfig)
		})
	}
}

func Test_fromFlags(t *testing.T) {
	// Set vars dbURI, tours and so on or set os.Args and then Parse?
	assert.Equal(t, &defaultConfig, fromFlags())

	//os.Args = []string{
	//	"program",
	//	"-uri", "db://",
	//	"-tours", "tours",
	//	"-players", "players",
	//	"-port", "8000",
	//	"-log-level", "info",
	//}

	*dbURI = "db://"
	*tours = "tours"
	*players = "players"
	*port = "8000"
	*debugLevel = log.InfoLevel

	expectedConfig := &Configuration{
		DB: dbConfig{
			URI:              "db://",
			TourCollection:   "tours",
			PlayerCollection: "players",
		},
		Port:  "8000",
		Debug: log.InfoLevel,
	}

	//flag.Parse() Package testing also calls flag.Parse
	assert.Equal(t, expectedConfig, fromFlags())
}
