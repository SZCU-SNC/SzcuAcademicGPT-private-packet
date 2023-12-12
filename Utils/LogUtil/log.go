package Utils

import (
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	maxLogSize    = 10 // MB
	maxLogBackups = 3
	maxLogAge     = 28 // days
)

type MyLogger struct {
	*logrus.Logger
}

type ErrorHook struct {
	ErrorLogger *logrus.Logger
}

func NewLogger(serviceName string) *MyLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logFile := &lumberjack.Logger{
		Filename:   "./" + serviceName + "_" + time.Now().Format("06-01-02") + ".log",
		MaxSize:    maxLogSize,
		MaxBackups: maxLogBackups,
		MaxAge:     maxLogAge,
		Compress:   true,
	}
	logger.SetOutput(logFile)

	errorLogger := logrus.New()
	errorLogger.SetFormatter(&logrus.JSONFormatter{})
	errorLogger.SetOutput(&lumberjack.Logger{
		Filename:   "./" + serviceName + "_" + time.Now().Format("06-01-02") + "_err.log",
		MaxSize:    maxLogSize,
		MaxBackups: maxLogBackups,
		MaxAge:     maxLogAge,
		Compress:   true,
	})

	logger.AddHook(&ErrorHook{ErrorLogger: errorLogger})

	return &MyLogger{Logger: logger}
}

func (hook *ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *ErrorHook) Fire(entry *logrus.Entry) error {
	hook.ErrorLogger.WithFields(entry.Data).Error(entry.Message)
	return nil
}

func getCaller() string {
	pc, file, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	return path.Base(file) + "/" + fn.Name()
}

func (logger *MyLogger) WithCaller() *logrus.Entry {
	return logger.WithField("caller", getCaller())
}

// func main() {
// 	serviceName := "security"
// 	logger := NewLogger(serviceName)

// 	logger.WithCaller().Info("这是一条普通信息")
// 	logger.WithCaller().Error("这是一条错误信息")
// }
