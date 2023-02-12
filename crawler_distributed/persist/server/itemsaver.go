package main

import (
	"fmt"
	"log"

	"github.com/gaotingwang/go-learn/crawler_distributed/config"
	"github.com/gaotingwang/go-learn/crawler_distributed/persist"
	"github.com/gaotingwang/go-learn/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v6"
)

// 提供了一个对外的ItemSave的rpc接口
func main() {
	//Fatal，若有异常，则挂了,没有机会recover。panic还有recover的机会
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
