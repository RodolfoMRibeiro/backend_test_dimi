package configs

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
			Database:   os.Getenv("DATABASE"),
			DbPort:     os.Getenv("DB_PORT"),
			DbHost:     os.Getenv("DB_HOST"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbConfigs:  os.Getenv("DB_CONFIGS"),
		}

		assert.Equal(t, expectedStruct, recivedStruct)
	})

	t.Run("set struct with Environment Variables", func(t *testing.T) {
		godotenv.Load("../../.env")

		recivedStruct := EnvironmentVariables{}
		recivedStruct.FeedStruct()

		expectedStruct := EnvironmentVariables{
			Database:   os.Getenv("DATABASE"),
			DbPort:     os.Getenv("DB_PORT"),
			DbHost:     os.Getenv("DB_HOST"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbConfigs:  os.Getenv("DB_CONFIGS"),
		}

		assert.Equal(t, expectedStruct, recivedStruct)
	})
}
