package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// GetPageContent gets different content from provided url depending on action
func GetPageContent(url, action string) {
	pageDataCollection := []string{}
	pageDataToWrite := ""
	resp := pageResponse(url)

	statusCodeFirstNumber := string((strconv.Itoa(resp.StatusCode)[0]))
	if statusCodeFirstNumber != "2" {
		fmt.Printf("Page returned %v - no content returned.\n", resp.StatusCode)
		return
	}

	if action == pageActionSavePage {
		pageDataToWrite := getPageHTML(resp)
		if len(pageDataToWrite) > 0 {
			writePageContentsToFile(url, pageDataToWrite, action)
		}
	} else {
		pageDataCollection = loopGetPage(resp.Body, action)
		pageDataToWrite = strings.Join(pageDataCollection, "\n")
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
			} else if strings.Contains(attr, "href") && attributeContainsLink(attr) && action == pageActionSaveLinks {
				return attr
			}
		}
	}
	return ""
}

func getPageHTML(resp *http.Response) string {
	pageResponseData, _ := ioutil.ReadAll(resp.Body)
	pageBodyString := pageBodyString(pageResponseData)
	return pageBodyString
}
