package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{
		Logger: logrus.New(),
	}
}

//获取打印级别
func (l *LogrusLogger) GetLevel() LogLevel {
	return LogLevel(l.Logger.GetLevel())
}

//设置打印级别
func (l *LogrusLogger) SetLevel(level string) error {
	logLevel, err := ParseLevel(level)
	if err != nil {
		return err
	}
	l.Logger.SetLevel(logrus.Level(logLevel))
	return nil
}

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *LogrusLogger) Printf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.Info(args...)
}

func (l *LogrusLogger) Print(args ...interface{}) {
	l.Print(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.Error(args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.Fatal(args...)
}

func (l *LogrusLogger) Panic(args ...interface{}) {
	l.Panic(args...)
}

func (l *LogrusLogger) Debugln(args ...interface{}) {
	l.Debugln(args...)
}

func (l *LogrusLogger) Infoln(args ...interface{}) {
	l.Infoln(args...)
}

func (l *LogrusLogger) Println(args ...interface{}) {
	l.Println(args...)
}

func (l *LogrusLogger) Warnln(args ...interface{}) {
	l.Warnln(args...)
}

func (l *LogrusLogger) Errorln(args ...interface{}) {
	l.Errorln(args...)
}

func (l *LogrusLogger) Fatalln(args ...interface{}) {
	l.Fatalln(args...)
}

func (l *LogrusLogger) Panicln(args ...interface{}) {
	l.Panicln(args...)
}
