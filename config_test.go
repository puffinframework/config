package config_test

import (
	"github.com/puffinframework/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMode(t *testing.T) {
	assert.Equal(t, "development", config.GetMode())
}
