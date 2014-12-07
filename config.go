package config

import (
	"os"
)

var (
	mode string
)

func init() {
	mode = os.Getenv("PF_MODE")
	if mode == "" {
		mode = "development"
	}
}

func Mode() string {
	return mode
}
