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
	assert.Equal(t, config.MODE_TEST, config.MustGetMode())

	os.Setenv("PF_MODE", "devELOPmenT")
	assert.Equal(t, config.MODE_DEVELOPMENT, config.MustGetMode())
}
