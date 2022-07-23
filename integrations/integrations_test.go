package integrations

import (
	"testing"

	"github.com/go-playground/assert"
)

var tran TransactionStatus

func TestConnectWithExternalAPI(t *testing.T) {

	t.Run("verify if the func is able to catch transference status", func(t *testing.T) {

		tran.ConnectWithExternalAPI()
		recivedBrokenString := tran.Authorization

		expectedCleanString := true

		assert.Equal(t, expectedCleanString, recivedBrokenString)
	})
}
