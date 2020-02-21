package lib

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

func pageBodyString(pageBody []byte) string {
	p := bluemonday.UGCPolicy()
	buffer := bytesBuffer(pageBody)
	pageResponseString := buffer.String()
	pageBodyStringSanitised := p.Sanitize(pageResponseString)

	return pageBodyStringSanitised
}

func bytesBuffer(responseData []byte) bytes.Buffer {
	var buffer bytes.Buffer
	for sliceByte := range responseData {
		buffer.WriteString(string(responseData[sliceByte]))
	}

	return buffer
}

func writePageContentsToFile(url, pageBodyString, action string) {
	nameDirectoryParent := getNameDirectoryFromURL(url, "")
	nameDirectoryContent := nameDirectoryParent + "/" + getNameDirectoryFromURL(url, action)
	nameFile := getNameFileFromURL(url)
	writeToFile(nameDirectoryContent, nameFile, pageBodyString)
}

func writeToFile(nameDirectory, nameFile, content string) {
	os.MkdirAll(nameDirectory, os.ModePerm)
	file, err := os.Create("./" + nameDirectory + "/" + nameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(content)
}

func getNameDirectoryFromURL(url, action string) string {
	nameSplit := strings.Split(url, "//")
	name := nameSplit[len(nameSplit)-1]
	suffix := getNameSuffixByAction(action)
	if name == "" {
		name = "Untitled_" + timeStamp()
	}
	return name + suffix
}

func getNameFileFromURL(url string) string {
	indexStartNamePage := strings.LastIndex(url, "/") + 1
	namePage := url[indexStartNamePage:len(url)]
	// Omit backslash at end of url
	if indexStartNamePage == len(url) {
		url := url[0 : len(url)-1]
		indexStartNamePage = strings.LastIndex(url, "/") + 1
		namePage = url[indexStartNamePage:len(url)]
	}

	nameTimeStamped := namePage + "_" + timeStamp()
	if namePage == "" {
		nameTimeStamped = "Untitled_" + timeStamp()
	}
	return nameTimeStamped
}

func getNameSuffixByAction(action string) string {
	suffix := ""

	switch action {
	case pageActionSaveImageLinks:
		suffix = suffixImageLinks
	case pageActionSaveLinks:
		suffix = suffixLinks
	case pageActionSavePage:
		suffix = suffixHTML
	case pageActionSaveText:
		suffix = suffixText
	}
	return suffix
}

func identifyTag(tag string) string {

	if tagName, ok := HTMLMap[tag]; ok {

		return tagName
	}
	return ""
}

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

func removeCharacters(stringToStrip, charToRemove string) string {
	stringCleaned := strings.ReplaceAll(stringToStrip, charToRemove, "")

	return stringCleaned
}

func stringToInt(amountString string) int {
	stringToClean := amountString
	removeNonDigts := regexp.MustCompile("[^0-9]+")
	cleanedString := removeNonDigts.ReplaceAllString(stringToClean, "")
	amountInt, _ := strconv.Atoi(cleanedString)

	return amountInt
}

func timeStamp() string {
	now := time.Now()
	return now.Format("200601021504")
}
