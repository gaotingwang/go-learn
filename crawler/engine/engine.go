package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		content, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch error: url %s : %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(content)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
