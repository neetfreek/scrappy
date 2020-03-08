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

func addToLinksToDoNext(pageLinks []string, pageDomain string) []string {
	for _, link := range pageLinks {
		link = strings.Trim(link, "/")
		if strings.Contains(link, pageDomain) &&
			itemInSlice(link, linksCurrentMaster) == false &&
			itemInSlice(link, linksNextMaster) == false &&
			itemInSlice(link, linksDoneMaster) == false {
			linksNextMaster = append(linksNextMaster, link)
		}
	}
	return linksNextMaster
}

func itemInSlice(item string, slice []string) bool {
	for _, itemSlice := range slice {
		if itemSlice == item {
			return true
		}
	}
	return false
}

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

func printCollection(collection []string, name string) {
	fmt.Printf("------------------%v------------------\n", name)
	for counter, link := range collection {
		fmt.Printf("%v: %v\n", counter, link)
	}
}
