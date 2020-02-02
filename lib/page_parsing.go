package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetNodeContent gets specified node content
func GetNodeContent(pageResponseString, tag string) {

	// doc is *html.Node which can be crawled. Here it's parsed using golang's /x/net/html func Parse example for links
	doc, err := html.Parse(strings.NewReader(pageResponseString))
	if err != nil {
		log.Fatal(err)
	}
	crawlNodes(tag, doc)
}

func crawl(n *html.Node) {
	// n.Data refers to tag name for element nodes, content for text
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		crawl(c)
	}
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func crawlNodes(tag string, n *html.Node) {

	if n.Type == html.ElementNode && n.Data == tag {
		fmt.Printf("\nNODE ALL: %s\n", renderNode(n))
		fmt.Printf("NODE DATA: %s\n", n.Data)            // node tag
		fmt.Printf("NODE DATA_ATOM: %s\n", n.DataAtom)   // node attr
		fmt.Printf("NODE 1ST_CHILD: %v\n", n.FirstChild) // node 1st_child **
		fmt.Printf("NODE TYPE: %v\n", n.Type)            // node type
	}

	for childNode := n.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		crawlNodes(tag, childNode)
	}
}
// Tester tests
func Tester(url string) {
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
		case html.StartTagToken:
			// fmt.Printf("\n%v\n", responsToken)
			tag := responseToken.Token()
			tagType := tag.Data
			tokenType(tagType)
			// fmt.Printf("\n%v\n", tag.Data)

			// Print href links (url and image links)
			for _, a := range tag.Attr {
				if a.Key == "href" {
					fmt.Println("Found :", a.Val)
					break
				}
			}
		}
	}
}

func tokenType(tokenType string) {
	switch tokenType {
	case "p":
		fmt.Print("PARAGRAPH")
	case "a":
		fmt.Print("ANCHOR")
	default:
		fmt.Print("SOME OTHER TAG TYPE")
	}
}
