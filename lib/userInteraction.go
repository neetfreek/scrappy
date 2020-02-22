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
		showUserInputOptions()
		fmt.Print(messageChoosePageAction)
		userInput = getUserInputOption()
		userInputInt = stringToInt(userInput) - 1
	}
	option := InputOptionsMap[userInputInt]

	if option == messageActionExit {
		ExitRequested = true
	} else {
		pageURL := getUserPageOption()
		getPageContent(pageURL, option)
	}

}

func showUserInputOptions() {
	fmt.Println()
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

func getUserPageOption() string {
	fmt.Print(messageChoosePageURL)
	pageURL := getUserInputOption()

	return pageURL
}
