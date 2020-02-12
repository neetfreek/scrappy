package main

import (
	"http-testing/lib"
)

func main() {
	// url := "https://news.ycombinator.com/item?id=22213210"
	url := "https://www.neetfreek.net"
	lib.SetupTags()
	lib.PrintAttribute(url)
}
