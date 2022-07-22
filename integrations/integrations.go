package integrations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TransactionStatus struct {
	Approved bool `json:"approved"`
}

func (b *TransactionStatus) feedTransactionsStruct(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}

func (b *TransactionStatus) hfghfg() {

	response, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	ValidateErros(err)

	responseData, err := ioutil.ReadAll(response.Body)
	ValidateErros(err)

	feedTransactionsStruct(string(responseData))
}

func ValidateErros(err error) {
	// err != nil {
	// 	fmt.Println("Error: ", err)
	// }
}
