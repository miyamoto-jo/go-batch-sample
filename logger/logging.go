package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry
var logFile *os.File

func SetupLog(useFile bool) {
	lg := logrus.New()
	mode := int32(0777)
	os.Mkdir(`.`+string(filepath.Separator)+`log`, os.FileMode(mode))
	file, err := os.OpenFile(`./log/application.log`, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
	f := new(logrus.JSONFormatter)
	f.TimestampFormat = "2006-01-02T15:04:05.999Z07:00"
	lg.Formatter = f

	hostname, _ := os.Hostname()
	logger = lg.WithField("host", hostname)

	if useFile {
		writer := io.MultiWriter(os.Stdout, file)
		lg.SetOutput(writer)
	} else {
		lg.SetOutput(os.Stdout)
	}
	logger.Info("Logrus is Setup for logging.")
}

func TerminateLog() {
	logFile.Close()
}

func Println(args ...interface{}) {
	if logger == nil {
		log.Println(args...)
		return
	}
	logger.Println(args...)
}

func Printf(format string, args ...interface{}) {
	if logger == nil {
		log.Printf(format, args...)
		return
	}
	logger.Printf(format, args...)
}

func Infoln(args ...interface{}) {
	if logger == nil {
		log.Println(args...)
		return
	}
	logger.Infoln(args...)
}

func Error(args ...interface{}) {
	if logger == nil {
		log.Println(args...)
		return
	}
	logger.Error(args...)
}

func Errorln(args ...interface{}) {
	if logger == nil {
		log.Println(args...)
		return
	}
	logger.Errorln(args...)
}

// スタックトレースを出力
func StackTrace() {
	for depth := 0; ; depth++ {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		logger.Printf("%d: %v:%d", depth, file, line)
	}
}
