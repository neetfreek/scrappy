package lib

// HTMLTags instance contains global tag string literal values
var HTMLTags tags

// HTMLMap maps html tag node values to global string literals
var HTMLMap map[string]string

// SetupTags initialises HTMLMap and HTMLTags
func SetupTags() {
	makeHTMLTags()
	makeHTMLMap()
}

func makeHTMLTags() {
	HTMLTags.tagHyperLink = "tagHyperLink"
	HTMLTags.tagBlockQuote = "tagBlockQuote"
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
	HTMLMap[HTMLTags.tagHyperLink] = "a"
	HTMLMap[HTMLTags.tagBlockQuote] = "blockquote"
	HTMLMap[HTMLTags.tagBody] = "body"
	HTMLMap[HTMLTags.tagButton] = "button"
	HTMLMap[HTMLTags.tagCite] = "cite"
	HTMLMap[HTMLTags.tagCode] = "code"
	HTMLMap[HTMLTags.tagDescriptionValue] = "dd"
	HTMLMap[HTMLTags.tagDelete] = "del"
	HTMLMap[HTMLTags.tagDetails] = "details"
	HTMLMap[HTMLTags.tagDialogue] = "dialog"
	HTMLMap[HTMLTags.tagDescriptionList] = "dl"
	HTMLMap[HTMLTags.tagDescriptionListTerm] = "dt"
	HTMLMap[HTMLTags.tagEmphasised] = "em"
	HTMLMap[HTMLTags.tagCaptionFigure] = "figcaption"
	HTMLMap[HTMLTags.tagFigure] = "figure"
	HTMLMap[HTMLTags.tagFooter] = "footer"
	HTMLMap[HTMLTags.tagForm] = "form"
	HTMLMap[HTMLTags.tagHeader1] = "h1"
	HTMLMap[HTMLTags.tagHeader2] = "h2"
	HTMLMap[HTMLTags.tagHeader3] = "h3"
	HTMLMap[HTMLTags.tagHeader4] = "h4"
	HTMLMap[HTMLTags.tagHeader5] = "h5"
	HTMLMap[HTMLTags.tagHeader6] = "h6"
	HTMLMap[HTMLTags.tagHead] = "head"
	HTMLMap[HTMLTags.tagHeader] = "header"
	HTMLMap[HTMLTags.tagHTML] = "html"
	HTMLMap[HTMLTags.tagAlternativeVoice] = "i"
	HTMLMap[HTMLTags.tagImage] = "img"
	HTMLMap[HTMLTags.tagInput] = "input"
	HTMLMap[HTMLTags.tagInserted] = "ins"
	HTMLMap[HTMLTags.tagLabel] = "label"
	HTMLMap[HTMLTags.tagLegend] = "legend"
	HTMLMap[HTMLTags.tagList] = "li"
	HTMLMap[HTMLTags.tagLink] = "link"
	HTMLMap[HTMLTags.tagMark] = "mark"
	HTMLMap[HTMLTags.tagMeter] = "meter"
	HTMLMap[HTMLTags.tagNav] = "nav"
	HTMLMap[HTMLTags.tagObject] = "object"
	HTMLMap[HTMLTags.tagListOrdered] = "ol"
	HTMLMap[HTMLTags.tagOptions] = "optgroup"
	HTMLMap[HTMLTags.tagOption] = "option"
	HTMLMap[HTMLTags.tagOutput] = "output"
	HTMLMap[HTMLTags.tagParagrah] = "p"
	HTMLMap[HTMLTags.tagPicture] = "picture"
	HTMLMap[HTMLTags.tagPreformatted] = "pre"
	HTMLMap[HTMLTags.tagQuotation] = "q"
	HTMLMap[HTMLTags.tagTextIncorrect] = "s"
	HTMLMap[HTMLTags.tagSample] = "samp"
	HTMLMap[HTMLTags.tagScript] = "script"
	HTMLMap[HTMLTags.tagSection] = "section"
	HTMLMap[HTMLTags.tagSelect] = "select"
	HTMLMap[HTMLTags.tagSmall] = "small"
	HTMLMap[HTMLTags.tagSource] = "source"
	HTMLMap[HTMLTags.tagSpan] = "span"
	HTMLMap[HTMLTags.tagStrong] = "strong"
	HTMLMap[HTMLTags.tagSubscript] = "sub"
	HTMLMap[HTMLTags.tagHDetailsHeading] = "summary"
	HTMLMap[HTMLTags.tagSuperscript] = "sup"
	HTMLMap[HTMLTags.tagTable] = "table"
	HTMLMap[HTMLTags.tagTableBody] = "tbody"
	HTMLMap[HTMLTags.tagTableCell] = "td"
	HTMLMap[HTMLTags.tagTextArea] = "textarea"
	HTMLMap[HTMLTags.tagTableFooter] = "textfoot"
	HTMLMap[HTMLTags.tagTableHeaderCell] = "th"
	HTMLMap[HTMLTags.tagTableHeader] = "thead"
	HTMLMap[HTMLTags.tagDateTime] = "time"
	HTMLMap[HTMLTags.tagTitle] = "title"
	HTMLMap[HTMLTags.tagTableRow] = "tr"
	HTMLMap[HTMLTags.tagUnderline] = "u"
	HTMLMap[HTMLTags.tagListUnordered] = "ul"
	HTMLMap[HTMLTags.tagVariable] = "var"

	return HTMLMap
}

type tags struct {
	tagHyperLink           string
	tagBlockQuote          string
	tagBody                string
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
