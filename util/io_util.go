package util

import (
	"context"
	"github.com/easyship/util/log"
	"io"
)

func CloseReader(closer io.ReadCloser) {
	if closer != nil {
		defer func(respReader io.ReadCloser) {
			err := respReader.Close()
			if err != nil {
				log.ErrorWithContext(context.Background(), "close io resp error:%v", err)
			}
		}(closer)
	}
}
