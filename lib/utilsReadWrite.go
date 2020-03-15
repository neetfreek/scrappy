package lib

import "strings"

func sliceToString(slice []string) string {
	return strings.Join(slice, "\n")
}

func writeDataToFile(url, data, action string) {
	if len(data) > 0 {
		writePageContentsToFile(url, data, action)
	}
}
