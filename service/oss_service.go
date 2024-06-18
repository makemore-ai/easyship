package service

import (
	"context"
	"github.com/easyship/infra"
	"github.com/easyship/infra/constant"
	"github.com/easyship/util/log"
	"io/fs"

	"strings"
)

var DefaultOssFileSystem = &ossFileSystem{}

type ossFileSystem struct {
}

func (*ossFileSystem) Open(name string) (fs.File, error) {
	return GetOssDirectFile(context.Background(), name)
}

// GetOssDirectFile PathName = bucketName + '/' + filePath
func GetOssDirectFile(ctx context.Context, pathName string) (fs.File, error) {
	bucketName, filePath := transferPath(ctx, pathName)
	log.ErrorWithContext(ctx, "bucketName empty:%v, %v,%v", pathName, bucketName, filePath)
	if len(bucketName) == 0 || len(filePath) == 0 {
		log.ErrorWithContext(ctx, "bucketName empty:%v,%v", bucketName, filePath)
		return nil, infra.NewDefaultSystemError()
	}
	file, err := infra.GetOssFile(ctx, bucketName, filePath)
	if err != nil {
		log.ErrorWithContext(ctx, "GetOssFile error:%v,%v,%v", bucketName, pathName, err)
		return nil, err
	}
	return file, nil
}

func transferPath(ctx context.Context, pathName string) (bucketName string, filePath string) {
	pathSplitList := strings.SplitN(pathName, "/", 2)
	if len(pathSplitList) != 2 {
		log.WarnWithContext(ctx, "transferPath error:%v", pathName)
		bucketName = constant.EMPTY_STRING
		filePath = constant.EMPTY_STRING
		return
	}
	bucketName = pathSplitList[0]
	filePath = pathSplitList[1]
	return
}
