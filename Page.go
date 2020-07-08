package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	subdHeaders = map[string]string{
		"User-Agent": "🦊Mozilla🦊/5.0 (X11; 🐧Linux🐧 x86_64; rv:68.0) 🦎Gecko🦎/20100101 🔥Firefox🔥/68.0"}
)

func getPage(url string, payload string) []byte {

	req, _ := http.NewRequest("GET", url+payload, nil)

	req.Header.Add("User-Agent", subdHeaders["User-Agent"])

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return body
}
