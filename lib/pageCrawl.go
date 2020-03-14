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
	"time"
)

var linksCurrent = []string{}
var linksDone = []string{}
var linksNext = []string{}
var linksPageCurrent = []string{}
var linksInProgress = []string{}

func crawlSite(url string) {
	resetSlices()
	defer Timer(time.Now(), "Total time:")
	// Setup domain, start pages
	var domain = pageDomainName(url)
	linksCurrent = append(linksCurrent, strings.Trim(url, "/"))
	if url != domain {
		linksNext = append(linksCurrent, strings.Trim(domain, "/"))
	}
	// Setup sync helpers
	var wg sync.WaitGroup
	var m sync.Mutex
	// Crawl all pages' links
	for len(linksCurrent) > 0 {
		for _, link := range linksCurrent {
			if indexItem(link, linksInProgress) == -1 {
				wg.Add(1)
				linksInProgress = append(linksInProgress, link)
				go crawlPageForLinks(&wg, &m, link, domain)
				fmt.Printf("Crawled: %v\n", link)
			}
		}
	}
	printCollection(linksDone, "Links found from "+domain)
	defer wg.Wait()
}

func crawlPageForLinks(wg *sync.WaitGroup, m *sync.Mutex, link, domain string) {
	defer wg.Done()
	// get page links
	resp := pageResponse(link)
	defer resp.Body.Close()
	linksPageCurrent = loopGetPage(resp.Body, pageActionSaveLinks)
	linksPageCurrent = domainLinks(linksPageCurrent, domain)
	if indexItem(link, linksCurrent) != -1 && indexItem(link, linksDone) == -1 {

		// update lists of links
		m.Lock()
		linksDone = append(linksDone, link)
		for _, item := range linksPageCurrent {
			if indexItem(item, linksCurrent) == -1 &&
				indexItem(item, linksDone) == -1 {
				linksCurrent = append(linksCurrent, item)
			}
		}
		linksCurrent = removeItemFromSlice(indexItem(link, linksCurrent), linksCurrent)
		linksInProgress = removeItemFromSlice(indexItem(link, linksInProgress), linksInProgress)
		m.Unlock()
	}
}
