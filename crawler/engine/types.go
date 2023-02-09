package engine

import (
	"crawler/fetcher"
	"log"
)

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

func worker(r Request) (ParseResult, error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error: url %s : %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(content), nil
}
