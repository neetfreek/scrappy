package lib

/*==================================================================================*
* Interaction with user, including printing menus, and getting, processing input	*
*===================================================================================*
* MenuMain represents entry-point into menu routine. Options printed represent		*
*	global CONSTANT variables stored in slices.										*
* User input selections are represented by the same CONSTANT variables presented	*
*	in menu printing.																*
*===================================================================================*/

import (
	"bufio"
	"fmt"
	"os"
)

// GreetUser on application start
func GreetUser() {
	fmt.Println(messageGreeting)
}

// Menu functions

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

func printMenuOptions(menuMap []string) {
	fmt.Println()
	for counter := range menuMap {
		fmt.Printf("%v. %v\n", counter+1, menuMap[counter])
	}
	fmt.Println()
}

// User input functions

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

func getUserInputOption() string {
	reader := bufio.NewReader(os.Stdin)
	userInputString, _ := reader.ReadString('\n')
	userInputStringCleaned := removeCharacters(userInputString, "\n")

	return userInputStringCleaned
}

func getUserPageOption() string {
	fmt.Print(messageChoosePageURL)
	pageURL := getUserInputOption()

func getDomainScope(url string) {
	scopesParts := []string{}
	scopes := []string{}
	scopes = append(scopes, pageDomain(url))
	urlSplit := strings.Split(url, "/")

	for counter := 1; counter < len(urlSplit); counter++ {
		if len(urlSplit[counter]) > 0 {
			scopesParts = append(scopesParts, urlSplit[counter])
		}
	}

	for counter := 1; counter < len(scopesParts); counter++ {
		scopes = append(scopes, scopes[counter-1]+slash+scopesParts[counter])
	}
}
