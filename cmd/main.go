package main

import (
	"chibiplay/anime"
	"chibiplay/episode"
	"chibiplay/video"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

var (
	mainMenu []prompt.Suggest = []prompt.Suggest{
		{Text: "search", Description: "find content to play"},
		{Text: "play", Description: "plays media "},
	}

	searchMenu []prompt.Suggest = []prompt.Suggest{
		{Text: "--episodes", Description: "list episodes"},
		{Text: "--only-title", Description: "only show the title"},
		{Text: "--only-href", Description: "only show the url"},
	}

	playMenu []prompt.Suggest = []prompt.Suggest{
		{Text: "--episode", Description: "find content to play"},
		{Text: "--all", Description: "plays media "},
	}
)

func executor(in string) {
	in = strings.TrimSpace(in)

	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}

	blocks := strings.Split(in, " ")

	switch blocks[0] {
	case "search":
		modifiers := map[string]bool{
			"--episodes":   true,
			"--only-href":  true,
			"--only-title": true,
		}
		switch len_blocks := len(blocks); {
		case len_blocks > 3:
			fmt.Println("search with 4 arguments are not support yet!")
			return
		case len_blocks == 3 && modifiers[blocks[1]] && blocks[1] == "--episodes":
			episodes := episode.SearchEpisode("https://ww2.gogoanime2.org/" + blocks[2])
			for _, epi := range episodes {
				fmt.Println(epi.Href)
			}
			/* 		case len_blocks == 3 && modifiers[blocks[1]] && blocks[1] == "--only-title":
			   			animes := anime.SearchAnime("https://ww2.gogoanime2.org/search/" + blocks[2])
			   			for _, anime := range animes {
			   				fmt.Println(anime.Title)
			   			}
			   		case len_blocks == 3 && modifiers[blocks[1]] && blocks[1] == "--only-href":
			   			animes := anime.SearchAnime("https://ww2.gogoanime2.org/search/" + blocks[2])
			   			for _, anime := range animes {
			   				fmt.Println(anime.Href)
			   			} */
		case len_blocks == 2:
			animes := anime.SearchAnime("https://ww2.gogoanime2.org/search/" + blocks[1])
			for _, anime := range animes {
				fmt.Println(anime.Title, " -- ", anime.Href)
			}
		default:
			fmt.Println("search has incorrect number of arguments")
			fmt.Println("search <anime title>")
			fmt.Println("search --only-href <anime title>")
			fmt.Println("search --only-title <anime title>")
			fmt.Println("search --episodes <anime url>")
		}
	case "play":
		video.PlayVideo("https://ww2.gogoanime2.org" + blocks[1])

	}

	//LivePrefixState.LivePrefix = in + "> "
	//LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	text := strings.TrimSpace(in.Text)
	blocks := strings.Split(text, " ")
	switch blocks[0] {
	case "search":
		return prompt.FilterHasPrefix(searchMenu, in.GetWordBeforeCursor(), true)
	case "play":
		return prompt.FilterHasPrefix(playMenu, in.GetWordBeforeCursor(), true)
	default:
		return prompt.FilterHasPrefix(mainMenu, in.GetWordBeforeCursor(), true)
	}
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func Request(url string, handle func(r io.Reader) ([]string, error)) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	return handle(resp.Body)
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	p.Run()
}
