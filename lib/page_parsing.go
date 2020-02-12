package lib

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// PrintAttribute for testing print http tag attributes to standard output
func PrintAttribute(url string) {
	// Get response
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Make response token
	responseToken := html.NewTokenizer(resp.Body)
	tagTypeCurrent := ""

	for {
		responsToken := responseToken.Next()

		switch responsToken {
		case html.ErrorToken:
			return
		//For working with tags, e.g. inspecting HTML element tag values or attributes
		case html.StartTagToken:
			tag := responseToken.Token()
			tagString := tag.String()
			tagType := tag.Data
			tagTypeCurrent = tagType

			// Print start tags' tagHyperLink attributes containing http
			if (IdentifyTag(tagType)) == HTMLMap[HTMLTags.tagHyperLink] &&
				strings.Contains(tagString, "http") {
				for _, attribute := range tag.Attr {
					fmt.Printf("LINK: %v\n\n", attribute.Val)
					break
				}
			}
		// For working with text content of HTML elements
		case html.TextToken:
			text := strings.TrimSpace(responseToken.Token().String())
			textTrimmed := strings.TrimSpace(text)
			// Print non-empty strings
			if (IdentifyTag(tagTypeCurrent)) == HTMLTags.tagParagrah &&
				len(textTrimmed) > 0 {
				fmt.Printf("%v \n\n", textTrimmed)
				tagTypeCurrent = ""
			}
		}
	}
}
