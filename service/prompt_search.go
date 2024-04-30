package service

import (
	"context"
	"github.com/easyship/config"
	"github.com/easyship/convert"
	"github.com/easyship/infra/dao"
	dto_model "github.com/easyship/model/dto"
	"github.com/easyship/model/vo"
	"github.com/easyship/util/log"

	"strings"
)

func SearchPrompt(ctx context.Context, searchText *string) (*dto_model.PromptSearchDTO, error) {
	doQuery := !(searchText == nil || len(strings.TrimSpace(*searchText)) == 0)
	var promptList []*vo.PromptVO
	if doQuery {
		// 没做searchText判空，是前面doQuery会自动为false
		promptDoList, err := dao.SearchPrompt(ctx, *searchText)
		if err != nil {
			log.ErrorWithContext(ctx, "SearchPrompt error:%v", err)
			return nil, err
		}
		promptList = convert.ConvertPromptDO2VOList(promptDoList)
	} else {
		promptList = make([]*vo.PromptVO, 0, 0)
	}

	recommendPromptList := config.GetRecommendPromptList(ctx)

	return &dto_model.PromptSearchDTO{
		SearchPromptList:    promptList,
		RecommendPromptList: recommendPromptList,
	}, nil
}
