package api

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// NewArticle returns slice of strings or an error upon recieving the tag and url
// from which the user wishes to retrieve.
func NewArticle(tag, url string) ([]string, error) {
	var article Article
	resp, err := http.Get("https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04")
	if err != nil {
		return nil, err
	}
	parsedBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(strings.NewReader(string(parsedBody)))
	err = article.Body(doc, tag)
	if err != nil {
		return nil, err
	}
	return article.Lines, nil
}

// Article struct
type Article struct {
	Lines []string `xml:"code"`
}

// Body retrieves the html nodes, converts them to strings
func (a *Article) Body(doc *html.Node, tag string) error {
	var body []*html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			body = append(body, node)
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	for _, v := range body {
		a.Lines = append(a.Lines, renderNode(v))
	}
	if body != nil {
		return nil
	}
	return errors.New("Missing <code> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
