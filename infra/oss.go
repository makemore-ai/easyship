package infra

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/easyship/util/log"
)

var provider = credentials.NewEnvironmentVariableCredentialsProvider()
var cfg = oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithEndpoint("oss-cn-beijing.aliyuncs.com").WithRegion("cn-beijing")

func GetOssFile(ctx context.Context, bucketName string, filePath string) (*oss.ReadOnlyFile, error) {
	client := oss.NewClient(cfg)
	// 直接读取文件
	file, err := client.OpenFile(ctx, bucketName, filePath)
	if err != nil {
		log.ErrorWithContext(ctx, "failed to open file %v", err)
		return nil, err
	}
	return file, nil
}
