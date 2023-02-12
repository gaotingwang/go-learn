package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"
	"github.com/gaotingwang/go-learn/crawler_distributed/worker"
)

// 17-7
func TestCrawService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second) //保证服务起来
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://www.zhenai.com/zhenghun/beijing",
		Parser: worker.SerializedParser{
			Name: config.ParseCity,
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
