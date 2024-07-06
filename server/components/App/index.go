package components

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
)

type Components interface {
	Render() string
}

func Duder[P any](state P) string {
	markup := `
		<div>
			<h1>Hey Duder</h1>
		</div>
	`
	return RenderComponent[P](markup, state)
}

func RenderComponent[S any](markup string, state S) string {
	componentsMap := map[string]func(state S) string {
		"Duder": Duder[S],
	}

	// needs to be a synatx tree traversal
	htmlLines := strings.Split(markup, "\n")

	newHTMLLines := make([]string, 0)

	for _, line := range htmlLines {

		replacers := []string{ "<", ">", "/", "\t" }

		tagName := line

		for _, token := range replacers {
			tagName = strings.ReplaceAll(tagName, token, "")			
		}
		tagName = strings.TrimSpace(tagName)
		
		componentFunction, present := componentsMap[tagName]

		if !present {
			newHTMLLines = append(newHTMLLines, line)
		} else {
			componentHTML := componentFunction(state)
			newHTMLLines = append(newHTMLLines, strings.Split(componentHTML, "\n")...)
		}
	}
	return strings.Join(newHTMLLines, "\n")
}

func App[P any](state P) string {
	fmt.Println("App!")
	markup := `
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
	return RenderComponent[P](markup, state)
}

func ParseHTML(htmlString string) {

	r := strings.NewReader(htmlString)

	doc, err := html.Parse(r)
	if err != nil {
		// ...
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// Do something with n...
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
