package lib

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

func pageResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp
}

func pageHeader(url string) http.Header {
	resp := pageResponse(url)
	defer resp.Body.Close()
	pageResponseHeader := resp.Header

	return pageResponseHeader
}

func pageBody(url string) []byte {
	resp := pageResponse(url)
	defer resp.Body.Close()
	pageResponseData, _ := ioutil.ReadAll(resp.Body)

	return pageResponseData
}

// pageBodyString returns page body string from response
func pageBodyString(url string) string {
	p := bluemonday.UGCPolicy()

	pageBody := pageBody(url)
	buffer := bytesBuffer(pageBody)

	pageResponseString := buffer.String()

	pageBodyStringSanitised := p.Sanitize(pageResponseString)

	return pageBodyStringSanitised
}
