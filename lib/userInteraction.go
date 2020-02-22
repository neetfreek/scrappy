package lib

import (
	"bufio"
	"fmt"
	"os"
)

// GreetUser on application start
func GreetUser() {
	fmt.Println(messageGreeting)
}

// MenuMain represents top-level menu for user application options
func MenuMain() {
	option := userMenuInput(InputOptionsMapMain)
	startActionMain(option)
}

func menuSite() {
	option := userMenuInput(InputOptionsMapSite)
	startActionSite(option)
}

func menuPage() {
	option := userMenuInput(InputOptionsMapPage)
	startActionPage(option)
}

func userMenuInput(menuMap []string) string {
	option := ""
	userInput := ""
	userInputInt := -1

	for option == "" {
		for userInputInt < 0 || userInputInt > len(menuMap)-1 {
			printMenuOptions(menuMap)
		fmt.Print(messageChoosePageAction)
		userInput = getUserInputOption()
		userInputInt = stringToInt(userInput) - 1
	}
		option = menuMap[userInputInt]
	}

	return option
}

func printMenuOptions(menuMap []string) {
	fmt.Println()
	for counter := range menuMap {
		fmt.Printf("%v. %v\n", counter+1, menuMap[counter])
	}
	fmt.Println()
}

func startActionMain(option string) {
	switch option {
	case actionPageContent:
		menuPage()
		break
	case actionSiteContent:
		menuSite()
		break
	case actionExit:
		ExitRequested = true
		break
	}
}

func startActionSite(option string) {
	if option == actionMenuMain {
		MenuMain()
	} else {
	}
}

func startActionPage(option string) {
	if option == actionMenuMain {
		MenuMain()
	} else {
		pageURL := getUserPageOption()
		getPageContent(pageURL, option)
	}
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
