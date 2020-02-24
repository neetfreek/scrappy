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
			itemInSlice(link, linksToDoCurrent) == false &&
			itemInSlice(link, linksToDoNext) == false &&
			itemInSlice(link, linksDone) == false {
			linksToDoNext = append(linksToDoNext, link)
		}
	}
	return linksToDoNext
}

func itemInSlice(item string, slice []string) bool {
	for _, itemSlice := range slice {
		if itemSlice == item {
			return true
		}
	}
	return false
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
