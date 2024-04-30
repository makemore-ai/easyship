package model

import "github.com/easyship/model/vo"

type PromptSearchDTO struct {
	SearchPromptList    []*vo.PromptVO `json:"search_prompt_list"`
	RecommendPromptList []*vo.PromptVO `json:"recommend_prompt_list"`
}

type PromptRefreshDTO struct {
	Prompt            string   `json:"prompt"`
	PromptContextList []string `json:"prompt_context_list"`
}
