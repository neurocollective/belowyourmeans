package components

import (
	"testing"
	"strings"
	"golang.org/x/net/html"
	"fmt"
)

func TestAppOne(t *testing.T) {
	component := App[any](nil)

	expected := App[any](nil)

	if expected != component {
		t.Error("did not equal!")
	}
}

// func TraverseAndPrint(node *html.Node, level int) {

// }

func Traverse(n *html.Node, level int) {
	if n.Type == html.ElementNode {

		fmt.Println("level:", level)
		fmt.Println("tag:", n.Data)

		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("href:", a.Val)
					break
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Traverse(c, level + 1)
	}
}

func TestParseHTML(t *testing.T) {

	htmlString := `
		<html>
			<head>
			</head>
			<body>
				<header>
				</header>
				<main>
					<h1>Hey Bud</h1>
					<Duder />
				</main>
				<footer>
				</footer>
			</body>
		</html>
	`
	newHTML := htmlString

	replacers := []string{ "\n", "\t", " " }

	for _, token := range replacers {
		newHTML = strings.ReplaceAll(newHTML, token, "")
	}
	newHTML = strings.TrimSpace(newHTML)

	t.Log("htmlString:", htmlString)
	t.Log("newHTML:", newHTML)

	rootNode, err := html.Parse(strings.NewReader(htmlString))

	if err != nil {
		t.Error(err)
	}

	Traverse(rootNode, 0)
}

func TestParseHTMLTwo(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	Traverse(doc, 0)
}