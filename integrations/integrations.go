package integrations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	config "transaction/configs"
	"transaction/util"
)

type TransactionStatus struct {
	Authorization bool `json:"authorization"`
}

func (b *TransactionStatus) ConnectWithExternalAPI() {
	response, err := http.Get(config.API.URL)
	util.PresentateErros(err)

	responseData, err := ioutil.ReadAll(response.Body)
	util.PresentateErros(err)

	b.feedTransactionStruct(string(responseData))
}

func (b *TransactionStatus) feedTransactionStruct(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}
