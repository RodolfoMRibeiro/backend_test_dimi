package integrations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	config "transaction/configs"
)

type TransactionStatus struct {
	Authorization bool `json:"authorization"`
}

func (b *TransactionStatus) ConnectWithExternalAPI() {

	response, err := http.Get(config.API.URL)
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
