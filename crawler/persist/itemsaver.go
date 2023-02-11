package persist

import (
	"errors"
	"github.com/gaotingwang/go-learn/crawler/engine"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func ItemSaver(itemIndex string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Got item: count=%d, content = %v\n", itemCount, item)
			itemCount++

			err := SaveItem(client, itemIndex, item)
			if err != nil {
				log.Printf("Item save error : %+v : %v \n", item, err)
			}
		}
	}()
	return out, nil
}

func SaveItem(client *elastic.Client, itemIndex string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply item type")
	}

	indexService := client.Index().
		Index(itemIndex).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
