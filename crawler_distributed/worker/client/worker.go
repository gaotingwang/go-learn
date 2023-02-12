package client

import (
	"fmt"

	"github.com/gaotingwang/go-learn/crawler_distributed/worker"

	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(request engine.Request) (engine.ParseResult, error) {
		req := worker.SerializedRequest(request)

		var result worker.ParseResult
		// jsonrpc 调用返回结果，result.Requests[0].Parser.Args 为struct 会丢失原有对象类型，变为map[string]interface{}
		err = client.Call(config.CrawlServiceRpc, req, &result)
		if err != nil {
			return engine.ParseResult{}, nil
		}

		return worker.DeserializedResult(result), nil
	}, nil

}
