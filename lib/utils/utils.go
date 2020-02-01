package utils

import (
	"bytes"
	"log"
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
