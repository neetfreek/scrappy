package lib

/*==================================================================================*
* Call functions handling page, site actions depending on user input				*
*===================================================================================*
* Functions are called from different menu selections made by user.					*
* Functions handle starting routines based on menu selections made by the user.		*
* Options are represented as CONSTANTS mapped to the options displayed to the user	*
*	in menus.																		*
*===================================================================================*/

func startActionMain(userAction string) {
	switch userAction {
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

func startActionSite(userAction string) {
	if userAction == actionMenuMain {
		MenuMain()
	} else {
		pageURL := getPageSelection()
		if validResponse(pageURL) {
			scopes := getDomainScopes(pageURL)
			scope := getScopeSelection(scopes)
			crawlSite(scope, userAction)
		}
	}
}

func startActionPage(userAction string) {
	if userAction == actionMenuMain {
		MenuMain()
	} else {
		pageURL := getPageSelection()
		if validResponse(pageURL) {
			getPageContent(pageURL, userAction)
		}
	}
}
