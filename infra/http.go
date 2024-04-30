package infra

import (
	"context"
	"github.com/easyship/util/log"
	"net/http"
	"time"
)

type EasyHttpClient struct {
}

var DefaultHttpClient = &EasyHttpClient{}

func (*EasyHttpClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	duration := time.Since(start)
	log.InfoWithContext(ctx, "http request cost:%vms", duration.Milliseconds())
	return resp, err
}

func GetHttpClient() *EasyHttpClient {
	return DefaultHttpClient
}
