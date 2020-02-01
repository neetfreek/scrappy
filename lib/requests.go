package requests

import (
	utils "http-testing/utils"
	"io/ioutil"
	"log"
	"net/http"
)

// GetPageResponseData returns page body byte array from response
func GetPageResponseData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	pageResponseData, err := ioutil.ReadAll(resp.Body)

	return pageResponseData
}

// GetPageResponseString returns page body string from response
func GetPageResponseString(url string) string {

	pageResponseData := GetPageResponseData(url)
	buffer := utils.BytesBuffer(pageResponseData)

	pageResponseString := buffer.String()
	return pageResponseString
}
