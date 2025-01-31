package lib

/*==================================================================================*
* Global variables, definitions of menu options, and HTML tags and maps				*
*===================================================================================*
* Many variables are defined here as CONSTANTS to allow for re-use, to prevent the 	*
*	use of "magic" numbers and strings, and to facilitate for easy modifications.	*
* HTML tags are stored in the HTMLTags struct, and are accessed from here for,		*
*	e.g., comparing their values to those returned in accessing page HTML.			*
* HTML tags are mapped to CONSTANT values for easier use.							*
* Collections of menu options are stored as CONSTANT variables, which reflect both	*
*	the options presented to the user in menus as well as the internal values used	*
*	in determining which routines and functions are called.							*
*===================================================================================*/

// SetupConstants initialises HTMLMap and HTMLTags
func SetupConstants() {
	makeHTMLTags()
	makeHTMLMap()
}

// HTMLTags instance contains global tag string literal values
var HTMLTags tags

// HTMLMap maps html tag node values to global string literals
var HTMLMap map[string]string

// InputOptionsMapMain maps user input to top-level application options
var InputOptionsMapMain = []string{
	actionPageContent,
	// actionSiteContent,
	actionExit,
}

// InputOptionsMapSite maps user input to site-content options
var InputOptionsMapSite = []string{
	siteActionSaveImageLinks,
	siteActionSaveLinks,
	siteActionSaveText,
	actionMenuMain,
}

// InputOptionsMapPage maps user input options to page-content options
var InputOptionsMapPage = []string{
	pageActionSaveText,
	pageActionSaveLinks,
	pageActionSavePage,
	pageActionSaveImageLinks,
	actionMenuMain,
}

// ExitRequested handles whether app loop terminates
var ExitRequested = false

// Standard use messagesTimer to player
var messageGreeting = "Welcome to scrappy!"
var messageChoosePageAction = "Selection: "
var messageChoosePageURL = "Enter a URL: "
var messageNoPage = "Could not find page"
var messageNotURLFormat = "URL format error:"
var messageSelectScope = "Select scope: "

// Top-level action input options
var actionPageContent = "Get page content"
var actionSiteContent = "Get site content"
var actionMenuMain = "Return to main menu"
var actionExit = "Exit application"

// Site action input options
var siteActionSaveImageLinks = "Save site image links to single file (.txt)"
var siteActionSaveLinks = "Save site links to single file (.txt)"
var siteActionSaveText = "Save site text content to single file (.txt)"

// Page action input options
var pageActionSaveText = "Save page text content to file (.txt)"
var pageActionSaveLinks = "Save page links to file (.txt)"
var pageActionSavePage = "Save page to file (.html)"
var pageActionSaveImageLinks = "Save page image links to file (.txt)"

// File information
var pageURL = "Page URL"
var pageAccessed = "Page Accessed"
var pageUntitled = "Untitled_"

// Common strings
var addressRobotsTxt = "/robots.txt"
var delimiter = ": "
var delimiterDomain = "://"
var doubleSlash = "//"
var slash = "/"
var newLine = "\n"
var thisDirectoryDelimiter = "./"
var underscore = "_"
var dot = "."

// Directory, file name suffixes
var suffixImageLinks = "_linksImage"
var suffixLinks = "_links"
var suffixText = "_textContent"
var suffixHTML = "_htmlContent"

// Parameters
var typeDirectory = "directory"
var typeFile = "file"
var goClient = "Go-http-client/1.1"

var imageFormats = []string{".jpg", ".jpeg", ".tiff", ".gif", ".bmp", ".pbm", ".pgm", ".ppm", "pnm", ".png", ".svg", ".webp"}

func makeHTMLTags() {
	HTMLTags.tagAbbreviation = "tagAbbreviation"
	HTMLTags.tagArea = "tagArea"
	HTMLTags.tagArticle = "tagArticle"
	HTMLTags.tagCaption = "tagCaption"
	HTMLTags.tagColumn = "tagColumn"
	HTMLTags.tagColumnGroup = "tagColumnGroup"
	HTMLTags.tagData = "tagData"
	HTMLTags.tagDataList = "tagDataList"
	HTMLTags.tagDefiningInstance = "tagDefiningInstance"
	HTMLTags.tagBiDirectionalIsolation = "tagBiDirectionalIsolation"
	HTMLTags.tagBiDirectionalOverride = "tagBiDirectionalOverride"
	HTMLTags.tagHyperLink = "tagHyperLink"
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
	HTMLTags.tagDivision = "tagDivision"
	HTMLTags.tagDescriptionList = "tagDescriptionList"
	HTMLTags.tagDescriptionListTerm = "tagDescriptionListTerm"
	HTMLTags.tagEmphasised = "tagEmphasised"
	HTMLTags.tagCaptionFigure = "tagCaptionFigure"
	HTMLTags.tagFieldSet = "tagFieldSet"
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
	HTMLTags.tagInlineFrame = "tagInlineFrame"
	HTMLTags.tagKeyboardInput = "tagKeyboardInput"
	HTMLTags.tagRoot = "tagRoot"
	HTMLTags.tagAlternativeVoice = "tagAlternativeVoice"
	HTMLTags.tagImage = "tagImage"
	HTMLTags.tagInput = "tagInput"
	HTMLTags.tagInserted = "tagInserted"
	HTMLTags.tagLabel = "tagLabel"
	HTMLTags.tagLegend = "tagLegend"
	HTMLTags.tagList = "tagList"
	HTMLTags.tagLink = "tagLink"
	HTMLTags.tagMain = "tagMain"
	HTMLTags.tagMark = "tagMark"
	HTMLTags.tagMetadata = "tagMetadata"
	HTMLTags.tagMeter = "tagMeter"
	HTMLTags.tagNav = "tagNav"
	HTMLTags.tagNoScript = "tagNoScript"
	HTMLTags.tagObject = "tagObject"
	HTMLTags.tagListOrdered = "tagListOrdered"
	HTMLTags.tagOptions = "tagOptions"
	HTMLTags.tagOption = "tagOption"
	HTMLTags.tagOutput = "tagOutput"
	HTMLTags.tagParagrah = "tagParagrah"
	HTMLTags.tagParameter = "tagParameter"
	HTMLTags.tagPicture = "tagPicture"
	HTMLTags.tagPreformatted = "tagPreformatted"
	HTMLTags.tagRubyAnnotation = "tagRubyAnnotation"
	HTMLTags.tagRubyParentheses = "tagRubyParentheses"
	HTMLTags.tagRubyPronounciation = "tagRubyPronounciation"
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
	HTMLTags.tagSVG = "tagSVG"
	HTMLTags.tagHDetailsHeading = "tagHDetailsHeading"
	HTMLTags.tagSuperscript = "tagSuperscript"
	HTMLTags.tagTable = "tagTable"
	HTMLTags.tagTableBody = "tagTableBody"
	HTMLTags.tagTableCell = "tagTableCell"
	HTMLTags.tagTemplate = "tagTemplate"
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
	HTMLTags.tagVideo = "tagVideo"
}

func makeHTMLMap() map[string]string {
	HTMLMap = make(map[string]string)
	HTMLMap["a"] = HTMLTags.tagHyperLink
	HTMLMap["abbr"] = HTMLTags.tagAbbreviation
	HTMLMap["address"] = HTMLTags.tagAddress
	HTMLMap["article"] = HTMLTags.tagArticle
	HTMLMap["area"] = HTMLTags.tagArea
	HTMLMap["audio"] = HTMLTags.tagAudio
	HTMLMap["blockquote"] = HTMLTags.tagBlockQuote
	HTMLMap["body"] = HTMLTags.tagBody
	HTMLMap["b"] = HTMLTags.tagBold
	HTMLMap["bdi"] = HTMLTags.tagBiDirectionalIsolation
	HTMLMap["bdo"] = HTMLTags.tagBiDirectionalOverride
	HTMLMap["blockquote"] = HTMLTags.tagBlockQuote
	HTMLMap["button"] = HTMLTags.tagButton
	HTMLMap["caption"] = HTMLTags.tagCaption
	HTMLMap["cite"] = HTMLTags.tagCite
	HTMLMap["code"] = HTMLTags.tagCode
	HTMLMap["col"] = HTMLTags.tagColumn
	HTMLMap["colgroup"] = HTMLTags.tagColumnGroup
	HTMLMap["data"] = HTMLTags.tagData
	HTMLMap["datalist"] = HTMLTags.tagDataList
	HTMLMap["del"] = HTMLTags.tagDelete
	HTMLMap["dd"] = HTMLTags.tagDescriptionValue
	HTMLMap["details"] = HTMLTags.tagDetails
	HTMLMap["dfn"] = HTMLTags.tagDefiningInstance
	HTMLMap["dialog"] = HTMLTags.tagDialogue
	HTMLMap["div"] = HTMLTags.tagDivision
	HTMLMap["dl"] = HTMLTags.tagDescriptionList
	HTMLMap["dt"] = HTMLTags.tagDescriptionListTerm
	HTMLMap["em"] = HTMLTags.tagEmphasised
	HTMLMap["fieldset"] = HTMLTags.tagFieldSet
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
	HTMLMap["html"] = HTMLTags.tagRoot
	HTMLMap["i"] = HTMLTags.tagAlternativeVoice
	HTMLMap["img"] = HTMLTags.tagImage
	HTMLMap["input"] = HTMLTags.tagInput
	HTMLMap["ins"] = HTMLTags.tagInserted
	HTMLMap["iframe"] = HTMLTags.tagInlineFrame
	HTMLMap["kbd"] = HTMLTags.tagKeyboardInput
	HTMLMap["label"] = HTMLTags.tagLabel
	HTMLMap["legend"] = HTMLTags.tagLegend
	HTMLMap["li"] = HTMLTags.tagList
	HTMLMap["link"] = HTMLTags.tagLink
	HTMLMap["main"] = HTMLTags.tagMain
	HTMLMap["mark"] = HTMLTags.tagMark
	HTMLMap["meta"] = HTMLTags.tagMetadata
	HTMLMap["meter"] = HTMLTags.tagMeter
	HTMLMap["nav"] = HTMLTags.tagNav
	HTMLMap["noscript"] = HTMLTags.tagNoScript
	HTMLMap["object"] = HTMLTags.tagObject
	HTMLMap["ol"] = HTMLTags.tagListOrdered
	HTMLMap["optgroup"] = HTMLTags.tagOptions
	HTMLMap["option"] = HTMLTags.tagOption
	HTMLMap["output"] = HTMLTags.tagOutput
	HTMLMap["p"] = HTMLTags.tagParagrah
	HTMLMap["parameter"] = HTMLTags.tagParameter
	HTMLMap["picture"] = HTMLTags.tagPicture
	HTMLMap["pre"] = HTMLTags.tagPreformatted
	HTMLMap["q"] = HTMLTags.tagQuotation
	HTMLMap["rp"] = HTMLTags.tagRubyParentheses
	HTMLMap["rt"] = HTMLTags.tagRubyPronounciation
	HTMLMap["ruby"] = HTMLTags.tagRubyAnnotation
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
	HTMLMap["svg"] = HTMLTags.tagSVG
	HTMLMap["table"] = HTMLTags.tagTable
	HTMLMap["tbody"] = HTMLTags.tagTableBody
	HTMLMap["td"] = HTMLTags.tagTableCell
	HTMLMap["template"] = HTMLTags.tagTemplate
	HTMLMap["textarea"] = HTMLTags.tagTextArea
	HTMLMap["tfoot"] = HTMLTags.tagTableFooter
	HTMLMap["th"] = HTMLTags.tagTableHeaderCell
	HTMLMap["thead"] = HTMLTags.tagTableHeader
	HTMLMap["time"] = HTMLTags.tagDateTime
	HTMLMap["title"] = HTMLTags.tagTitle
	HTMLMap["tr"] = HTMLTags.tagTableRow
	HTMLMap["u"] = HTMLTags.tagUnderline
	HTMLMap["ul"] = HTMLTags.tagListUnordered
	HTMLMap["var"] = HTMLTags.tagVariable
	HTMLMap["video"] = HTMLTags.tagVideo

	return HTMLMap
}

type tags struct {
	tagAbbreviation           string
	tagAddress                string
	tagArea                   string
	tagArticle                string
	tagCaption                string
	tagColumn                 string
	tagColumnGroup            string
	tagData                   string
	tagDataList               string
	tagDefiningInstance       string
	tagDivision               string
	tagHyperLink              string
	tagAudio                  string
	tagBiDirectionalIsolation string
	tagBiDirectionalOverride  string
	tagBlockQuote             string
	tagBody                   string
	tagBold                   string
	tagButton                 string
	tagCite                   string
	tagCode                   string
	tagDescriptionValue       string
	tagDelete                 string
	tagDetails                string
	tagDialogue               string
	tagDescriptionList        string
	tagDescriptionListTerm    string
	tagEmphasised             string
	tagFieldSet               string
	tagCaptionFigure          string
	tagFigure                 string
	tagFooter                 string
	tagForm                   string
	tagHeader1                string
	tagHeader2                string
	tagHeader3                string
	tagHeader4                string
	tagHeader5                string
	tagHeader6                string
	tagHead                   string
	tagHeader                 string
	tagInlineFrame            string
	tagRoot                   string
	tagAlternativeVoice       string
	tagImage                  string
	tagInput                  string
	tagInserted               string
	tagKeyboardInput          string
	tagLabel                  string
	tagLegend                 string
	tagList                   string
	tagLink                   string
	tagMain                   string
	tagMark                   string
	tagMetadata               string
	tagMeter                  string
	tagNav                    string
	tagNoScript               string
	tagObject                 string
	tagListOrdered            string
	tagOptions                string
	tagOption                 string
	tagOutput                 string
	tagParagrah               string
	tagParameter              string
	tagPicture                string
	tagPreformatted           string
	tagQuotation              string
	tagRubyAnnotation         string
	tagRubyParentheses        string
	tagRubyPronounciation     string
	tagTextIncorrect          string
	tagSample                 string
	tagScript                 string
	tagSection                string
	tagSelect                 string
	tagSmall                  string
	tagSource                 string
	tagSpan                   string
	tagStrong                 string
	tagSubscript              string
	tagSVG                    string
	tagHDetailsHeading        string
	tagSuperscript            string
	tagTable                  string
	tagTableBody              string
	tagTableCell              string
	tagTemplate               string
	tagTextArea               string
	tagTableFooter            string
	tagTableHeaderCell        string
	tagTableHeader            string
	tagDateTime               string
	tagTitle                  string
	tagTableRow               string
	tagUnderline              string
	tagListUnordered          string
	tagVariable               string
	tagVideo                  string
}
