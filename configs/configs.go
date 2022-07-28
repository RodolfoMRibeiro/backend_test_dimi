package config

import (
	"encoding/json"
	"os"
	"reflect"
	"transaction/util"

	"github.com/joho/godotenv"
)

var (
	Mysql  MysqlConfig
	Server ServerConfig
)

type MysqlConfig struct {
	DB                 string `env:"MYSQL_DATABASE"`
	PORT               string `env:"MYSQL_PORT" envDefault:"3306"`
	HOST               string `env:"MYSQL_HOST"`
	USER               string `env:"MYSQL_USER"`
	PASSWORD           string `env:"MYSQL_PASSWORD"`
	ADDITIONAL_CONFIGS string `env:"ADDITIONAL_CONFIGS"`
}

type ServerConfig struct {
	PORT string `env:"SERVER_PORT"`
	HOST string `env:"SERVER_HOST"`
}

func Load() {
	loadStructWithEnvVars(&Mysql)
	loadStructWithEnvVars(&Server)
}

func loadStructWithEnvVars(configStructure interface{}) {
	reflectElement := reflect.ValueOf(configStructure).Elem()

	environments := make(map[string]string)
	loadEnvFile()

	elements := reflectElement.Type()
	for i := 0; i < elements.NumField(); i++ {
		field := elements.Field(i)
		fieldName, envVarName := field.Name, field.Tag.Get("env")

		envVarValue := os.Getenv(envVarName)
		environments[fieldName] = envVarValue
	}

	parsed := util.ParseMapToJson(environments)
	json.Unmarshal([]byte(parsed), &configStructure)
}

func loadEnvFile() {
	godotenv.Load(".env")
}
