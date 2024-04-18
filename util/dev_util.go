package util

import "os"

type EnvName string

var (
	DEV  EnvName = "dev"
	PROD EnvName = "prod"
)

var env EnvName = DEV

func InitEnv() {
	if os.Getenv("env") == string(PROD) {
		env = PROD
	}
	env = DEV

}

// 是否为测试环境
func IsDev() bool {
	return env == DEV
}

func IsProd() bool {
	return env == PROD
}
