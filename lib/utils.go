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

// BytesBuffer writes byte date into a returned buffer
func BytesBuffer(responseData []byte) bytes.Buffer {
	var buffer bytes.Buffer
	for sliceByte := range responseData {
		buffer.WriteString(string(responseData[sliceByte]))
	}

	return buffer
}

// WritePageToFile writes HTML page contents to file
func WritePageToFile(url string) {
	fileName := getFileNameFromURL(url)
	pageString := pageString(url)
	WriteToFile(fileName, pageString)
}

// WriteToFile writes data to a file
func WriteToFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(content)
}

func getFileNameFromURL(url string) string {
	fileNameByHost := regexp.MustCompile("(.*)[/]")
	matchfileNameByHost := fileNameByHost.FindStringSubmatch(url)
	matchString := matchfileNameByHost[0]
	matchStringDelimited := strings.Split(matchString, "/")
	fileName := matchStringDelimited[2]
	if fileName == "" {
		now := time.Now()
		fileName = "UntitledFile_" + now.Format("200601021504")
	}
	fileNameWithSuffix := fileName + ".html"

	return fileNameWithSuffix
}

// IdentifyTag returns string constant value of tag
func IdentifyTag(tag string) string {
	if tagName, ok := HTMLMap[tag]; ok {
		return tagName
	}
	return ""
}

func removeCharacters(stringToStrip, charToRemove string) string {
	stringCleaned := strings.ReplaceAll(stringToStrip, charToRemove, "")

	return stringCleaned
}

func stringToInt(amountString string) int {
	stringToClean := amountString
	removeNonDigts := regexp.MustCompile("[^0-9]+")
	cleanedString := removeNonDigts.ReplaceAllString(stringToClean, "")
	amountInt, err := strconv.Atoi(cleanedString)
	if err != nil {
		log.Fatal(err)
	}

	return amountInt
}
