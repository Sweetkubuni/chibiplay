// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

func main() {
	response, err := http.Get("https://ww2.gogoanime2.org/search/pokemon")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	htmlTokens := html.NewTokenizer(response.Body)
loop:
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
		case html.StartTagToken:
			t := htmlTokens.Token()
			if t.Data == "a" {
				var href string
				for _, a := range t.Attr {
					if a.Key == "href" {
						href = a.Val
					}
				}
				re := regexp.MustCompile(`([a-zA-z0-9-]+)`)
				match := re.FindAllString(href, -1)
				if len(match) > 1 && match[1] == "anime" {
					fmt.Println("href: " + href + "  title: " + match[1])
				}
			}
		}
	}
}
