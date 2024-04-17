package model

type AppConfig struct {
	url      string   //网站链接
	appName  string   //App名称
	appLogo  string   //app Logo
	appText  string   //app简介
	appPic   string   //app介绍图片
	appLabel []string // app标签列表

}

func (a AppConfig) Url() string {
	return a.url
}

func (a AppConfig) AppName() string {
	return a.appName
}

func (a AppConfig) AppLogo() string {
	return a.appLogo
}

func (a AppConfig) AppText() string {
	return a.appText
}

func (a AppConfig) AppPic() string {
	return a.appPic
}

func (a AppConfig) AppLabel() []string {
	return a.appLabel
}

func NewAppConfig(url string, appName string, appLogo string, appText string, appPic string, appLabel []string) *AppConfig {
	return &AppConfig{url: url, appName: appName, appLogo: appLogo, appText: appText, appPic: appPic, appLabel: appLabel}
}
