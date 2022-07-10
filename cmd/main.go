package main

import (
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
		{Text: "list", Description: "find content to play"},
		{Text: "play", Description: "plays media "},
	}

	listMenu []prompt.Suggest = []prompt.Suggest{
		{Text: "--episodes", Description: "find content to play"},
		{Text: "--movies", Description: "plays media "},
		{Text: "--new-season", Description: "find content to play"},
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
		fmt.Println("you called search")
	}

	//LivePrefixState.LivePrefix = in + "> "
	//LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	text := strings.TrimSpace(in.Text)
	blocks := strings.Split(text, " ")
	switch blocks[0] {
	case "list":
		return prompt.FilterHasPrefix(listMenu, in.GetWordBeforeCursor(), true)
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
