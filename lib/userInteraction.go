package lib

import (
	"bufio"
	"fmt"
	"os"
)

// GreetUser on application start
func GreetUser() {
	fmt.Println(messageGreeting)
	fmt.Println()
}

// UserActionPage handles users actions regard HTTP pages
func UserActionPage() {
	userInput := ""
	userInputInt := -1

	for userInputInt < 0 || userInputInt > len(InputOptionsMap)-1 {
		fmt.Printf(messageMakeSelection)
		showUserInputOptions()
		userInput = getUserInputOption()
		userInputInt = stringToInt(userInput) - 1
	}

	callUserActionPageFunction(userInputInt)
}

func showUserInputOptions() {
	for counter := range InputOptionsMap {
		fmt.Printf("%v. %v\n", counter+1, InputOptionsMap[counter])
	}
	fmt.Println()
}

func getUserInputOption() string {
	reader := bufio.NewReader(os.Stdin)
	userInputString, _ := reader.ReadString('\n')
	userInputStringCleaned := removeCharacters(userInputString, "\n")

	return userInputStringCleaned
}

func callUserActionPageFunction(userOption int) {

	userOptionCommand := InputOptionsMap[userOption]
	switch userOptionCommand {
	case pageActionSaveLinks:
	case pageActionSavePage:
	case pageActionSaveText:
	case actionExit:
		ExitRequested = true

	}
	fmt.Println(userOptionCommand)
}
