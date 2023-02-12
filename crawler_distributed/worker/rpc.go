package worker

import "github.com/gaotingwang/go-learn/crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializedRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializedResult(engineResult)
	return nil
}
