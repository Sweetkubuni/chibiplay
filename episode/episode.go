package episode

import (
	"regexp"
	"test_client/request"

	"golang.org/x/net/html"
)

type Episode struct {
	Href    string
	Episode string
}

//e.g https://ww2.gogoanime2.org/anime/digimon-adventure-2020

func SearchEpisode(url string) []Episode {
	var resp []Episode

	request.GetHyperLinks(url, func(t html.Token) {
		episode := Episode{}
		for _, a := range t.Attr {
			if a.Key == "href" {
				episode.Href = a.Val
			}
		}
		re := regexp.MustCompile(`([a-zA-z0-9.-]+)`)
		match := re.FindAllString(episode.Href, -1)
		if len(match) == 3 && match[0] == "watch" {
			episode.Episode = match[2]
			resp = append(resp, episode)
		}
	})

	return resp
}
