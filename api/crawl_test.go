package api

import (
	"testing"
)

func compareSlice(s1, s2 []string) bool {
	same := true
	if len(s1) == len(s2) {
		for i, v := range s1 {
			if s2[i] != v {
				same = false
				break
			}
		}
	}
	return same
}

func TestParseBody(t *testing.T) {
	var article Article
	err := article.parseHTML(htm, "code")
	if err != nil {
		t.Error(err)
	}
	var want []string
	want = append(want, "<code>get github.com/firstimedeveloper/getcode</code>")
	got := article.Lines
	if !compareSlice(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

}

const htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
	<p>more content</p>
	<code>get github.com/firstimedeveloper/getcode</code>
</body>
</html>`
