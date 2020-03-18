package lib

/*==================================================================================*
* Examine pages, sites and retrieve data from multiple pages						*
*===================================================================================*
* The crawlSite routine goes through link and domain URL looking for link URLs to	*
*	other pages within the domain. Each link is crawled by a seperate go routine,	*
*	making use of wait groups and mutexes.											*
*	1) Link URLs are put into linksCurrent, from where each are in turn crawled for	*
*		link URLs they contain.														*
*	2) On start crawl, each link is added to linksInProgress so other go routines	*
*		don't also crawl that link.												 	*
*	2) URLS not already in any lists are added to linksCurrent.						*
*	3) Once crawled, URLS are placed into linksDone and removed from linksCurrent	*
*	4) Once all linksCurrent URLs have been crawled, the list's contents are		*
*		replaced by those of linksToDoNext, and linksToDoNext is cleared.			*
*	5) The process is complete when there are no more links from linksCurrent left	*
*		to crawl.																	*
*===================================================================================*/

import (
	"strings"
	"sync"
	"time"
)

var linksCurrent = []string{}
var linksDone = []string{}
var linksImages = []string{}
var linksInProgress = []string{}

func crawlSite(url, userAction string) {
	resetSlices()
	defer Timer(time.Now(), "Total time:")
	// Setup domain, start pages
	var domain = pageDomainName(url)
	linksCurrent = append(linksCurrent, strings.Trim(url, "/"))
	// Setup sync helpers
	var m sync.Mutex
	// Crawl all pages' links
	for len(linksCurrent) > 0 {
		for _, link := range linksCurrent {
			if indexItem(link, linksInProgress) == -1 {
				linksInProgress = append(linksInProgress, link)
				go crawlPageForLinks(&m, link, domain, userAction)
			}
		}
	}
	crawlOutput(url, userAction)
}

func crawlPageForLinks(m *sync.Mutex, link, domain, userAction string) {
	// get page links
	resp := pageResponse(link)
	defer resp.Body.Close()

	linksImagesCurrent := []string{}
	linksPageCurrent := loopGetPage(resp.Body, pageActionSaveLinks)
	// store image links if userAction to save image links
	if userAction == siteActionSaveImageLinks {
		linksImagesCurrent = addToLinksImages(linksPageCurrent, linksImagesCurrent)
	}
	if userAction == siteActionSaveText {

	}
	// store only URL links in domain
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
		linksImages = addToLinksImages(linksImagesCurrent, linksImages)
		linksCurrent = removeItemFromSlice(indexItem(link, linksCurrent), linksCurrent)
		linksInProgress = removeItemFromSlice(indexItem(link, linksInProgress), linksInProgress)
		m.Unlock()
	}
}
