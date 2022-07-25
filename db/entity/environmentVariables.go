package entity

import (
	parser "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	Database string `env:"db_DATABASE"`
	Port     string `env:"db_PORT" envDefault:"3306"`
	Host     string `env:"db_HOST"`
	User     string `env:"db_USER"`
	Password string `env:"db_PASSWORD"`
	Configs  string `env:"db_CONFIGS"`
}

func (env *EnvironmentVariables) FeedStruct() {
	env.setWithEnvFile()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	parser.Parse(env)
}
