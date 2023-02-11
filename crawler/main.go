package main

import (
	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/persist"
	"github.com/gaotingwang/go-learn/crawler/scheduler"
	"github.com/gaotingwang/go-learn/crawler/zhenai/parser"
)

func main() {
	// 单任务爬虫
	//engine.SingleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 并发爬虫,实现1
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 实现2
	itemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemSaver,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
