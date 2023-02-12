package main

import (
	"fmt"

	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/scheduler"
	"github.com/gaotingwang/go-learn/crawler/zhenai/parser"
	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	itemsaver "github.com/gaotingwang/go-learn/crawler_distributed/persist/client"
	worker "github.com/gaotingwang/go-learn/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
