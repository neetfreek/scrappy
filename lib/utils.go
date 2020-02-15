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
	nameDirectory := getNameFromURL(url, typeDirectory)
	nameFile := getNameFromURL(url, typeFile)
	pageString := pageString(url)
	writeToFile(nameDirectory, nameFile, pageString)
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

func getNameFromURL(url, fileOrDir string) string {
	nameByHost := regexp.MustCompile("(.*)[/]")
	matchNameByHost := nameByHost.FindStringSubmatch(url)
	matchString := matchNameByHost[0]
	matchStringDelimited := strings.Split(matchString, "/")
	name := ""

	if fileOrDir == typeDirectory {
		name = matchStringDelimited[2]
	} else {
		name = matchStringDelimited[len(matchStringDelimited)-2] + "_" + timeStamp()
	}

	if name == "" {
		name = "Untitled_" + timeStamp()
	}
	return name
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
