package lib

import (
	"log"
	"net/http"

	"github.com/temoto/robotstxt"
)

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
