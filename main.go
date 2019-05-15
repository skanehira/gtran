package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// required to translate language flags
var (
	// source language
	source = flag.String("source", "en", "translate source")
	// target language
	target = flag.String("target", "ja", "translate traget")
	// source language text
	text = flag.String("text", "", "translate source text")
	// Use the Google Apps Script to translate language
	endpoint = flag.String("endpoint", "https://script.google.com/macros/s/AKfycbywwDmlmQrNPYoxL90NCZYjoEzuzRcnRuUmFCPzEqG7VdWBAhU/exec", "translate endpoint")
)

type post struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

// translate language
func translate(text, source, target string) (string, error) {
	postData, err := json.Marshal(post{text, source, target})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, *endpoint, bytes.NewBuffer([]byte(postData)))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func run(args []string) int {
	envEndpoint := os.Getenv("GTRAN_ENDPOINT")
	if envEndpoint != "" {
		*endpoint = envEndpoint
	}

	if len(args) == 0 && *text == "" {
		flag.Usage()
		return -1
	}

	if *text == "" && args[0] != "" {
		*text = args[0]
	}

	result, err := translate(*text, *source, *target)
	if err != nil {
		return -1
	}
	fmt.Println(result)
	return 0
}

func main() {
	flag.Parse()
	os.Exit(run(flag.Args()))
}
