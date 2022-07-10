package request

import (
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func Request(url string, action func(io.ReadCloser)) {
	response, err := http.Get(url)
	if err != nil {
		panic("could not connect!")
	}
	defer response.Body.Close()
	//Do something with request
	action(response.Body)
}

func GetHyperLinks(url string, action func(html.Token)) {
	Request(url, func(body io.ReadCloser) {
		htmlTokens := html.NewTokenizer(body)
	loop:
		for {
			tt := htmlTokens.Next()
			switch tt {
			case html.ErrorToken:
				break loop
			case html.StartTagToken:
				t := htmlTokens.Token()
				if t.Data == "a" {
					action(t)
				}
			}
		}
	})
}
