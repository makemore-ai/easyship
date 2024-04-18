package infra

import (
	"context"
	"github.com/easyship/util"
)

func Init(ctx context.Context) error {
	err := initEsClient(ctx)
	if err != nil {
		return err
	}
	util.InitEnv()
	return nil
}
