package main

import "http-testing/lib"

func main() {
	// url := "https://news.ycombinator.com/item?id=22213210"
	url := "https://www.neetfreek.net"
	// lib.WriteToFile("test", lib.PageString(url))
	// lib.WriteToFile("test", lib.PageString(url))
	lib.Tester(url)
}
