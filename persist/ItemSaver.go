package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	// ItemSaver里面就可以做事情
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %s", itemCount, item)
			itemCount++

			// Elasticsearch Clients
			// Community Contributed Clients
			// Go -> elastic: Elasticsearch client for Google Go.
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	// 开一个elastic服务
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.127.200:9200"),
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return "", err
	}

	// 往里面存数据
	resp, err := client.Index().
		Index("database").
		Type("table").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
