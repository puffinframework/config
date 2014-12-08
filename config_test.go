package config_test

import (
	"github.com/puffinframework/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMode(t *testing.T) {
	assert.Panics(t, func() { config.MustGetMode() })

	os.Setenv(config.ENV_VAR_NAME, "something")
	assert.Panics(t, func() { config.MustGetMode() })

	os.Setenv(config.ENV_VAR_NAME, config.MODE_TEST)
	assert.Equal(t, config.MODE_TEST, config.MustGetMode())

	os.Setenv(config.ENV_VAR_NAME, "devELOPmenT")
	assert.Equal(t, config.MODE_DEVELOPMENT, config.MustGetMode())
}

func TestReadConfig(t *testing.T) {
	os.Setenv(config.ENV_VAR_NAME, config.MODE_TEST)

	type A struct {
		A1 string
		A2 string
	}

	type B struct {
		Bs []string
	}

	type C struct {
		D struct {
			Cd1 string
		}
	}

	type configAB struct {
		A A
		B B
	}

	cfg1 := &configAB{}
	config.MustReadConfig(cfg1)
	assert.Equal(t, "a1", cfg1.A.A1)
	assert.Equal(t, []string{"b1", "b2"}, cfg1.B.Bs)

	type configAC struct {
		A A
		C C
	}

	cfg2 := &configAC{}
	config.MustReadConfig(cfg2)
	assert.Equal(t, "a1", cfg2.A.A1)
	assert.Equal(t, "cd1", cfg2.C.D.Cd1)
}
