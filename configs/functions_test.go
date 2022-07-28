package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeStruct struct {
	Name string `env:"NAME"`
	Age  string `env:"AGE"`
}

func TestConfigFunctions(t *testing.T) {
	t.Run("ParseMapToJson --> must parse a string map", func(t *testing.T) {
		testMap := map[string]string{
			"rodolfo": "handsome",
			"skill":   "nice",
		}

		expectedJsonElement := `{"rodolfo":"handsome","skill":"nice"}`

		receivedJsonElement := parseMapToJson(testMap)

		assert.Equal(t, expectedJsonElement, receivedJsonElement)

	})

}
