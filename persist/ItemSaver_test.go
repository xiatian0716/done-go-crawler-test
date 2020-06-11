package persist

import (
	"context"
	"go-crawler-test/engine"
	"go-crawler-test/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "https://book.douban.com",
		Type: "douban_test",
		Id:   "001",

		Payload: model.Bookdetail{
			BookName:  "BookName",
			Author:    "刘慈欣",
			Publicer:  "重庆出版社",
			Bookpages: 470,
			Price:     "32.00",
			Score:     "9.3",
			Into:      "三体人在利用魔法般的科技锁死了地球人的科学之后...",
		},
	}

	// 然后我们去拿出来
	// Fetch saved item
	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.127.200:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	// 测试save方法
	// Save expected item
	const index = "database"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

}
