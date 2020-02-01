package requests

import (
	utils "http-testing/lib/utils"
	"io/ioutil"
	"net/http"
)

// GetPageResponseData returns page body byte array from response
func GetPageResponseData(url string) []byte {
	resp, err := http.Get(url)
	utils.HandleError(err)
	defer resp.Body.Close()

	pageResponseData, err := ioutil.ReadAll(resp.Body)
	utils.HandleError(err)

	return pageResponseData
}
