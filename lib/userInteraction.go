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

	for userInputInt < 0 || userInputInt > 3 {
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

func callInteractionFunction(userInput int) {
	fmt.Printf("Great. Just image we did option %v with a web page. Impressive!\n\n", userInput)
}
