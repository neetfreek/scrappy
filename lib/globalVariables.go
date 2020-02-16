package lib

// HTMLTags instance contains global tag string literal values
var HTMLTags tags

// HTMLMap maps html tag node values to global string literals
var HTMLMap map[string]string

// InputOptionsMap maps user input options TODO: CHANGE TO ARRAY FOR ORDERED ITERATING
var InputOptionsMap = []string{
	pageActionSaveText,
	pageActionSaveLinks,
	pageActionSavePage,
	pageActionSaveImages,
	messageActionExit,
}

// SetupConstants initialises HTMLMap and HTMLTags
func SetupConstants() {
	makeHTMLTags()
	makeHTMLMap()
}

func makeHTMLTags() {
	HTMLTags.tagHyperLink = "tagHyperLink"
	HTMLTags.tagBlockQuote = "tagBlockQuote"
	HTMLTags.tagBold = "tagBold"
	HTMLTags.tagBody = "tagBody"
	HTMLTags.tagButton = "tagButton"
	HTMLTags.tagCite = "tagCite"
	HTMLTags.tagCode = "tagCode"
	HTMLTags.tagDescriptionValue = "tagDescriptionValue"
	HTMLTags.tagDelete = "tagDelete"
	HTMLTags.tagDetails = "tagDetails"
	HTMLTags.tagDialogue = "tagDialogue"
	HTMLTags.tagDescriptionList = "tagDescriptionList"
	HTMLTags.tagDescriptionListTerm = "tagDescriptionListTerm"
	HTMLTags.tagEmphasised = "tagEmphasised"
	HTMLTags.tagCaptionFigure = "tagCaptionFigure"
	HTMLTags.tagFigure = "tagFigure"
	HTMLTags.tagFooter = "tagFooter"
	HTMLTags.tagForm = "tagForm"
	HTMLTags.tagHeader1 = "tagHeader1"
	HTMLTags.tagHeader2 = "tagHeader2"
	HTMLTags.tagHeader3 = "tagHeader3"
	HTMLTags.tagHeader4 = "tagHeader4"
	HTMLTags.tagHeader5 = "tagHeader5"
	HTMLTags.tagHeader6 = "tagHeader6"
	HTMLTags.tagHead = "tagHead"
	HTMLTags.tagHeader = "tagHeader"
	HTMLTags.tagHTML = "tagHTML"
	HTMLTags.tagAlternativeVoice = "tagAlternativeVoice"
	HTMLTags.tagImage = "tagImage"
	HTMLTags.tagInput = "tagInput"
	HTMLTags.tagInserted = "tagInserted"
	HTMLTags.tagLabel = "tagLabel"
	HTMLTags.tagLegend = "tagLegend"
	HTMLTags.tagList = "tagList"
	HTMLTags.tagLink = "tagLink"
	HTMLTags.tagMark = "tagMark"
	HTMLTags.tagMeter = "tagMeter"
	HTMLTags.tagNav = "tagNav"
	HTMLTags.tagObject = "tagObject"
	HTMLTags.tagListOrdered = "tagListOrdered"
	HTMLTags.tagOptions = "tagOptions"
	HTMLTags.tagOption = "tagOption"
	HTMLTags.tagOutput = "tagOutput"
	HTMLTags.tagParagrah = "tagParagrah"
	HTMLTags.tagPicture = "tagPicture"
	HTMLTags.tagPreformatted = "tagPreformatted"
	HTMLTags.tagQuotation = "tagQuotation"
	HTMLTags.tagTextIncorrect = "tagTextIncorrect"
	HTMLTags.tagSample = "tagSample"
	HTMLTags.tagScript = "tagScript"
	HTMLTags.tagSection = "tagSection"
	HTMLTags.tagSelect = "tagSelect"
	HTMLTags.tagSmall = "tagSmall"
	HTMLTags.tagSource = "tagSource"
	HTMLTags.tagSpan = "tagSpan"
	HTMLTags.tagStrong = "tagStrong"
	HTMLTags.tagSubscript = "tagSubscript"
	HTMLTags.tagHDetailsHeading = "tagHDetailsHeading"
	HTMLTags.tagSuperscript = "tagSuperscript"
	HTMLTags.tagTable = "tagTable"
	HTMLTags.tagTableBody = "tagTableBody"
	HTMLTags.tagTableCell = "tagTableCell"
	HTMLTags.tagTextArea = "tagTextArea"
	HTMLTags.tagTableFooter = "tagTableFooter"
	HTMLTags.tagTableHeaderCell = "tagTableHeaderCell"
	HTMLTags.tagTableHeader = "tagTableHeader"
	HTMLTags.tagDateTime = "tagDateTime"
	HTMLTags.tagTitle = "tagTitle"
	HTMLTags.tagTableRow = "tagTableRow"
	HTMLTags.tagUnderline = "tagUnderline"
	HTMLTags.tagListUnordered = "tagListUnordered"
	HTMLTags.tagVariable = "tagVariable"
}

func makeHTMLMap() map[string]string {
	HTMLMap = make(map[string]string)
	HTMLMap["a"] = HTMLTags.tagHyperLink
	HTMLMap["blockquote"] = HTMLTags.tagBlockQuote
	HTMLMap["body"] = HTMLTags.tagBody
	HTMLMap["b"] = HTMLTags.tagBold
	HTMLMap["button"] = HTMLTags.tagButton
	HTMLMap["cite"] = HTMLTags.tagCite
	HTMLMap["code"] = HTMLTags.tagCode
	HTMLMap["dd"] = HTMLTags.tagDescriptionValue
	HTMLMap["del"] = HTMLTags.tagDelete
	HTMLMap["details"] = HTMLTags.tagDetails
	HTMLMap["dialog"] = HTMLTags.tagDialogue
	HTMLMap["dl"] = HTMLTags.tagDescriptionList
	HTMLMap["dt"] = HTMLTags.tagDescriptionListTerm
	HTMLMap["em"] = HTMLTags.tagEmphasised
	HTMLMap["figcaption"] = HTMLTags.tagCaptionFigure
	HTMLMap["figure"] = HTMLTags.tagFigure
	HTMLMap["footer"] = HTMLTags.tagFooter
	HTMLMap["form"] = HTMLTags.tagForm
	HTMLMap["h1"] = HTMLTags.tagHeader1
	HTMLMap["h2"] = HTMLTags.tagHeader2
	HTMLMap["h3"] = HTMLTags.tagHeader3
	HTMLMap["h4"] = HTMLTags.tagHeader4
	HTMLMap["h5"] = HTMLTags.tagHeader5
	HTMLMap["h6"] = HTMLTags.tagHeader6
	HTMLMap["head"] = HTMLTags.tagHead
	HTMLMap["header"] = HTMLTags.tagHeader
	HTMLMap["html"] = HTMLTags.tagHTML
	HTMLMap["i"] = HTMLTags.tagAlternativeVoice
	HTMLMap["img"] = HTMLTags.tagImage
	HTMLMap["input"] = HTMLTags.tagInput
	HTMLMap["ins"] = HTMLTags.tagInserted
	HTMLMap["label"] = HTMLTags.tagLabel
	HTMLMap["legend"] = HTMLTags.tagLegend
	HTMLMap["li"] = HTMLTags.tagList
	HTMLMap["link"] = HTMLTags.tagLink
	HTMLMap["mark"] = HTMLTags.tagMark
	HTMLMap["meter"] = HTMLTags.tagMeter
	HTMLMap["nav"] = HTMLTags.tagNav
	HTMLMap["object"] = HTMLTags.tagObject
	HTMLMap["ol"] = HTMLTags.tagListOrdered
	HTMLMap["optgroup"] = HTMLTags.tagOptions
	HTMLMap["option"] = HTMLTags.tagOption
	HTMLMap["output"] = HTMLTags.tagOutput
	HTMLMap["p"] = HTMLTags.tagParagrah
	HTMLMap["picture"] = HTMLTags.tagPicture
	HTMLMap["pre"] = HTMLTags.tagPreformatted
	HTMLMap["q"] = HTMLTags.tagQuotation
	HTMLMap["s"] = HTMLTags.tagTextIncorrect
	HTMLMap["samp"] = HTMLTags.tagSample
	HTMLMap["script"] = HTMLTags.tagScript
	HTMLMap["section"] = HTMLTags.tagSection
	HTMLMap["select"] = HTMLTags.tagSelect
	HTMLMap["small"] = HTMLTags.tagSmall
	HTMLMap["source"] = HTMLTags.tagSource
	HTMLMap["span"] = HTMLTags.tagSpan
	HTMLMap["strong"] = HTMLTags.tagStrong
	HTMLMap["sub"] = HTMLTags.tagSubscript
	HTMLMap["summary"] = HTMLTags.tagHDetailsHeading
	HTMLMap["sup"] = HTMLTags.tagSuperscript
	HTMLMap["table"] = HTMLTags.tagTable
	HTMLMap["tbody"] = HTMLTags.tagTableBody
	HTMLMap["td"] = HTMLTags.tagTableCell
	HTMLMap["textarea"] = HTMLTags.tagTextArea
	HTMLMap["textfoot"] = HTMLTags.tagTableFooter
	HTMLMap["th"] = HTMLTags.tagTableHeaderCell
	HTMLMap["thead"] = HTMLTags.tagTableHeader
	HTMLMap["time"] = HTMLTags.tagDateTime
	HTMLMap["title"] = HTMLTags.tagTitle
	HTMLMap["tr"] = HTMLTags.tagTableRow
	HTMLMap["u"] = HTMLTags.tagUnderline
	HTMLMap["ul"] = HTMLTags.tagListUnordered
	HTMLMap["var"] = HTMLTags.tagVariable

	return HTMLMap
}

type tags struct {
	tagHyperLink           string
	tagBlockQuote          string
	tagBody                string
	tagBold                string
	tagButton              string
	tagCite                string
	tagCode                string
	tagDescriptionValue    string
	tagDelete              string
	tagDetails             string
	tagDialogue            string
	tagDescriptionList     string
	tagDescriptionListTerm string
	tagEmphasised          string
	tagCaptionFigure       string
	tagFigure              string
	tagFooter              string
	tagForm                string
	tagHeader1             string
	tagHeader2             string
	tagHeader3             string
	tagHeader4             string
	tagHeader5             string
	tagHeader6             string
	tagHead                string
	tagHeader              string
	tagHTML                string
	tagAlternativeVoice    string
	tagImage               string
	tagInput               string
	tagInserted            string
	tagLabel               string
	tagLegend              string
	tagList                string
	tagLink                string
	tagMark                string
	tagMeter               string
	tagNav                 string
	tagObject              string
	tagListOrdered         string
	tagOptions             string
	tagOption              string
	tagOutput              string
	tagParagrah            string
	tagPicture             string
	tagPreformatted        string
	tagQuotation           string
	tagTextIncorrect       string
	tagSample              string
	tagScript              string
	tagSection             string
	tagSelect              string
	tagSmall               string
	tagSource              string
	tagSpan                string
	tagStrong              string
	tagSubscript           string
	tagHDetailsHeading     string
	tagSuperscript         string
	tagTable               string
	tagTableBody           string
	tagTableCell           string
	tagTextArea            string
	tagTableFooter         string
	tagTableHeaderCell     string
	tagTableHeader         string
	tagDateTime            string
	tagTitle               string
	tagTableRow            string
	tagUnderline           string
	tagListUnordered       string
	tagVariable            string
}

// ExitRequested handles whether app loop terminates
var ExitRequested = false

// Standard use messages to player
var messageGreeting = "Welcome to scrappy!"
var messageChoosePageAction = "Choose an action: "
var messageChoosePageURL = "Enter a URL: "
var messageNoPage = "Could not find page"
var messageActionExit = "Exit application"

// Page action input options
var pageActionSaveText = "Save page text content to file (.txt)"
var pageActionSaveLinks = "Save page links to file (.txt)"
var pageActionSavePage = "Save page to file (.html)"
var pageActionSaveImages = "Save page images to directory"

// Parameters
var typeDirectory = "directory"
var typeFile = "file"

var imageFormats = []string{".jpg", ".jpeg", ".tiff", ".gif", ".bmp", ".pbm", ".pgm", ".ppm", "pnm", ".png", ".webp"}
