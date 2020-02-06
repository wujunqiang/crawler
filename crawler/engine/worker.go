package engine

import (
	"crawler_company/crawler/fatcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {

	body, err := fatcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
