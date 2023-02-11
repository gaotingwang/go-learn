package engine

import (
	"log"

	"github.com/gaotingwang/go-learn/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error: url %s : %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parser(content, r.Url), nil
}
