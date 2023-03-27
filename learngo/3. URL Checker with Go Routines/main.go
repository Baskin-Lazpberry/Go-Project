package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status string
}

func main() {
	c := make(chan urlStatus)
	urls := []string{
		"https://www.youtube.com/",
		"https://www.google.com/",
		"https://www.npr.org/",
		"https://github.com/",
		"https://paperswithcode.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://illinois.edu/",
		"https://www.stanford.edu/",
	}
	results := make(map[string]string)

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, stat := range results {
		fmt.Println(url, stat)
	}
}

func hitURL(url string, c chan<- urlStatus) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- urlStatus{url: url, status: "FAILED"}
	} else {
		c <- urlStatus{url: url, status: "OK"}
	}
}
