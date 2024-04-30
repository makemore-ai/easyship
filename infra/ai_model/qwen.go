package ai_model

import (
	"context"
	"github.com/easyship/infra"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"io"
	"net/http"
)

const QWEN_API_KEY = "sk-Z2B6CJsFNA"

const (
	QWEN_URL_PREFIX = "https://dashscope.aliyuncs.com"

	QWEN_TEXT_URL = QWEN_URL_PREFIX + "/api/v1/services/aigc/text-generation/generation"
)

// QwenTextRequestWithSSE SSE长连接Qwen
func QwenTextRequestWithSSE(ctx context.Context, prompt string) (io.ReadCloser, error) {
	req, err := buildQwenTextRequest(ctx, prompt)
	// 添加SSE Header
	req.Header.Add("X-DashScope-SSE", "enable")
	if err != nil {
		log.ErrorWithContext(ctx, "QwenTextRequestWithSSE req error:%v", err)
	}
	resp, err := infra.DefaultHttpClient.Do(ctx, req)
	if err != nil {
		log.ErrorWithContext(ctx, "QwenTextRequestWithSSE Do error:%v", req)
		return nil, err
	}
	return resp.Body, nil
}

func QwenTextRequest(ctx context.Context, prompt string) (map[string]interface{}, error) {
	req, err := buildQwenTextRequest(ctx, prompt)
	if err != nil {
		log.ErrorWithContext(ctx, "QwenTextRequest req error:%v", err)
		return nil, err
	}
	resp, err := infra.DefaultHttpClient.Do(ctx, req)
	if err != nil {
		log.ErrorWithContext(ctx, "QwenTextRequest Do error:%v", req)
		return nil, err
	}
	resMap, err := util.DecodeJson(resp.Body)
	log.InfoWithContext(ctx, "resp:%v,%v", resMap, err)
	return resMap, nil
}

func buildQwenTextRequest(ctx context.Context, prompt string) (*http.Request, error) {
	requestParams := map[string]interface{}{
		//先默认使用turbo了，其他太贵了
		"model": "qwen-turbo",
		"input": map[string]interface{}{
			"prompt": prompt,
		},
		"parameters": map[string]interface{}{
			"top_p": 0.1,
		},
	}
	req, err := http.NewRequest("POST", QWEN_TEXT_URL, util.ParseBody(requestParams))
	if err != nil {
		log.ErrorWithContext(ctx, "buildQwenTextRequest error:%v", req)
		return nil, err
	}
	req.Header.Add("Authorization", QWEN_API_KEY)
	req.Header.Add(util.CONTENT_TYPE, util.APPLICATION_JSON)
	return req, nil
}
