package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	"github.com/gaotingwang/go-learn/crawler_distributed/persist"
	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v6"
)

// 命令行参数
var port = flag.Int("port", 0, "the port for itemsaver listen on")

// 提供了一个对外的ItemSave的rpc接口
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	//Fatal，若有异常，则挂了,没有机会recover。panic还有recover的机会
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	//docker , sniff false
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServerRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
