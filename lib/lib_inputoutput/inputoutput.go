package inputoutput

import (
	"fmt"
	requests "http-testing/lib/lib_requests"
	"os"
)

// PrintPageBody prints the url's page's body to standard output
func PrintPageBody(url string) {
	fmt.Print(requests.GetPageResponseString(url))
}

// WriteToFile writes the url's page's body to standard output
func WriteToFile(nameFile, content string) {
	file, err := os.Create(nameFile + ".html")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(content)
}
