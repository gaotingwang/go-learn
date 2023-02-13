package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"

	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/scheduler"
	"github.com/gaotingwang/go-learn/crawler/zhenai/parser"
	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	itemsaver "github.com/gaotingwang/go-learn/crawler_distributed/persist/client"
	worker "github.com/gaotingwang/go-learn/crawler_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "",
		"itemsaver host")
	workerHosts = flag.String("worker_hosts", "",
		"worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if err != nil {
			log.Printf("Error connecting to %s: %v", host, err)
		} else {
			clients = append(clients, client)
		}
	}

	pool := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				pool <- client
			}
		}
	}()

	return pool
}
