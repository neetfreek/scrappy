package lib

/*==================================================================================*
* Examine pages, sites and retrieve data from multiple pages						*
*===================================================================================*
* The crawlSite routine goes through link and domain URL looking for link URLs to	*
*	other pages within the domain.													*
*	1) Link URLs are put into linksToDoCurrent, from where each are in turn crawled	*
*		for link URLs they contain.													*
*	2) URLS not alread in any lists are added to linksToDoNext.						*
*	3) Once crawled, URLS are placed into linksDone.								*
*	4) Once all linksToDoCurrent URLs have been crawled, the list's contents are	*
*		replaced by those of linksToDoNext, and linksToDoNext is cleared.			*
*	5) The process is complete when there are no more links from linksToDoNext to	*
*		put into linksToDoCurrent.													*
*===================================================================================*/

import (
	"fmt"
	"strings"
	"sync"
)

var linksToDoCurrent = []string{}
var linksToDoNext = []string{}
var linksDone = []string{}
var domain = ""

func crawlSite(url string) {

	domain = pageDomain(url)
	linksToDoCurrent = append(linksToDoCurrent, strings.Trim(url, "/"))
	if url != domain {
		linksToDoNext = append(linksToDoNext, strings.Trim(domain, "/"))
	}

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	for len(linksToDoCurrent) > 0 {
		for _, link := range linksToDoCurrent {
			waitGroup.Add(1)
			go crawlPageForLinks(&waitGroup, &mutex, link)
		}
		// crawlPageForLinks() // OLD IMPLEMENTATION
	}
	waitGroup.Wait()
	printCollection(linksDone, "DONE:")
}

func crawlPageForLinks(waitGroup *sync.WaitGroup, mutex *sync.Mutex, link string) {
	// func crawlPageForLinks() { // OLD IMPLEMENTATION
	var linksToAddToDone = []string{}
	var linksToAddToLinksToDoNext = []string{}
	var pageLinks = []string{}

	// for _, link := range linksToDoCurrent {
	// pageLinks = getPageLinks(link)
	// printCollection(pageLinks, "PAGE LINKS...")
	// resp := pageResponse(linksToDoCurrent[counter])
	// defer resp.Body.Close()

	resp := pageResponse(link)
	defer resp.Body.Close()
	pageLinks = loopGetPage(resp.Body, pageActionSaveLinks)

	// pageLinks = loopGetPage(resp.Body, pageActionSaveLinks)
	linksToAddToDone = append(linksToAddToDone, link)
	linksToAddToLinksToDoNext = addToLinksToDoNext(pageLinks, domain)
	// printCollection(linksToAddToLinksToDoNext, "LINKS TO ADD  EXT")
	// linksToDoNext = addToLinksToDoNext(pageLinks, domain) // OLD IMPLEMENTATION
	// linksDone = append(linksDone, link) // OLD IMPLEMENTATION
	// }
	mutex.Lock()
	linksDone = copyItemsToSlice(linksToAddToDone, linksDone)
	printCollection(linksDone, "LINKS DONE SO FAR")
	linksToDoCurrent = nil
	// linksToDoCurrent = copyItemsToSlice(linksToDoNext, linksToDoCurrent) // OLD IMPLEMENTATION
	linksToDoCurrent = copyItemsToSlice(linksToAddToLinksToDoNext, linksToDoCurrent)
	// linksToDoNext = nil
	mutex.Unlock()
	waitGroup.Done()
	fmt.Printf("Found %v links so far...\n", len(linksDone))
}

func getPageLinks(link string) []string {
	resp := pageResponse(link)
	defer resp.Body.Close()

	return loopGetPage(resp.Body, pageActionSaveLinks)
}
