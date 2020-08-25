package main

import (
	"fmt"
	"log"

	"github.com/firstimedeveloper/getcode/api"
)

func main() {
	url := "https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04"
	list, err := api.NewArticle("code", url)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range list {
		fmt.Println(v)
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
