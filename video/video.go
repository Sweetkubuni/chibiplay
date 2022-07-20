package video

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func getStringFromURL(url, contains string, action func(string)) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(body))
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), contains) {
			action(scanner.Text())
		}
	}
}

// e.g "https://ww2.gogoanime2.org/watch/100-man-no-inochi-no-ue-ni-ore-wa-tatteiru/1"
func PlayVideo(url string) {

	getStringFromURL(url, `"playerframe"`, func(t string) {
		re := regexp.MustCompile(`"[^"]+"`)
		matches := re.FindAllString(t, -1)

		getStringFromURL("https://ww2.gogoanime2.org"+strings.Trim(matches[1], "\""), `"file":`, func(t string) {
			re := regexp.MustCompile(`'(.*)'`)
			match := re.FindString(t)
			m3u8File := strings.Trim(match, `'`)
			cmd := exec.Command("play.bat", "https://ww2.gogoanime2.org"+m3u8File)
			fmt.Println(cmd.Run())
		})
	})
}
