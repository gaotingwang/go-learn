package client

import (
	"net/rpc"

	"github.com/gaotingwang/go-learn/crawler_distributed/worker"

	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler_distributed/config"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	// 这里只是返回一个函数定义，返回的结果会被外面不断调用，相当于会不断 client := <-clientChan
	return func(request engine.Request) (engine.ParseResult, error) {
		req := worker.SerializedRequest(request)

		var result worker.ParseResult
		client := <-clientChan
		// jsonrpc 调用返回结果，result.Requests[0].Parser.Args 为interface{} 会丢失原有对象类型，变为map[string]interface{}
		err := client.Call(config.CrawlServiceRpc, req, &result)
		if err != nil {
			return engine.ParseResult{}, nil
		}

		return worker.DeserializedResult(result), nil
	}
}
