package ai_model

import (
	"context"
	"github.com/easyship/infra"
	"github.com/easyship/infra/constant"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
)

const (
	FINISH_REASON = "finish_reason"
	OUTPUT        = "output"

	//finish reason val
	STOP   = "stop"
	NULL   = "null"
	LENGTH = "length" //过长
)

type QwenModel struct {
}

func (*QwenModel) HandleStreamRealData(ctx context.Context, resData string) (handleStreamResult *HandleStreamResult, err error) {
	handleStreamResult = &HandleStreamResult{}
	resMap, err := util.ParseJson(resData)
	if err != nil {
		log.ErrorWithContext(ctx, "ParseJson val:%v error:%v", resData, err)
		return nil, err
	}
	outputMap, err := handleTextRequest(resMap)
	if err != nil {
		log.ErrorWithContext(ctx, "handleTextRequest val:%v error:%v", util.ToJson(resMap), err)
		return nil, err
	}
	handleStreamResult.ModelRes, handleStreamResult.IsEnd = commonHandleText(outputMap)
	return handleStreamResult, nil
}

func handleTextRequest(requestRes map[string]interface{}) (map[string]interface{}, error) {
	// 调用失败
	if res, ok := requestRes[constant.CODE]; ok {
		errMsg, ok := requestRes[constant.CODE].(string)
		if !ok || len(errMsg) == 0 {
			errMsg = res.(string)
		}
		return nil, infra.NewSystemError(errMsg)
	}
	outputMap := requestRes[OUTPUT].(map[string]interface{})
	return outputMap, nil
}

func commonHandleText(outputMap map[string]interface{}) (res string, isEnd bool) {
	finishReason, ok := outputMap[FINISH_REASON].(string)
	if ok && finishReason == STOP {
		isEnd = true
	} else {
		// 显式赋值
		isEnd = false
	}
	res, ok = outputMap[constant.TEXT].(string)
	if !ok {
		// 显式赋值
		res = constant.EMPTY_STRING
	}
	return res, isEnd
}
