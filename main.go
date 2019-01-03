package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type site struct {
	Name string
	Link string
}

func main() {
	s := convertJSONToStruck()

	u := getUsername()

	fmt.Printf("🔎 Checking username \"%v\"\n", u)

	verifySites(s, u)
}

func getUsername() string {
	if len(os.Args) <= 1 {
		fmt.Println("Please inform the username")
		os.Exit(1)
	}

	return os.Args[1]
}

func convertJSONToStruck() []site {
	var s []site

	json.Unmarshal(readJSON(), &s)

	return s
}

func readJSON() []byte {
	bs, err := ioutil.ReadFile("sites.json")
	if err != nil {
		fmt.Println("File not found:", err)
		os.Exit(1)
	}
	return bs
}

func checkWebSite(s site, u string, c chan string) {
	resp, err := http.Get(s.Link + u)

	if err != nil {
		c <- s.Name + ": ❌ Verification fail"
		return
	}

	if resp.StatusCode != 200 {
		c <- s.Name + ": 👍 Free"
		return
	}
	c <- s.Name + ": 👎 Used"
}

func verifySites(sites []site, u string) {
	c := make(chan string)
	for _, s := range sites {
		go checkWebSite(s, u, c)
	}

	for i := 0; i < len(sites); i++ {
		fmt.Println(<-c)
	}
}
