package entity

import (
	parser "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	Database string `env:"DATABASE"`
	Port     string `env:"MYSQL_PORT" envDefault:"3306"`
	Host     string `env:"MYSQL_HOST"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}

func (env *EnvironmentVariables) Load() {
	env.setWithEnvFile()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	parser.Parse(env)
}
