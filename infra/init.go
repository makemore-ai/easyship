package infra

import (
	"context"
)

func Init(ctx context.Context) error {
	err := initEsClient(ctx)
	if err != nil {
		return err
	}
	return nil
}
