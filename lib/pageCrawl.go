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
var domain = ""

func crawlSite(url string) {
	// Setup domain, start pages
	domain = pageDomain(url)
	linksCurrentMaster = append(linksCurrentMaster, strings.Trim(url, "/"))
	if url != domain {
		linksNextMaster = append(linksNextMaster, strings.Trim(domain, "/"))
	}
	// Setup sync helpers
	var wg sync.WaitGroup
	var m sync.Mutex
	// Crawl all pages' links
	for len(linksCurrentMaster) > 0 && len(linksNextMaster) > 0 {
		for counter, link := range linksCurrentMaster {
			wg.Add(1)
			fmt.Printf("%v: Crawling page at URL %v\n", counter, link)
			go crawlPageForLinks(&wg, &m, link)
		}
	}
	defer wg.Wait()
	printCollection(linksDoneMaster, "DONE:")
}

func crawlPageForLinks(wg *sync.WaitGroup, m *sync.Mutex, link string) {
	defer wg.Done()
	// get page links
	resp := pageResponse(link)
	defer resp.Body.Close()
	linksPageCurrent = loopGetPage(resp.Body, pageActionSaveLinks)

	if itemInSlice(link, linksCurrentMaster) && !itemInSlice(link, linksDoneMaster) {
		m.Lock()

		removeItemFromSlice(indexItem(link, linksCurrentMaster), linksCurrentMaster)
		linksDoneMaster = append(linksDoneMaster, link)

		for _, item := range linksPageCurrent {
			if indexItem(item, linksCurrentMaster) == -1 &&
				indexItem(item, linksCurrentMaster) == -1 {
				linksCurrentMaster = append(linksCurrentMaster, item)
			}
		}

		m.Unlock()
	}

	fmt.Printf("Found %v links so far...\n", len(linksDoneMaster))
}
