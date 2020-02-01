package utils

import (
	"bytes"
	"log"
	"os"
)

// BytesBuffer writes byte date into a returned buffer
func BytesBuffer(responseData []byte) bytes.Buffer {
	var buffer bytes.Buffer
	for sliceByte := range responseData {
		buffer.WriteString(string(responseData[sliceByte]))
	}

	return buffer
}

// WriteToFile writes contents to file
func WriteToFile(nameFile, content string) {
	file, err := os.Create(nameFile + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(content)
}
