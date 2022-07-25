package entity

import (
	"os"
	"testing"

	"github.com/go-playground/assert"
	"github.com/joho/godotenv"
)

func TestEnvironmentVariables(t *testing.T) {
	t.Run("set struct with Environment Variables", func(t *testing.T) {
		godotenv.Load("../../.env")

		recivedStruct := EnvironmentVariables{}
		recivedStruct.setWithEnvFile()

		expectedStruct := EnvironmentVariables{
			Database: os.Getenv("DATABASE"),
			Port:     os.Getenv("DB_PORT"),
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Configs:  os.Getenv("DB_CONFIGS"),
		}

		assert.Equal(t, expectedStruct, recivedStruct)
	})
}
