package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	// Use the Google Apps Script to translate language
	endpoint = "https://script.google.com/macros/s/AKfycbzi15QCo0IsjutiMnI5FYf43-TKqfrUDiaM03x5C5IcH7-setg/exec?"
)

// required to translate language flags
var (
	// source language
	source = flag.String("source", "en", "translate source")
	// target language
	target = flag.String("target", "ja", "translate traget")
	// source language text
	text = flag.String("text", "", "translate source text")
)

// translate language
func translate() (string, error) {
	v := url.Values{}
	v.Add("text", *text)
	v.Add("source", *source)
	v.Add("target", *target)

	resp, err := http.Get(endpoint + v.Encode())
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func run() int {
	flag.Parse()
	if *text == "" {
		flag.Usage()
		return -1
	}

	result, err := translate()
	if err != nil {
		return -1
	}
	fmt.Println(result)
	return 0
}

func main() {
	os.Exit(run())
}
