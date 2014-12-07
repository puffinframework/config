package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	ENV_VAR_NAME     string = "PF_MODE"
	MODE_PRODUCTION  string = "production"
	MODE_DEVELOPMENT string = "development"
	MODE_TEST        string = "test"
)

var (
	ErrEnvVartNotSet error = fmt.Errorf("config: %s environment variable is not set", ENV_VAR_NAME)
)

func MustGetMode() string {
	mode := strings.ToLower(os.Getenv(ENV_VAR_NAME))
	switch mode {
	case MODE_PRODUCTION, MODE_DEVELOPMENT, MODE_TEST:
		return mode
	}
	panic(ErrEnvVartNotSet)
}
