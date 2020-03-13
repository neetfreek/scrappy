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

var linksCurrentMaster = []string{}
var linksDoneMaster = []string{}
var linksNextMaster = []string{}
var linksPageCurrent = []string{}
var linksInProgress = []string{}

func crawlSite(url string) {
	// Setup domain, start pages
	var domain = pageDomainName(url)
	linksCurrentMaster = append(linksCurrentMaster, strings.Trim(url, "/"))
	if url != domain {
		linksNextMaster = append(linksCurrentMaster, strings.Trim(domain, "/"))
	}
	// Setup sync helpers
	var wg sync.WaitGroup
	var m sync.Mutex
	// Crawl all pages' links
	for len(linksCurrentMaster) > 0 {
		for _, link := range linksCurrentMaster {
			if !itemInSlice(link, linksInProgress) {
				wg.Add(1)
				linksInProgress = append(linksInProgress, link)
				go crawlPageForLinks(&wg, &m, link, domain)
				fmt.Printf("Crawled: %v\n", link)
			}
		}
	}
	printCollection(linksDoneMaster, "Links found from "+domain)
	defer wg.Wait()
}

func crawlPageForLinks(wg *sync.WaitGroup, m *sync.Mutex, link, domain string) {
	defer wg.Done()
	// get page links
	resp := pageResponse(link)
	defer resp.Body.Close()
	linksPageCurrent = loopGetPage(resp.Body, pageActionSaveLinks)
	linksPageCurrent = domainLinks(linksPageCurrent, domain)
	if itemInSlice(link, linksCurrentMaster) && !itemInSlice(link, linksDoneMaster) {

		// update lists of links
		m.Lock()
		linksDoneMaster = append(linksDoneMaster, link)
		for _, item := range linksPageCurrent {
			if !itemInSlice(item, linksCurrentMaster) &&
				!itemInSlice(item, linksDoneMaster) {
				linksCurrentMaster = append(linksCurrentMaster, item)
			}
		}
		linksCurrentMaster = removeItemFromSlice(indexItem(link, linksCurrentMaster), linksCurrentMaster)
		linksInProgress = removeItemFromSlice(indexItem(link, linksInProgress), linksInProgress)
		m.Unlock()
	}
}
