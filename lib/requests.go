package lib

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

func pageHeader(url string) http.Header {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	pageResponseHeader := resp.Header

	return pageResponseHeader
}

func pageBody(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	pageResponseData, _ := ioutil.ReadAll(resp.Body)

	return pageResponseData
}

// pageString returns page body string from response
func pageString(url string) string {
	p := bluemonday.UGCPolicy()

	pageResponseData := pageBody(url)
	buffer := bytesBuffer(pageResponseData)

	pageResponseString := buffer.String()

	pageResponseStringSanitized := p.Sanitize(pageResponseString)

	return pageResponseStringSanitized
}
