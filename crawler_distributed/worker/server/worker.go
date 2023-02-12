package main

import (
	"fmt"
	"log"

	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"
	"github.com/gaotingwang/go-learn/crawler_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
