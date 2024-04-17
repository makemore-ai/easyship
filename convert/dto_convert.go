package convert

import (
	"github.com/easyship/model"
	dto_model "github.com/easyship/model/dto"
)

func ConvertAppConfig2Dto(appConfig *model.AppConfig) *dto_model.AppConfigDto {
	if appConfig == nil {
		return nil
	}
	return &dto_model.AppConfigDto{
		Url:      appConfig.Url(),
		AppName:  appConfig.AppName(),
		AppLogo:  appConfig.AppLogo(),
		AppText:  appConfig.AppText(),
		AppPic:   appConfig.AppPic(),
		AppLabel: appConfig.AppLabel(),
	}
}

func ConvertAppConfig2DtoList(appConfigList []*model.AppConfig) []*dto_model.AppConfigDto {
	if appConfigList == nil || len(appConfigList) == 0 {
		return []*dto_model.AppConfigDto{}
	}
	appConfigDtoList := make([]*dto_model.AppConfigDto, 0, len(appConfigList))
	for _, appConfig := range appConfigList {
		appConfigDtoList = append(appConfigDtoList, ConvertAppConfig2Dto(appConfig))
	}
	return appConfigDtoList
}
