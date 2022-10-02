package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Get(u string) error {
	// Build fileName from fullPath
	getURL, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	reqTime := time.Now()
	// Put content on file
	resp, err := client.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	respTime := time.Now()
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	latency := respTime.Sub(reqTime)
	size := len(b)
	fmt.Printf("Fetched the url of a file %s with size %d and latency %d", getURL, size, latency)
	return err
}
