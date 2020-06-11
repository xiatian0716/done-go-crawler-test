package persist

import (
	"context"
	"go-crawler-test/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSaver(t *testing.T) {
	bookdetail := model.Bookdetail{
		BookName:  "BookName",
		Author:    "刘慈欣",
		Publicer:  "重庆出版社",
		Bookpages: 470,
		Price:     "32.00",
		Score:     "9.3",
		Into:      "三体人在利用魔法般的科技锁死了地球人的科学之后...",
	}

	// 测试save方法
	id, err := save(bookdetail)

	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.127.200:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("database").
		Type("table").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

}
