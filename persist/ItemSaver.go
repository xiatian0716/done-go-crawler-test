package persist

import (
	"context"
	"errors"
	"go-crawler-test/engine"
	"log"

	"github.com/olivere/elastic"
)

// insex 我这个ItemSaver往哪里存
func ItemSaver(index string) (chan engine.Item, error) {
	// save()里面每save一个item就需要开一个elastic服务
	// 我们希望的是什么呢？只有一个client处理item就可以了
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.127.200:9200"),
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	// 连不上返回err
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
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
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, insex string, item engine.Item) error {

	// 送上来的Type没有填怎么办？
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	// 送上来的Id没有填怎么办？
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	// 往里面存数据
	// Index-Type-Id定位一个元素
	// Index由程序的配置人员来配置
	// Type-Id两个维度由parse来给
	indexService := client.Index().
		Index(insex).
		Type(item.Type).
		BodyJson(item)

	// 送上来的Id没有填怎么办？
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
