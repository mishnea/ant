package scan

import (
	"strings"

	"golang.org/x/net/html"
)

// Attr returns an attribute value by its key.
//
// `ok` is true if the value exists.
func Attr(n *html.Node, key string) (val string, ok bool) {
	if n == nil {
		return
	}

	for _, a := range n.Attr {
		if a.Key == key {
			val, ok = a.Val, true
			break
		}
	}
	return
}

// Text returns the inner text of n.
func Text(n *html.Node) string {
	var b strings.Builder

	if n == nil {
		return ""
	}

	if n.Type == html.TextNode {
		return n.Data
	}

	for n := n.FirstChild; n != nil; n = n.NextSibling {
		switch n.Type {
		case html.TextNode:
			s := strings.Trim(n.Data, " \n|")
			b.WriteString(s)
		case html.ElementNode:
			switch n.Data {
			case "br":
				b.WriteString("\n")
			default:
				b.WriteString(Text(n))
			}
		}
	}

	return b.String()
}
