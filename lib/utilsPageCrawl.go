package lib

import (
	"fmt"
	"strings"
)

/*==================================================================================*
* Collection of helper functions for page crawling									*
*===================================================================================*
* Functions responsible for moving data between linksToDoCurrent, linksToDoNext, 	*
*	and linksDone collections for page crawling through pages and tracking what has	*
*	been crawled.																	*
*===================================================================================*/

func indexItem(item string, slice []string) int {
	for index, itemSlice := range slice {
		if itemSlice == item {
			return index
		}
	}
	return -1
}

// Remove item at index by moving to end of slice, returning slice of all items but last
func removeItemFromSlice(indexItem int, slice []string) []string {
	slice[len(slice)-1], slice[indexItem] = slice[indexItem], slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func copyItemsToSlice(src, dest []string) []string {
	slice := dest
	for _, item := range src {
		slice = append(slice, item)
	}
	return slice
}

func domainLinks(slice []string, domain string) []string {
	var sliceStripped = []string{}

	for _, link := range slice {
		if strings.Contains(link, domain) && !isImageLink(link) {
			sliceStripped = append(sliceStripped, link)
		}
	}

	return sliceStripped
}

func isImageLink(link string) bool {
	for _, imageFormat := range imageFormats {
		if strings.Contains(link, imageFormat) {
			return true
		}
	}
	return false
}

func resetSlices() {
	linksCurrent = []string{}
	linksDone = []string{}
	linksImages = []string{}
	linksInProgress = []string{}
	linksPageCurrent = []string{}
}

func printCollection(collection []string, name string) {
	fmt.Printf("------------------%v------------------\n", name)
	for counter, link := range collection {
		fmt.Printf("%v: %v\n", counter, link)
	}
}

func crawlOutput(url, userAction string) {

	switch userAction {
	case siteActionSaveLinks:
		writePageContentsToFile(url, convertStringSliceToString(linksDone), userAction)
	case siteActionSaveImageLinks:
		writePageContentsToFile(url, convertStringSliceToString(linksImages), userAction)
	}
}

func addToLinksImages(linksPageCurrent, linksImages []string) []string {
	updatedlinksImages := linksImages

	for _, item := range linksPageCurrent {
		if isImageLink(item) {
			updatedlinksImages = append(updatedlinksImages, item)
		}
	}
	return updatedlinksImages
}
