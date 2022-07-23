package entity

import (
	parser "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	Database string `env:"DATABASE"`
	Port     string `env:"PORT" envDefault:"3306"`
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Configs  string `env:"CONFIGS"`
}

func (env *EnvironmentVariables) FeedStruct() {
	env.setWithEnvFile()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	parser.Parse(env)
}
