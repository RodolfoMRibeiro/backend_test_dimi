package configs

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func LoadStructWithEnvVars(configStructure interface{}) {
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

	parsed := parseMapToJson(environments)
	json.Unmarshal([]byte(parsed), &configStructure)
}

func loadEnvFile() {
	godotenv.Load(".env")
}

func parseMapToJson(mp map[string]string) string {
	str, _ := json.Marshal(mp)
	return string(str)
}
