package lib

/*==================================================================================*
* Collection of helpers functions for page parsing									*
*===================================================================================*
* Helper functions divided into several sections:									*
*	1) Processing HTTP response body.												*
*	2) Returning appropriate directory, file names for saving data.					*
*	3) Writing data files to local storage in root folder.							*
*===================================================================================*/

import (
	"bytes"
	"log"
	"os"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

// Helpers for process body
func pageBodyString(pageBody []byte) string {
	p := bluemonday.UGCPolicy()
	buffer := bytesBuffer(pageBody)
	pageResponseString := buffer.String()
	pageBodyStringSanitised := p.Sanitize(pageResponseString)

	return pageBodyStringSanitised
}

func bytesBuffer(byteArray []byte) bytes.Buffer {
	var buffer bytes.Buffer
	for sliceByte := range byteArray {
		buffer.WriteString(string(byteArray[sliceByte]))
	}

	return buffer
}

// Helpers for get directory, file names
func getNameDirectoryFromURL(url, action string) string {
	nameSplit := strings.Split(url, doubleSlash)
	name := nameSplit[len(nameSplit)-1]
	if name == "" {
		name = pageUntitled + timeStamp()
	}
	suffix := getNameSuffixByAction(action)

	return name + suffix
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

func getNameFileFromURL(url string) string {
	indexStartNamePage := strings.LastIndex(url, slash) + 1
	namePage := url[indexStartNamePage:len(url)]
	// Omit backslash at end of url
	if indexStartNamePage == len(url) {
		url := url[0 : len(url)-1]
		indexStartNamePage = strings.LastIndex(url, slash) + 1
		namePage = url[indexStartNamePage:len(url)]
	}
	var nameTimeStamped []string
	nameTimeStamped = append(nameTimeStamped, namePage, underscore, timeStamp())
	if namePage == "" {
		nameTimeStamped = append(nameTimeStamped, pageUntitled, timeStamp())
	}
	return strings.Join(nameTimeStamped, "")
}

func timeStamp() string {
	now := time.Now()
	return now.Format("200601021504")
}

// Helpers for write data
func writePageContentsToFile(url, pageBodyString, action string) {
	nameDirectoryParent := getNameDirectoryFromURL(url, "")
	var nameDirectoryContent []string
	nameDirectoryContent = append(nameDirectoryContent, nameDirectoryParent, slash, getNameDirectoryFromURL(url, action))
	nameFile := getNameFileFromURL(url)
	writeToFile(fileInformation(url), strings.Join(nameDirectoryContent, ""), nameFile, pageBodyString)
}

func fileInformation(url string) string {
	var fileInformation []string
	fileInformation = append(fileInformation,
		pageURL, delimiter, url, newLine,
		pageAccessed, delimiter, timeStamp(), newLine, newLine)

	return strings.Join(fileInformation, "")
}

func writeToFile(url, nameDirectory, nameFile, content string) {
	os.MkdirAll(nameDirectory, os.ModePerm)
	var filePath []string
	filePath = append(filePath, thisDirectoryDelimiter, nameDirectory, slash, nameFile)
	file, err := os.Create(strings.Join(filePath, ""))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(url)
	file.WriteString(content)
}

func identifyTag(tag string) string {

	if tagName, ok := HTMLMap[tag]; ok {

		return tagName
	}
	return ""
}
