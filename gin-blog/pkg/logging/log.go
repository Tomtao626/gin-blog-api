package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      string = ""
	DefaultCallerDepth int    = 2
	logger             *log.Logger
	logPrefix          string = ""
	levelFlags                = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	SetPrefix(DEBUG)
	log.Println("DEBUG")
}

func Info(v ...interface{}) {
	SetPrefix(INFO)
	log.Println("INFO")
}

func Warn(v ...interface{}) {
	SetPrefix(WARNING)
	log.Println("WARNING")
}

func Error(v ...interface{}) {
	SetPrefix(ERROR)
	log.Println("ERROR")
}

func Fatal(v ...interface{}) {
	SetPrefix(FATAL)
	log.Println("FATAL")
}

func SetPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
