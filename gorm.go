package klog

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

var _ logger.Interface = &GORMLogger{}

// NewGORMLogger Create GORMLogger instance
func NewGORMLogger(l Interface) logger.Interface {
	return &GORMLogger{l: l}
}

// GORMLogger A structure that implements gorm logger.Interface
type GORMLogger struct {
	l Interface
}

// LogMode You cannot change the log level by calling LogMode
// You can only change the display level of the log through logger.SetLevel()
func (g *GORMLogger) LogMode(level logger.LogLevel) logger.Interface {
	// You cannot change the log level by calling LogMode
	switch level {
	case logger.Info:
		return &GORMLogger{l: g.l.WithLevel(InfoLevel)}
	case logger.Warn:
		return &GORMLogger{l: g.l.WithLevel(WarnLevel)}
	case logger.Error:
		return &GORMLogger{l: g.l.WithLevel(ErrorLevel)}
	case logger.Silent:
		return &GORMLogger{l: g.l.WithLevel(OffLevel)}
	}
	return g
}

// Info Print information level log
func (g *GORMLogger) Info(ctx context.Context, s string, i ...interface{}) {
	_ = g.l.Output(InfoLevel, 4, g.l.Fields(), fmt.Sprintf(s, i...))
}

// Warn Print Warn log
func (g *GORMLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	_ = g.l.Output(WarnLevel, 4, g.l.Fields(), fmt.Sprintf(s, i...))
}

// Error Print error log
func (g *GORMLogger) Error(ctx context.Context, s string, i ...interface{}) {
	_ = g.l.Output(ErrorLevel, 4, g.l.Fields(), fmt.Sprintf(s, i...))
}

// Trace Print gorm sql detailed log
func (g *GORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l := g.l.WithDurationField("duration", time.Since(begin))
	if err != nil {
		l = l.WithError(err)
	}
	sql, rows := fc()
	l = l.WithIntField("rows", int(rows))
	_ = l.Output(TraceLevel, 4, l.Fields(), sql)
}
