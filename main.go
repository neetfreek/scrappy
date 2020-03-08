package main

/*==================================================================================*
* Application entry point and loop							<(^^)><(^ ^)><( ^^)>	*
*===================================================================================*
* Application loop serves to have application run in a REPL manner - after options	*
*	input and pages, sites processed, the user is able to continue to work with 	*
*	other pages, sites.																*
*===================================================================================*/

import (
	"http-testing/lib"
	"time"
)

func main() {
	lib.SetupConstants()
	lib.GreetUser()
	appLoop()
}

func appLoop() {
	// lib.Test()
	defer lib.Timer(time.Now(), "total:")
	for !lib.ExitRequested {
		lib.MenuMain()
	}
}
