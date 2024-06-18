package model

import (
	"io"
	"time"
)

type CommonOssFile struct {
	Reader  io.ReadSeeker
	ModTime time.Time
}
