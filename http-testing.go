package main

import (
	lib "http-testing/lib"
)

func main() {
	url := "https://news.ycombinator.com/item?id=22201337"
	lib.WriteToFile("test", lib.GetPageResponseString(url))
}
