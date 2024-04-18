package infra

import (
	"context"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"github.com/elastic/go-elasticsearch/v7"
)

var ES_CLIENT *elasticsearch.Client

func initEsClient(ctx context.Context) error {
	var esClientAddresses []string
	if util.IsProd() {
		esClientAddresses = []string{"http://172.18.26.245:9200"}
	} else {
		esClientAddresses = []string{"http://localhost:9200"}
	}

	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: esClientAddresses,
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
