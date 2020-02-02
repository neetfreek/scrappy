package main

import "http-testing/lib"

func main() {
	url := "https://news.ycombinator.com/item?id=22213210"
	lib.WriteToFile("test", lib.PageString(url))
	lib.GetNodeContent(lib.PageString(url), "p")
}
