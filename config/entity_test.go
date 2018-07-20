package config

import (
	"testing"

	"gotest.tools/assert"
)

func TestDefault(t *testing.T) {
	d := Default()

	assert.Equal(t, defaultConfig, d)
	if &d == &defaultConfig {
		t.Fatal("pointers should differ")
	}
}
