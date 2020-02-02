package lib

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

// GetPageResponseData returns page body byte array from response
func PageData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	pageResponseData, err := ioutil.ReadAll(resp.Body)

	return pageResponseData
}

// GetPageResponseString returns page body string from response
func PageString(url string) string {
	p := bluemonday.UGCPolicy()

	pageResponseData := PageData(url)
	buffer := BytesBuffer(pageResponseData)

	pageResponseString := buffer.String()

	pageResponseStringSanitized := p.Sanitize(pageResponseString)

	return pageResponseStringSanitized
}
