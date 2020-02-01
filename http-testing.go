package main

import (
	requests "http-testing/lib"
	utils "http-testing/utils"
)

func main() {
	url := "https://news.ycombinator.com/item?id=22201337"
	utils.WriteToFile("test", requests.GetPageResponseString(url))
}
