package lib

/*==================================================================================*
* Call functions handling page, site actions depending on user input				*
*===================================================================================*
* Functions are called from different menu selections made by user.					*
* Functions handle starting routines based on menu selections made by the user.		*
* Options are represented as CONSTANTS mapped to the options displayed to the user	*
*	in menus.																		*
*===================================================================================*/

func startActionMain(option string) {
	switch option {
	case actionPageContent:
		menuPage()
		break
	case actionSiteContent:
		menuSite()
		break
	case actionExit:
		ExitRequested = true
		break
	}
}

func startActionSite(option string) {
	if option == actionMenuMain {
		MenuMain()
	} else {
		pageURL := getPageSelection()
		scopes := getDomainScopes(pageURL)
		scope := getScopeSelection(scopes)
		crawlSite(scope)
	}
}

func startActionPage(option string) {
	if option == actionMenuMain {
		MenuMain()
	} else {
		pageURL := getPageSelection()
		getPageContent(pageURL, option)
	}
}
