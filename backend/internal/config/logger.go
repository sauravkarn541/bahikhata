package config

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

var log = logrus.New()

func InitializeLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	log.SetReportCaller(true)
}

func GetLogger() *logrus.Logger {
	return log
}

// LogMode sets the log level for GORM
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs general info messages
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.logrusLogger.WithContext(ctx).Infof(msg, data...)
	}
}

// Warn logs warning messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.logrusLogger.WithContext(ctx).Warnf(msg, data...)
	}
}

// Error logs error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.logrusLogger.WithContext(ctx).Errorf(msg, data...)
	}
}

type GormLogger struct {
	logrusLogger *logrus.Logger
	LogLevel     logger.LogLevel
}

// Trace implements logger.Interface.
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error:
		sql, rows := fc()
		l.logrusLogger.WithContext(ctx).WithFields(logrus.Fields{
			"error": err,
			"sql":   sql,
			"rows":  rows,
			"took":  elapsed,
		}).Error("Trace error")

	case elapsed > 200*time.Millisecond && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.logrusLogger.WithContext(ctx).WithFields(logrus.Fields{
			"sql":  sql,
			"rows": rows,
			"took": elapsed,
		}).Warn("Trace slow query")

	case l.LogLevel == logger.Info:
		sql, rows := fc()
		l.logrusLogger.WithContext(ctx).WithFields(logrus.Fields{
			"sql":  sql,
			"rows": rows,
			"took": elapsed,
		}).Info("Trace")
	}
}

// New Gorm logger
func NewGormLogger(logLevel logger.LogLevel) *GormLogger {
	return &GormLogger{
		logrusLogger: log,
		LogLevel:     logLevel,
	}
}
