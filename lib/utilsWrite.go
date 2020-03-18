/*==================================================================================*
* Functions related to writing data from crawling to local storage					*
*===================================================================================*
* Functions are organised into two logically distinct groups:						*
*	1. Those responsible for setting appropriate file names and paths				*
*	2. Those reponsible for writing data to local storage							*
*===================================================================================*/

package lib

import (
	"log"
	"os"
	"strings"
	"time"
)

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
	case siteActionSaveImageLinks:
		suffix = suffixImageLinks
	case siteActionSaveLinks:
		suffix = suffixLinks
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

// Helpers for writing contents to local storage
func writePageContentsToFile(url, pageBodyString, userAction string) {
	nameDirectoryParent := getNameDirectoryFromURL(url, "")
	var nameDirectoryContent []string
	nameDirectoryContent = append(nameDirectoryContent, nameDirectoryParent, slash, getNameDirectoryFromURL(url, userAction))
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
