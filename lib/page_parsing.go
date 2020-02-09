package lib

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// PrintAttribute for testing print http tag attributes to standard output
func PrintAttribute(url string) {
	// Get response
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Make response token
	responseToken := html.NewTokenizer(resp.Body)

	for {
		responsToken := responseToken.Next()

		switch responsToken {
		case html.ErrorToken:
			return
		//For working with tags, e.g. inspecting tag values or attributes
		case html.StartTagToken:
			tag := responseToken.Token()
			tagType := tag.Data
			fmt.Print("\nStartTagToken tag: ", tag)
			fmt.Print("\ntagType: ", tagType)
			// if (IdentifyTag(tagType)) != "" {
			// 	for _, attribute := range tag.Attr {
			// 		fmt.Print("\n tagAttribute: ", attribute.Val)
			// 		break
			// 	}
			// }
		case html.TextToken:
			tag := responseToken.Token()
			fmt.Print("\nTextToken tag: ", tag)
			// if (IdentifyTag(tag.String())) == HTMLTags.tagParagrah {
			// 	fmt.Print(tag, "\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

			// }
			// fmt.Print(tag, "\n")
		}
	}
}
