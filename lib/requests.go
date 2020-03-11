package lib

/*==================================================================================*
* HTTP-request and -response-related functionalities								*
*===================================================================================*
* Contains functionality directly related to HTTP requests, receiving HTTP			*
*	responses.																		*
* Contains other top-level functionality like determining the domain URL and parsing.	*
*	robots.txt files for crawler permissions and behaviour.							*
*===================================================================================*/

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/temoto/robotstxt"
)

func pageResponse(url string) *http.Response {
	// fmt.Printf("GETT RESPONSE FROM _%v_\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v %v", messageNotURLFormat, err)
		MenuMain()
	}

	return resp
}

func robotsTxt(url string) {
	var pageDomainStringArray = []string{}
	pageDomain := pageDomainURL(url)
	pageDomainStringArray = append(pageDomainStringArray, pageDomain, addressRobotsTxt)
	addressRobotsTxt := strings.Join(pageDomainStringArray, "")

	resp := pageResponse(addressRobotsTxt)
	defer resp.Body.Close()
	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	userAgentGroup := data.FindGroup(goClient)
	if userAgentGroup != nil {
		if userAgentGroup.Test("/all") {
			// crawlDelayInt64 := int64(userAgentGroup.CrawlDelay)
			// fmt.Printf("userAgentGroup.CrawlDelay: %d", crawlDelayInt64)
			// TODO: Define collection of UGs to test against
		}
	}
}

func pageDomainURL(url string) string {
	re := regexp.MustCompile(`([^@\/\n]+@)?([^:\/\n]+)`)
	pageDomain := re.FindAllStringSubmatch(url, -1)
	pageDomainStringArray := []string{}
	pageDomainStringArray = append(pageDomainStringArray, pageDomain[0][0], delimiterDomain, pageDomain[1][0])
	return strings.Join(pageDomainStringArray, "")
}

func pageDomainName(url string) string {
	reSlash := regexp.MustCompile(`([^@\/\n]+@)?([^:\/\n]+)`)
	pageDomain := reSlash.FindAllStringSubmatch(url, -1)
	pageName := strings.Split(pageDomain[1][0], ".")
	return pageName[0] + dot + pageName[1]
}
