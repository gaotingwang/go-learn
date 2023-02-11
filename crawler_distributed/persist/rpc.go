package persist

import (
	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/persist"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (iss *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.SaveItem(iss.Client, iss.Index, item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v:%v", item, err)
	}
	return err
}
