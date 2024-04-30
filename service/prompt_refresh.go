package service

import (
	"context"
	"fmt"
	"github.com/easyship/infra/ai_model"
	"github.com/easyship/infra/constant"
	"github.com/easyship/infra/prompt"
	"github.com/easyship/model"
	ai_model2 "github.com/easyship/service/ai_model"
	"github.com/easyship/util"
	"github.com/easyship/util/log"
	"io"
	"strconv"
	"strings"
	"sync/atomic"
)

// AI_MODULE_HANDLE_MAP 考虑下用id吧 感觉字符串会有大小写问题
var AI_MODULE_HANDLE_MAP = map[int32]ai_model2.BaseModel{
	int32(1688): &ai_model2.QwenModel{},
}

var DEFAULT_HANDLE_MOUDLE = &ai_model2.QwenModel{}

func RefreshPromptWithSSE(ctx context.Context, reqPrompt string, aiModuleId int32) (func(w io.Writer) bool, error) {
	handleModule, ok := AI_MODULE_HANDLE_MAP[aiModuleId]
	if !ok {
		handleModule = DEFAULT_HANDLE_MOUDLE
	}
	respReader, err := ai_model.QwenTextRequestWithSSE(ctx, fmt.Sprintf(prompt.REFRESH_PROMPT, reqPrompt))
	if err != nil {
		log.ErrorWithContext(ctx, "QwenTextRequestWithSSE error:%v", err)
		return nil, err
	}
	// 暂时不会有并发问题
	idCal := model.IdCal{
		NowId: util.Int32Ptr(1),
	}
	return func(w io.Writer) bool {
		// 延迟推出close

		isEnd := false
		resBody := make([]byte, 4096)
		// 用于暂存处理异常
		var handleErr error

		_, err := respReader.Read(resBody)
		if err != nil {
			if err != io.EOF {
				log.ErrorWithContext(ctx, "QwenTextRequestWithSSE read resBody error:%v", err)
				handleErr = err
			} else {
				// 清楚EOF error
				isEnd = true
			}
		}
		// 没得到数据
		if len(resBody) == 0 {
			log.ErrorWithContext(ctx, "QwenTextRequestWithSSE read resBody empty")
			return true
		}
		handleStreamResult := &ai_model2.HandleStreamResult{}
		// 没发生异常
		if handleErr == nil {
			handleStreamResult, err = ai_model2.HandleStream(ctx, handleModule, resBody)
		}
		if err != nil {
			// 中止
			log.ErrorWithContext(ctx, "HandlePromptStreamData error:%v", err)
			handleErr = err
		}
		isEnd = isEnd || handleStreamResult.IsEnd
		resStr := buildStreamData(ctx, *idCal.NowId, handleStreamResult.StreamRes, isEnd, handleErr)
		log.InfoWithContext(ctx, "resStr:%v", resStr)
		_, err = w.Write([]byte(resStr))
		if err != nil || handleErr != nil {
			log.ErrorWithContext(ctx, "Write res error:%v, handleErr:%v", err, handleErr)
			// 关闭
			defer util.CloseReader(respReader)
			return false
		}
		if isEnd {
			//标志结束
			defer util.CloseReader(respReader)
			return false
		}
		// 先拿CAS加着
		idCal.NowId = util.Int32Ptr(atomic.AddInt32(idCal.NowId, 1))
		return true
	}, nil
}

func buildStreamData(ctx context.Context, streamId int32, streamResult *model.StreamResult, isEnd bool, err error) string {
	if streamResult == nil {
		streamResult = &model.StreamResult{}
	}
	streamResult.Id = streamId
	var data *model.CommonResponse
	// 返回失败
	if err != nil {
		data = util.ErrResp(err)
		streamResult.Event = constant.ERROR_EVENT
	} else {
		data = util.SuccessResp(streamResult.Data)
	}

	if isEnd {
		streamResult.Event = constant.STOP_EVENT
	}
	streamResult.Data = util.StrPtr(util.ToJson(data))
	//通过stream result 构建SSE返回流的string值
	streamResLine1 := strings.Join([]string{constant.ID, strconv.FormatInt(int64(streamResult.Id), 10)}, constant.COLON)
	streamResLine2 := strings.Join([]string{constant.EVENT, streamResult.Event}, constant.COLON)
	streamResLine3 := strings.Join([]string{constant.DATA, *streamResult.Data}, constant.COLON)
	streamRes := strings.Join([]string{streamResLine1, streamResLine2, streamResLine3}, "\n")
	streamRes = "data:" + streamRes
	// 预留换行
	streamRes += "\n\n"
	return streamRes
}
