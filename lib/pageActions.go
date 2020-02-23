package lib

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
		pageURL := getUserPageOption()
		crawlSite(pageURL)
	}
}

func startActionPage(option string) {
	if option == actionMenuMain {
		MenuMain()
	} else {
		pageURL := getUserPageOption()
		getPageContent(pageURL, option)
	}
}
