package utils

import (
	"bytes"
	"log"
	"os"
)

// HandleError handles global errors
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// BytesBuffer writes byte date into a returned buffer
func BytesBuffer(responseData []byte) bytes.Buffer {
	var buffer bytes.Buffer
	for sliceByte := range responseData {
		buffer.WriteString(string(responseData[sliceByte]))
	}
	return buffer
}

func writeToFile(nameFile, content string) {
	file, err := os.Create(nameFile + ".html")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(content)
}
