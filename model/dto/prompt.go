package model

import "github.com/easyship/model/vo"

type PromptDTO struct {
	SearchPromptList    []*vo.PromptVO `json:"search_prompt_list"`
	RecommendPromptList []*vo.PromptVO `json:"recommend_prompt_list"`
}
