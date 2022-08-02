package integrations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TransactionStatus struct {
	Authorization bool `json:"authorization"`
}

func (b *TransactionStatus) ConnectWithExternalAPI() {
	response, err := http.Get(`https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31`)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// util.HasError(err)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// util.HasError(err)

	b.feedTransactionStruct(string(responseData))
}

func (b *TransactionStatus) feedTransactionStruct(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}
