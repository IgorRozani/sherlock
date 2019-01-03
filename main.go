package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	sites := []string{
		"https://github.com/",
		"https://vimeo.com/",
		"https://www.instagram.com/",
		"https://www.twitter.com/",
		"https://www.facebook.com/",
		"https://medium.com/@",
	}

	c := make(chan string)

	for _, s := range sites {
		go checkWebSite(s+os.Args[1], c)
	}

	for i := 0; i < len(sites); i++ {
		fmt.Println(<-c)
	}
}

func checkWebSite(u string, c chan string) {
	resp, _ := http.Get(u)

	if resp.StatusCode != 200 {
		c <- u + ": It's free"
		return
	}
	c <- u + ": It's used"
}
