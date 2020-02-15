package main

import (
	"http-testing/lib"
)

func main() {
	// url := "https://news.ycombinator.com/item?id=22213210"
	lib.SetupConstants()
	lib.GreetUser()
	appLoop()
}

func appLoop() {
	for {
		lib.UserActionPage()
	}
}
