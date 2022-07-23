package integrations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"transaction/util"
)

type TransactionStatus struct {
	Authorization bool `json:"authorization"`
}

func (b *TransactionStatus) feedTransactionsStruct(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}

func (b *TransactionStatus) ConnectWithExternalAPI() {

	// the link just wasn't passed as an argument 'cause it's a constant.
	response, err := http.Get(`https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31`)
	util.PresentateErros(err)

	responseData, err := ioutil.ReadAll(response.Body)
	util.PresentateErros(err)

	b.feedTransactionsStruct(string(responseData))
}
