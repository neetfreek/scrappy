package lib

import (
	"fmt"
	"strings"
)

var linksToDoCurrent = []string{}
var linksToDoNext = []string{}
var linksDone = []string{}
var host = ""

func crawlSite(url string) {

	host = pageDomain(url)
	linksToDoNext = append(linksToDoNext, host)
	linksToDoCurrent = append(linksToDoCurrent, url)

	crawPageForLinks()

	for counter, link := range linksDone {
		fmt.Printf("%v: %v\n", counter, link)
	}
}

func crawPageForLinks() {
	for counter, link := range linksToDoCurrent {
		resp := pageResponse(linksToDoCurrent[counter])
		defer resp.Body.Close()

		pageLinks := loopGetPage(resp.Body, pageActionSaveLinks)
		addToLinksToDoNext(pageLinks, host)
		linksDone = append(linksDone, link)
	}
	linksToDoCurrent = nil
	linksToDoCurrent = copyItemsToSlice(linksToDoNext, linksToDoCurrent)
	linksToDoNext = nil

	if len(linksToDoCurrent) > 0 {
		crawPageForLinks()
	}
}

func addToLinksToDoNext(pageLinks []string, pageDomain string) {
	for _, link := range pageLinks {
		if strings.Contains(link, pageDomain) &&
			itemInSlice(link, linksToDoCurrent) == false &&
			itemInSlice(link, linksToDoNext) == false &&
			itemInSlice(link, linksDone) == false {
			linksToDoNext = append(linksToDoNext, link)
		}
	}
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
