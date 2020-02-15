package lib

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

func pageData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	pageResponseData, err := ioutil.ReadAll(resp.Body)

	return pageResponseData
}

// pageString returns page body string from response
func pageString(url string) string {
	p := bluemonday.UGCPolicy()

	pageResponseData := pageData(url)
	buffer := BytesBuffer(pageResponseData)

	pageResponseString := buffer.String()

	pageResponseStringSanitized := p.Sanitize(pageResponseString)

	return pageResponseStringSanitized
}
