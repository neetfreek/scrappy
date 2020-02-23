package lib

/*==================================================================================*
* Collection of helpers functions for user interaction								*
*===================================================================================*
* Functions relate to preparing user input for later use in other functions.		*
*===================================================================================*/

import (
	"regexp"
	"strconv"
	"strings"
)

func removeCharacters(stringToStrip, charToRemove string) string {
	stringCleaned := strings.ReplaceAll(stringToStrip, charToRemove, "")

	return stringCleaned
}

func stringToInt(amountString string) int {
	stringToClean := amountString
	removeNonDigts := regexp.MustCompile("[^0-9]+")
	cleanedString := removeNonDigts.ReplaceAllString(stringToClean, "")
	amountInt, _ := strconv.Atoi(cleanedString)

	return amountInt
}
