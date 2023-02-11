package engine

import (
	"github.com/gaotingwang/go-learn/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error: url %s : %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(content, r.Url), nil
}
