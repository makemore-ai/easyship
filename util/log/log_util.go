package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

func InfoWithContext(ctx context.Context, format string, args ...interface{}) {
	logrus.WithContext(ctx).Infof(format, args...)
}

func ErrorWithContext(ctx context.Context, format string, args ...interface{}) {
	logrus.WithContext(ctx).Errorf(format, args...)
}

func WarnWithContext(ctx context.Context, format string, args ...interface{}) {
	logrus.WithContext(ctx).Warnf(format, args...)
}
