package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//GetDollarValue API Client to request dollar value.
func GetDollarValue() (string, bool) {

	apiDollarToken := os.Getenv("API_DOLLAR_TOKEN")
	apiDollarUrl := os.Getenv("API_DOLLAR_URL")
	urlRequest := fmt.Sprintf("%v?apikey=%s&format=json", apiDollarUrl, apiDollarToken)

	response, err := http.Get(urlRequest)
	if err == nil {
		responseBody, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(responseBody), true
	}
	return "", false

}
