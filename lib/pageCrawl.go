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
	"strings"
)

var linksToDoCurrent = []string{}
var linksToDoNext = []string{}
var linksDone = []string{}
var domain = ""

func crawlSite(url string) {

	domain = pageDomain(url)
	linksToDoNext = append(linksToDoNext, strings.Trim(domain, "/"))
	linksToDoCurrent = append(linksToDoCurrent, strings.Trim(url, "/"))

	crawPageForLinks()
}

func crawPageForLinks() {
	for counter, link := range linksToDoCurrent {
		resp := pageResponse(linksToDoCurrent[counter])
		defer resp.Body.Close()

		pageLinks := loopGetPage(resp.Body, pageActionSaveLinks)
		linksToDoNext = addToLinksToDoNext(pageLinks, domain)
		linksDone = append(linksDone, link)
	}
	linksToDoCurrent = nil
	linksToDoCurrent = copyItemsToSlice(linksToDoNext, linksToDoCurrent)
	linksToDoNext = nil
}
