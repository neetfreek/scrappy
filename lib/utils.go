package lib

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func bytesBuffer(responseData []byte) bytes.Buffer {
	var buffer bytes.Buffer
	for sliceByte := range responseData {
		buffer.WriteString(string(responseData[sliceByte]))
	}

	return buffer
}

func writePageToFile(url string) {
	pageString := pageString(url)
	WriteToFile(fileName, pageString)
}

func writeToFile(nameDirectory, nameFile, content string) {
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(content)
}

func getNameFromURL(url, fileOrDir string) string {
	matchStringDelimited := strings.Split(matchString, "/")
	fileName := matchStringDelimited[2]
	if fileName == "" {
		now := time.Now()
		fileName = "UntitledFile_" + now.Format("200601021504")
	}
	fileNameWithSuffix := fileName + ".html"

	return fileNameWithSuffix
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
