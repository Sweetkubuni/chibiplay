package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	{
		resp, err := http.Get("https://ww2.gogoanime2.org/watch/100-man-no-inochi-no-ue-ni-ore-wa-tatteiru/1")
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
			if strings.Contains(scanner.Text(), `"playerframe"`) {
				fmt.Println(scanner.Text())
			}
		}
	}

	{
		resp, err := http.Get("https://ww2.gogoanime2.org/embed/MTQ2MDUz")
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
			if strings.Contains(scanner.Text(), `"file":`) {
				fmt.Println(scanner.Text())
			}
		}
	}
}
