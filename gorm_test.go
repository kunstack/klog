package klog_test

import (
	"context"
	"testing"

	"github.com/kunstack/klog"
	"gorm.io/gorm/logger"
)

func TestGORM(t *testing.T) {
	l := klog.WithoutContext()
	defer l.Flush()
	g := klog.NewGORMLogger(l)
	g.LogMode(logger.Info).Info(context.Background(), "this is info message")
}
