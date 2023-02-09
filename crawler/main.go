package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// 单任务爬虫
	//engine.SingleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenhun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 并发爬虫,实现1
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenhun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 实现2
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenhun",
		ParserFunc: parser.ParseCityList,
	})
}
