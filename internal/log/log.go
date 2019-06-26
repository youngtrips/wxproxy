package log

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	_log *zap.SugaredLogger
	once sync.Once
)

const (
	MAXN_SIZE    = 128 // MB
	MAXN_BACKUPS = 30  //
	MAXN_AGE     = 31  // a month
)

func MyTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05"))
}

func Init(app string) {
	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/%s.log", app),
		MaxSize:    MAXN_SIZE,
		MaxBackups: MAXN_BACKUPS,
		MaxAge:     MAXN_AGE,
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     MyTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)

	caller := zap.AddCaller()
	zlog := zap.New(core, caller, zap.AddCallerSkip(1))
	_log = zlog.Sugar()
}

func Sync() {
	if _log != nil {
		_log.Sync()
	}
}

type ElasticLogger struct {
}

func (el *ElasticLogger) Printf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func GetLogger() *zap.SugaredLogger {
	if _log == nil {
		panic("OOPS....Has not initialize logger yet.")
	}
	return _log
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}
