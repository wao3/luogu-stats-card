package common

import (
	"fmt"
	"log"
	"runtime"
)

func LogInfo(info string, args ...interface{}) {
	logs("INFO", info, args)
}

func logDebug(info string, args ...interface{}) {
	logs("DEBUG", info, args)
}

func LogWarn(info string, args ...interface{}) {
	logs("WARN", info, args)
}

func LogError(info string, args ...interface{}) {
	logs("ERROR", info, args)
}

func logs(level string, info string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	log.Printf("[%s] file=%s; line=%d; info=%s", level, file, line, fmt.Sprintf(info, args))
}
