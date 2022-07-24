package entity

import (
	parser "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	Database string `env:"DATABASE"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Configs  string `env:"DB_CONFIGS"`
}

func (env *EnvironmentVariables) FeedStruct() {
	env.setWithEnvFile()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	parser.Parse(env)
}
