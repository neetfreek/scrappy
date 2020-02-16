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
	pageDataCollection := []string{}
	pageDataToWrite := ""
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v, %v\n", messageNoPage, url)
		return
	}
	defer resp.Body.Close()

	if action == pageActionSavePage {
		pageDataToWrite = getPageHTML(url)
		writePageContentsToFile(url, pageDataToWrite, action)
	} else {
		pageDataCollection = loopGetPage(resp.Body, action)
		pageDataToWrite := strings.Join(pageDataCollection, "\n")
		writePageContentsToFile(url, pageDataToWrite, action)
	}
}

func loopGetPage(body io.Reader, action string) []string {
	responseToken := html.NewTokenizer(body)
	pageDataCollection := []string{}
	pageData := ""
	tagCurrent := ""

	for {
		responseTokenNext := responseToken.Next()
		switch responseTokenNext {
		case html.ErrorToken:
			return pageDataCollection

		default:
			token := responseToken.Token()
			tag := token.Data
			if responseTokenNext == html.StartTagToken {
				tagCurrent = tag
				if action == pageActionSaveImageLinks || action == pageActionSaveLinks {
					pageData = getPageImagesOrLinks(token, tag, action)
					if pageData != "" {
						pageDataCollection = append(pageDataCollection, pageData)
					}
				}
			}
			if responseTokenNext == html.TextToken && action == pageActionSaveText {
				pageData = getPageText(token, tagCurrent)
				if pageData != "" {
					pageDataCollection = append(pageDataCollection, pageData)
				}
				tagCurrent = ""
			}
		}
	}
}

func getPageText(token html.Token, tag string) string {
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
		return textTrimmed
	}
	return ""
}

func getPageImagesOrLinks(token html.Token, tag, action string) string {
	if identifyTag(tag) == HTMLTags.tagHyperLink {
		attributesToSplit := token.String()
		attributesSplit := strings.Split(attributesToSplit, " ")

		for _, attr := range attributesSplit {
			if strings.Contains(attr, "src") || strings.Contains(attr, "href") && attributeContainsImage(attr) && action == pageActionSaveImageLinks {
				return attr
			} else if strings.Contains(attr, "href") && action == pageActionSaveLinks {
				return attr
			}
		}
	}
	return ""
}

func getPageHTML(url string) string {
	return pageString(url)
}
