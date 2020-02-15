package lib

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetPage for testing print http tag attributes to standard output
func GetPage(url, action string) {
	// Get response
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v, %v\n", messageNoPage, url)
		return
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

		case html.StartTagToken:
			tag := responseToken.Token()
			// tagString := tag.String()
			tagType := tag.Data
			tagTypeCurrent = tagType

			// Get links, images
			if IdentifyTag(tagType) == HTMLTags.tagHyperLink {
				attributesToSplit := tag.String()
				attributesSplit := strings.Split(attributesToSplit, " ")
				for _, attr := range attributesSplit {
					if strings.Contains(attr, "src") || strings.Contains(attr, "href") && attributeContainsImage(attr) {
						fmt.Println("IMAGE: ", attr)
					} else if strings.Contains(attr, "href") {
						fmt.Println("LINK: ", attr)
					}
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
