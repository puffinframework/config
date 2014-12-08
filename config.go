package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
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
	ErrEnvVarNotSet error = fmt.Errorf("config: %s environment variable is not correctly set", ENV_VAR_NAME)
	ErrReadConfig   error = fmt.Errorf("config: couldn't read config file")
)

func MustGetMode() string {
	mode := strings.ToLower(os.Getenv(ENV_VAR_NAME))
	switch mode {
	case MODE_PRODUCTION, MODE_DEVELOPMENT, MODE_TEST:
		// ok
	default:
		log.Panic(ErrEnvVarNotSet)
	}
	return mode
}

func MustReadConfig(config interface{}) {
	mode := MustGetMode()
	file := strings.Join([]string{mode, "toml"}, ".")
	if _, err := toml.DecodeFile(file, config); err != nil {
		log.Println(err)
		log.Panic(ErrReadConfig)
	}
}
