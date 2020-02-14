package lib

import "fmt"

// UserInteraction handles showing options for, and getting, user interactions
func UserInteraction() {
	showUserInputOptions()
	userInput := getUserInputOption()
	fmt.Printf("User input received: %v\n", userInput)
	callInteractionFunction(userInput)
}

func showInstruction() {
	fmt.Println("Make a scrappy selection:")
}

func showUserInputOptions() {
	for counter := range InputOptionsMap {
		fmt.Printf("%v. %v\n", counter+1, InputOptionsMap[counter])
	}
}

func getUserInputOption() int {
	userInput := 1
	fmt.Println("Get user input, for now set to 1")
	// Read user input as int, decrement by one, assign to var
	// if error, call showUserInputOptions
	// Return user input variable
	return userInput
}

func callInteractionFunction(userInput int) {
	fmt.Println("Call appropriate function for user input", userInput)
}
