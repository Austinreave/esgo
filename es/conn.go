package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

var (
	Client *elastic.Client
	host  ="http://127.0.0.1:9200/"
	user = "elastic"
	pwd = "admin888"
)

// 初始化连接
func Init() (err error) {
	Client, err = elastic.NewSimpleClient(
		elastic.SetURL(host),
		elastic.SetBasicAuth(user,pwd),
	)
	if err != nil {
		fmt.Println("错误："+err.Error())
		return
	}
	info, code, err := Client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := Client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch version %s\n", esversion)
	return nil
}
