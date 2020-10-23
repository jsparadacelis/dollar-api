package main

import (
	"io/ioutil"
	"net/http"
)

//GetDollarValue API Client to request dollar value.
func GetDollarValue() (string, bool) {

	response, err := http.Get("https://api.exchangeratesapi.io/latest")
	if err == nil {
		responseBody, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(responseBody), true
	}
	return "", false

}
