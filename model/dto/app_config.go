package model

type AppConfigDto struct {
	Url      string   //网站链接
	AppName  string   //App名称
	AppLogo  string   //app Logo
	AppText  string   //app简介
	AppPic   string   //app介绍图片
	AppLabel []string // app标签列表
}
