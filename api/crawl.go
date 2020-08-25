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
	str, err := respToStr(url)
	if err != nil {
		return nil, err
	}
	err = article.parseHTML(str, tag)
	if err != nil {
		return nil, err
	}
	return article.Lines, nil
}

// Article struct
type Article struct {
	Lines []string `xml:"code"`
}

func respToStr(url string) (string, error) {
	resp, err := http.Get("https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04")
	if err != nil {
		return "", err
	}
	parsedBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(parsedBody), nil
}

func node(body string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Body retrieves the html nodes, converts them to strings
func (a *Article) parseHTML(s string, tag string) error {
	doc, err := node(s)
	if err != nil {
		return err
	}
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
