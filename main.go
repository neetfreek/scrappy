package main

import (
	"http-testing/lib"
)

func main() {
	lib.SetupConstants()
	lib.GreetUser()
	appLoop()
}

func appLoop() {
	for !lib.ExitRequested {
		lib.MenuMain()
	}
}
