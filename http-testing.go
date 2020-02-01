package main

import (
	"fmt"
	"http-testing/lib/requests"
	"http-testing/lib/utils"
)

func main() {
	printPageBody("https://news.ycombinator.com/item?id=22201337")
}

func printPageBody(url string) {

	pageResponseData := requests.GetPageResponseData(url)
	buffer := utils.BytesBuffer(pageResponseData)

	responseString := buffer.String()
	fmt.Print(responseString)
}
