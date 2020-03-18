package lib

/*==================================================================================*
* Collection of helpers functions for page parsing									*
*===================================================================================*
* Helper functions divided into several sections:									*
*	1) Processing HTTP response body.												*
*	2) Returning appropriate directory, file names for saving data.					*
*===================================================================================*/

import (
	"bytes"

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

func identifyTag(tag string) string {

	if tagName, ok := HTMLMap[tag]; ok {

		return tagName
	}
	return ""
}
