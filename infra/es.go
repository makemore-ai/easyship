package infra

import (
	"context"
	"github.com/easyship/util/log"
	"github.com/elastic/go-elasticsearch/v7"
)

var ES_CLIENT *elasticsearch.Client

func initEsClient(ctx context.Context) error {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.ErrorWithContext(ctx, "elasticsearch.NewTypedClient failed, err:%v", err)
		return err
	}
	ES_CLIENT = client
	return nil
}

func GetEsClient() *elasticsearch.Client {
	return ES_CLIENT
}
