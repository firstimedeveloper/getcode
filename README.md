# getcode

This api is used to retrieve all the code in a page given a url. 

I created this because I used a lot of articles/tutorials to set up things (a digitalocean server for example).

Using this tool will allow me to retrieve the code found in said articles so that I can use the code more easily for things I don't have in my long term memory.

## 
To use the api, enter the following in your terminal.
```bash
go get github.com/firstimedeveloper/getcode
```

An example use:
```go
import (	
    "fmt"
	"github.com/firstimedeveloper/getcode/api")

func main() {
	url := "https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04"
	list, err := api.NewArticle("pre", url)
	if err != nil {
		// handle error
	}

	for _, v := range list {
		fmt.Println(v)
	}
}
```
Output:
```
<pre class="code-pre command prefixed local-environment"><code><ul class="prefixed"><li class="line" data-prefix="$">ssh root@<span class="highlight">your_server_ip</span>
</li></ul></code></pre>
<pre class="code-pre super_user prefixed"><code><ul class="prefixed"><li class="line" data-prefix="#">adduser <span class="highlight">sammy</span>
</li></ul></code></pre>
...
```
