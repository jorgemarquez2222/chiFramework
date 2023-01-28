package services

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) []byte {
	resp, _ := http.Get(url)
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return dataBytes
}
