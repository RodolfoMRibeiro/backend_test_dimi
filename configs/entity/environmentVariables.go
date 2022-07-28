package entity

import (
	parser "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	Database   string `env:"DATABASE"`
	DbPort     string `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbConfigs  string `env:"DB_CONFIGS"`
	ServerPort string `env:"SERVER_PORT"`
	ServerHost string `env:"SERVER_HOST"`
}

func (env *EnvironmentVariables) FeedStruct() {
	env.setWithEnvFile()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	parser.Parse(env)
}
