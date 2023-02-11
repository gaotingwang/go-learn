package main

import (
	"fmt"
	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/scheduler"
	"github.com/gaotingwang/go-learn/crawler/zhenai/parser"
	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	"github.com/gaotingwang/go-learn/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
