// You can edit this code!
// Click here and start typing.
package anime

import (
	"regexp"

	"test_client/request"

	"golang.org/x/net/html"
)

type Anime struct {
	href  string
	title string
}

// e.g https://ww2.gogoanime2.org/search/pokemon

func SearchAnime(url string) []Anime {

	var resp []Anime

	request.GetHyperLinks(url, func(t html.Token) {
		anime := Anime{}
		for _, a := range t.Attr {
			if a.Key == "href" {
				anime.href = a.Val
			}

			if a.Key == "title" {
				anime.title = a.Val
			}
		}
		re := regexp.MustCompile(`/(.*)/`)
		match := re.FindStringSubmatch(anime.href)
		if len(match) > 1 && match[1] == "anime" {
			resp = append(resp, anime)
		}
	})

	return resp
}