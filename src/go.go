package main

import (
	"errors"
	"net/http"
	"regexp"
)

func checkRedirect2(req *http.Request, via []*http.Request) error {
	re := "https?://github\\.com/"
	if matched, _ := regexp.MatchString(re, req.URL.String()); matched { // go/regex/missing-regexp-anchor
		return nil
	}
	return errors.New("Invalid URL.")
}
