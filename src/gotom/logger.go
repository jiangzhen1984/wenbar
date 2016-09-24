

package gotom


import (
    "log"
    "os"
    "fmt"
    "errors"
)

const (
    
    VDebug  = iota
    VInfo
    VWarn
    VError
)

var GoTomLogger * log.Logger = log.New(os.Stdout, "[GOTOM]", log.Ldate | log.Ltime | log.Lshortfile)
var VLogLevel = VDebug


func LD(format string, v ...interface{}) {
    LV(VDebug, format, v...)
}

func LI(format string, v ...interface{}) {
    LV(VInfo, format, v...)
}

func LW(format string, v  ...interface{}) {
    LV(VWarn, format, v...)
}

func LE(format string, v  ...interface{}) {
    LV(VError, format, v...)
}

func LV(level int, format string, v ...interface{}) {
   GoTomLogger.Output(3, fmt.Sprintf(format, v...)) 
}


func LF() {
   GoTomLogger.Output(2, "") 
}


func LP(format string, v ...interface{}) {
    if v == nil {
        return
    }
    GoTomLogger.Panicf(format, v...)
}


func ErrorMsg(format string, a ...interface{}) error {
    return errors.New(fmt.Sprintf(format, a...))
}

func E(format string, a ...interface{}) error {
    return ErrorMsg(format, a...)
}


func SetLogLevel(level int) {
    VLogLevel = level
}
