package lib

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetPage gets different content from provided url depending on action
func GetPage(url, action string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v, %v\n", messageNoPage, url)
		return
	}
	defer resp.Body.Close()

	if action == pageActionSavePage {
		getPageHTML(url)
	} else {
		loopGetPage(resp.Body, action)
	}
}

func loopGetPage(body io.Reader, action string) {
	responseToken := html.NewTokenizer(body)
	tagCurrent := ""

	for {
		responseTokenNext := responseToken.Next()
		switch responseTokenNext {
		case html.ErrorToken:
			return

		default:
			token := responseToken.Token()
			tag := token.Data
			if responseTokenNext == html.StartTagToken {
				tagCurrent = tag
				if action == pageActionSaveImages || action == pageActionSaveLinks {
					getPageImagesOrLinks(token, tag, action)
				}
			}
			if responseTokenNext == html.TextToken && action == pageActionSaveText {
				getPageText(token, tagCurrent)
				tagCurrent = ""
			}
		}
	}
}

func getPageText(token html.Token, tag string) {
	tagCurrent := tag
	tagCurrentIdentified := identifyTag(tagCurrent)
	text := strings.TrimSpace(token.String())
	textTrimmed := strings.TrimSpace(text)

	if len(textTrimmed) > 0 &&
		tagCurrentIdentified == HTMLTags.tagParagrah ||
		tagCurrentIdentified == HTMLTags.tagEmphasised ||
		tagCurrentIdentified == HTMLTags.tagAlternativeVoice ||
		tagCurrentIdentified == HTMLTags.tagDelete ||
		tagCurrentIdentified == HTMLTags.tagHeader1 ||
		tagCurrentIdentified == HTMLTags.tagHeader2 ||
		tagCurrentIdentified == HTMLTags.tagHeader3 ||
		tagCurrentIdentified == HTMLTags.tagHeader4 ||
		tagCurrentIdentified == HTMLTags.tagHeader5 ||
		tagCurrentIdentified == HTMLTags.tagHeader6 {
		fmt.Printf("%v \n\n", textTrimmed)
	}
}

func getPageImagesOrLinks(token html.Token, tag, action string) {
	if identifyTag(tag) == HTMLTags.tagHyperLink {
		attributesToSplit := token.String()
		attributesSplit := strings.Split(attributesToSplit, " ")

		for _, attr := range attributesSplit {
			if strings.Contains(attr, "src") || strings.Contains(attr, "href") && attributeContainsImage(attr) && action == pageActionSaveImages {
				fmt.Println("IMAGE: ", attr)
			} else if strings.Contains(attr, "href") && action == pageActionSaveLinks {
				fmt.Println("LINK: ", attr)
			}
		}
	}
}

func getPageHTML(url string) {
	writePageToFile(url)
}
