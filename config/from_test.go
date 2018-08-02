package config

import (
	"testing"

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
