# Http client package to interract with [short.io api]([https://h.com/](https://developers.short.io/reference/apilinksget))

### Implemented functions:

```
1. Shorten link (simple to use returns only shorten link)
2. Shorten links bulk (returns response with all necessary info)
```

### Setting up

Set on your `.env` file variables 
SHORTENER_DOMAIN
SHORTENER_TOKEN 
To get your domain and token you must be registered in short.io

### Example usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/salesforceanton/short-io-go-client/shortener"
)

func main() {
	// Create shortener client
	client, err := shortener.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Prepare link to shorten
	originalLink := "https://sumo9-dev-ed.my.site.com/sumo/s/onlinescheduler?processId=a0T5j000003WLJEEA4&clientId=0035j00000v8QaPAAU"
	shortLink, err := client.ShortenLink(originalLink)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Get the result
	fmt.Println(shortLink)
```
