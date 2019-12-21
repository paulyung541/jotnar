package jotnar

import (
	"errors"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Logger

func InitLogrus() {
	defaultLogger = logrus.New()
	switch strings.ToLower(defualtLogConfig.Level) {
	case "panic":
		defaultLogger.Level = logrus.PanicLevel
	case "fatal":
		defaultLogger.Level = logrus.FatalLevel
	case "error":
		defaultLogger.Level = logrus.ErrorLevel
	case "Warn":
		defaultLogger.Level = logrus.WarnLevel
	case "info":
		defaultLogger.Level = logrus.InfoLevel
	case "debug":
		defaultLogger.Level = logrus.DebugLevel
	case "trace":
		defaultLogger.Level = logrus.TraceLevel
	}

	if defualtLogConfig.Format == "json" {
		jsonFormat := &logrus.JSONFormatter{
			CallerPrettyfier: myCallerPrettyfier,
		}

		jsonFormat.TimestampFormat = "2006-01-02 15:04:05"
		if defualtLogConfig.Timeformat != "" {
			jsonFormat.TimestampFormat = defualtLogConfig.Timeformat
		}

		defaultLogger.SetFormatter(jsonFormat)
	} else if defualtLogConfig.Format == "text" {
		textFormat := defaultLogger.Formatter.(*logrus.TextFormatter)
		textFormat.CallerPrettyfier = myCallerPrettyfier
		textFormat.TimestampFormat = "2006-01-02 15:04:05"
		if defualtLogConfig.Timeformat != "" {
			textFormat.TimestampFormat = defualtLogConfig.Timeformat
		}
		textFormat.FullTimestamp = true
	} else {
		errExit(errors.New("format value must be json or text"))
	}
}

func myCallerPrettyfier(f *runtime.Frame) (string, string) {
	s := strings.Split(f.Function, ".")
	funcname := s[len(s)-1]
	dir, filename := path.Split(f.File)
	tmpArray := strings.Split(dir, string(os.PathSeparator))
	if len(tmpArray) <= 3 {
		filename = f.File
	} else {
		tmpArray = tmpArray[len(tmpArray)-3:]
		filename = strings.Join(tmpArray, string(os.PathSeparator)) + filename
	}
	return funcname, filename
}

func GetLogger() *logrus.Logger {
	return defaultLogger
}
