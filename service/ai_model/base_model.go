package ai_model

import (
	"context"
	"github.com/easyship/infra"
	"github.com/easyship/infra/constant"
	"github.com/easyship/model"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"strings"
)

type HandleStreamResult struct {
	// ke为ID, DATA, EVENT
	StreamRes *model.StreamResult
	ModelRes  string
	IsEnd     bool
}

type BaseModel interface {
	HandleStreamRealData(ctx context.Context, resData string) (handleStreamResult *HandleStreamResult, err error)
}

// HandleStream
// readBytes示例
// data:{"output":
//
//id:1
//event:result
func HandleStream(ctx context.Context, aiModel BaseModel, readBytes []byte) (handleStreamResult *HandleStreamResult, err error) {
	if len(readBytes) == 0 {
		return handleStreamResult, nil
	}
	streamRes := string(readBytes)
	dataList := strings.Split(streamRes, "\n")
	streamResMap := &model.StreamResult{}
	isEnd := false
	log.InfoWithContext(ctx, "dataList:%v", streamRes)
	for _, data := range dataList {
		if data == constant.EMPTY_STRING || len(strings.TrimSpace(data)) == 0 {
			continue
		}
		// 只拆解第一个:
		kvRes := strings.SplitN(data, constant.COLON, 2)
		if len(kvRes) != 2 {
			continue
		}
		switch kvRes[0] {
		case constant.ID:
			// id在上游处理，
			// 可能需要避免错位，暂不处理
			break
		case constant.EVENT:
			//存在错误
			if kvRes[1] != constant.RESULT {
				log.ErrorWithContext(ctx, "HandlePromptStreamData error:%v", streamRes)
				streamResMap.Event = constant.ERROR_EVENT
				err = infra.NewDefaultSystemError()
				break
			}
			streamResMap.Event = constant.CONTINUE_EVENT
		case constant.DATA:
			tempHandleStreamResult, err := aiModel.HandleStreamRealData(ctx, kvRes[1])
			if err != nil {
				log.ErrorWithContext(ctx, "HandleStreamRealData value:%v, error:%v", kvRes[1], err)
				break
			}
			if tempHandleStreamResult.IsEnd {
				// 已经结束 停止
				streamResMap.Event = constant.STOP_EVENT
				isEnd = tempHandleStreamResult.IsEnd
			}
			streamResMap.Data = util.StrPtr(tempHandleStreamResult.ModelRes)
		}
		if err != nil {
			break
		}
	}
	handleStreamResult = &HandleStreamResult{
		StreamRes: streamResMap,
		IsEnd:     isEnd,
	}
	return handleStreamResult, nil
}
