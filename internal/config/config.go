package config

import (
	"errors"
	"os"
)

type Config struct {
	S3Backet string
	S3Endpoint   string
	S3AuthToken    string
}

const (

)

var environments map[string]string = {
	""
}

func NewConfig() (*Config, error) {
	config := &Config{}

	var ok bool

	envName := "S3_BACKET"
	if config.S3Backet, ok = os.LookupEnv(envName); !ok {
		return nil, errors.New(envName)
	}

	envName = "S3_ENDPOINT"
	if config.S3Endpoint, ok = os.LookupEnv(envName); !ok {
		return nil, errors.New(envName)
	}

	envName = "S3_AUTH_TOKEN"
	if config.S3AuthToken, ok = os.LookupEnv(envName); !ok {
		return nil, errors.New(envName)
	}
}
