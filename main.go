package main

import (
	"http-testing/lib"
)

func main() {
	// url := "https://news.ycombinator.com/item?id=22213210"
	// url := "https://www.neetfreek.net/tales-of-arrissia"
	// url := "https://www.oldbookillustrations.com/illustrations/de-groof-falling/"
	lib.SetupConstants()
	lib.GreetUser()
	appLoop()
}

func appLoop() {
	for !lib.ExitRequested {
		lib.UserActionPage()
	}
}
