package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/loadept/loadept.com/internal/config"
)

var (
	INFO       *log.Logger
	ERROR      *log.Logger
	onceLogger sync.Once
)

var (
	infoLoggerFile  *os.File
	errorLoggerFile *os.File
)

func NewLogger() {
	onceLogger.Do(func() {
		var err error

		currentDate := time.Now().Format("2006-01-02")
		logInfoFile := fmt.Sprintf("logs/access-%s.log", currentDate)
		logErrorFile := fmt.Sprintf("logs/error-%s.log", currentDate)

		infoLoggerFile, err = os.OpenFile(logInfoFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		errorLoggerFile, err = os.OpenFile(logErrorFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		var ioInfoLogger, ioErrLogger io.Writer
		if config.Env.DEBUG == "true" {
			ioInfoLogger = os.Stdout
			ioErrLogger = os.Stderr
		} else {
			ioInfoLogger = infoLoggerFile
			ioErrLogger = errorLoggerFile
		}

		infoLogger := log.New(ioInfoLogger, "INFO: ", log.Ldate|log.Ltime)
		errorLogger := log.New(ioErrLogger, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

		INFO = infoLogger
		ERROR = errorLogger
	})
}

func CloseLogger() {
	if infoLoggerFile != nil {
		infoLoggerFile.Close()
	}
	if errorLoggerFile != nil {
		errorLoggerFile.Close()
	}
}
