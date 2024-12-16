package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"server/internal/constants"
)

type Logger struct {
	logrus *logrus.Entry
}

type Fields map[string]interface{}

func New() *Logger {
	return &Logger{
		logrus: logrus.NewEntry(logrus.StandardLogger()),
	}
}

func MustInitGlobal(env constants.Environment) {
	switch env {
	case constants.ENVLocal:
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:      true,
			DisableTimestamp: true,
		})
	case constants.ENVDev:
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case constants.ENVProd:
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		panic("Unsupported environment: " + env)
	}
}

// Show logs

func (logger *Logger) Trace(args ...any) {
	logger.logrus.Trace(args)
}

func (logger *Logger) Debug(args ...any) {
	logger.logrus.Debug(args)
}

func (logger *Logger) Print(args ...any) {
	logger.logrus.Print(args)
}

func (logger *Logger) Info(args ...any) {
	logger.logrus.Info(args)
}

func (logger *Logger) Warn(args ...any) {
	logger.logrus.Warn(args)
}

func (logger *Logger) Warning(args ...any) {
	logger.logrus.Warning(args)
}

func (logger *Logger) Error(args ...any) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		logger.WithFields(Fields{
			"File": file,
			"Line": line,
		})
	}

	logger.logrus.Error(args)
}

func (logger *Logger) Fatal(args ...any) {
	logger.logrus.Fatal(args)
}

func (logger *Logger) Panic(args ...any) {
	logger.logrus.Panic(args)
}
