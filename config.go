package config

import (
	"os"
	"strings"
)

func GetMode() string {
	mode := os.Getenv("PF_MODE")
	if mode == "" {
		mode = "development"
	}
	return strings.ToLower(mode)
}

