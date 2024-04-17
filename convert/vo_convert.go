package convert

import (
	"github.com/easyship/model/do"
	"github.com/easyship/model/vo"
)

func ConvertPromptDO2VO(promptEsDo *do.PromptEsDo) *vo.PromptVO {
	if promptEsDo == nil {
		return nil
	}
	return &vo.PromptVO{
		Id:        promptEsDo.Id,
		PromptZh:  promptEsDo.PromptZh,
		PromptEn:  promptEsDo.PromptEn,
		LabelName: promptEsDo.LabelName,
	}
}

func ConvertPromptDO2VOList(promptEsDoList []*do.PromptEsDo) []*vo.PromptVO {
	if len(promptEsDoList) == 0 {
		return []*vo.PromptVO{}
	}
	promptVoList := make([]*vo.PromptVO, 0, len(promptEsDoList))
	for _, promptEsDo := range promptEsDoList {
		promptVoList = append(promptVoList, ConvertPromptDO2VO(promptEsDo))
	}
	return promptVoList
}
