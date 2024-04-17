package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/easyship/infra"
	"github.com/easyship/model/do"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
)

const (
	INDEX_NAME = "prompt"

	HITS = "hits"
)

const (
	//es mapping key
	ID         = "id"
	LABEL_NAME = "label_name"

	PROMPT_ZH = "prompt_zh"

	PROMPT_EN = "prompt_en"
)

func getHits(ctx context.Context, resp map[string]interface{}) []map[string]interface{} {
	hitsMap, ok := resp[HITS].(map[string]interface{})
	if !ok || len(hitsMap) == 0 {
		return []map[string]interface{}{}
	}
	tempHitsMapList, ok := hitsMap[HITS].([]interface{})
	if !ok || len(tempHitsMapList) == 0 {
		return []map[string]interface{}{}
	}
	realHitsMapList := make([]map[string]interface{}, 0, len(tempHitsMapList))
	for _, tempHitsMap := range tempHitsMapList {
		realHistsMap, ok := tempHitsMap.(map[string]interface{})
		if !ok || len(realHistsMap) == 0 {
			log.ErrorWithContext(ctx, "error transfer realHistsMap", util.ToJson(tempHitsMap))
			continue
		}
		realHitsMapList = append(realHitsMapList, realHistsMap)
	}
	return realHitsMapList
}

func commonHandleEsError(ctx context.Context, resp *esapi.Response, err error) error {
	if err != nil {
		log.ErrorWithContext(ctx, "esClient doQuery error : %v", err)
		return err
	}
	if resp == nil {
		err = infra.NewSystemError("es result nil")
		log.ErrorWithContext(ctx, "commonHandleEsError result: %v", err)
		return err
	}
	if resp.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			log.ErrorWithContext(ctx, "Error parsing the response body: %v", err)
			return err
		} else {
			errMsg := util.ToJson(e["error"])
			err = infra.NewSystemError(errMsg)
			log.ErrorWithContext(ctx, "esClient doQuery result error : %s", errMsg)
			return err
		}
	}
	return nil
}

func SearchPrompt(ctx context.Context, searchText string) ([]*do.PromptEsDo, error) {
	esClient := infra.GetEsClient()
	var body bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{"match": map[string]interface{}{LABEL_NAME: searchText}},
					{"match": map[string]interface{}{PROMPT_ZH: searchText}},
				},
			},
		},
		"sort": []map[string]interface{}{
			{
				"_score": map[string]interface{}{
					"order": "desc",
				}},
		},
	}
	if err := json.NewEncoder(&body).Encode(query); err != nil {
		log.ErrorWithContext(ctx, "Error encoding query: %v", err)
		return nil, err
	}
	searchRes, err := esClient.Search(esClient.Search.WithContext(ctx), esClient.Search.WithIndex(INDEX_NAME),
		esClient.Search.WithBody(&body), esClient.Search.WithPretty())
	err = commonHandleEsError(ctx, searchRes, err)
	if err != nil {
		log.ErrorWithContext(ctx, "do SearchPrompt searchRequest:%v, error:%v", util.ToJson(query), err)
		return nil, err
	}
	// 延迟执行
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.ErrorWithContext(ctx, "do SearchPrompt close body error:%v", err)
		}
	}(searchRes.Body)
	res := make(map[string]interface{})
	if err = json.NewDecoder(searchRes.Body).Decode(&res); err != nil {
		log.ErrorWithContext(ctx, "parse searchRes body error: %v", err)
		return nil, err
	}
	realHitsMapList := getHits(ctx, res)
	promptDoList := make([]*do.PromptEsDo, 0, len(realHitsMapList))
	for _, resMap := range realHitsMapList {
		source, ok := resMap["_source"].(map[string]interface{})
		if !ok || len(source) == 0 {
			continue
		}
		promptDoList = append(promptDoList, &do.PromptEsDo{
			Id:        int64(source[ID].(float64)),
			PromptZh:  util.StrPtr(source[PROMPT_ZH].(string)),
			PromptEn:  util.StrPtr(source[PROMPT_EN].(string)),
			LabelName: source[LABEL_NAME].(string),
		})
	}
	return promptDoList, nil
}
