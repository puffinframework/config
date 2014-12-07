package config_test

import (
	"github.com/puffinframework/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMode(t *testing.T) {
	assert.Equal(t, "development", config.GetMode())
	os.Setenv("PF_MODE", "test")
	assert.Equal(t, "test", config.GetMode())
}
