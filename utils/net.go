package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetBinary(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		// TODO: retry
		return []byte{}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: retry
		return []byte{}
	}

	return body
}

func GetText(url string) string {
	return strings.Trim(string(GetBinary(url)), " \r\n")
}
