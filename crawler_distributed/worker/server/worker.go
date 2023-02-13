package main

import (
	"flag"
	"fmt"
	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"
	"github.com/gaotingwang/go-learn/crawler_distributed/worker"
	"log"
)

// 命令行参数
var port = flag.Int("port", 0, "the port for worker listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
