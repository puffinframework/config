package config_test

import (
	"github.com/puffinframework/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMode(t *testing.T) {
	assert.Panics(t, func() { config.MustGetMode() })

	os.Setenv("PF_MODE", "something")
	assert.Panics(t, func() { config.MustGetMode() })

	os.Setenv("PF_MODE", "test")
	assert.Equal(t, "test", config.MustGetMode())
}
