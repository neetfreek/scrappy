package lib

import (
	"log"
	"net/http"
	"strings"

	"github.com/temoto/robotstxt"
)

func attributeContainsImage(attribute string) bool {
	for _, imageFormat := range imageFormats {
		if strings.Contains(attribute, imageFormat) {
			return true
		}
	}
	return false
}

func attributeContainsLink(attribute string) bool {
	if strings.Contains(attribute, "http") {
		return true
	}
	return false
}

func pageResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func robotsTxt(url string) {
	resp := pageResponse(url)
	defer resp.Body.Close()

	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	userAgentGroup := data.FindGroup(goClient)
	if userAgentGroup != nil {
		if userAgentGroup.Test("/all") {
			// TODO: Define collection of UGs to test against
		}
	}
}

func pageResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
