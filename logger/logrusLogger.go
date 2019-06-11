package logger

import "github.com/sirupsen/logrus"

var (
	//缺省的logrus logger
	DefaultLogrusLogger = &LogrusLogger{}
)

type LogrusLogger struct{}

//获取打印级别
func (l *LogrusLogger) GetLevel() LogLevel {
	return LogLevel(logrus.GetLevel())
}

//设置打印级别
func (l *LogrusLogger) SetLevel(level string) error {
	logLevel, err := ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(logrus.Level(logLevel))
	return nil
}

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (l *LogrusLogger) Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	logrus.Info(args...)
}

func (l *LogrusLogger) Print(args ...interface{}) {
	logrus.Print(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	logrus.Error(args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (l *LogrusLogger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (l *LogrusLogger) Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

func (l *LogrusLogger) Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

func (l *LogrusLogger) Println(args ...interface{}) {
	logrus.Println(args...)
}

func (l *LogrusLogger) Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

func (l *LogrusLogger) Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

func (l *LogrusLogger) Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func (l *LogrusLogger) Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
