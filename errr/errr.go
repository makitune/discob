package errr

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func Printf(format string, argv ...interface{}) {
	pt, file, line, ok := runtime.Caller(1)
	if ok {
		funcName := runtime.FuncForPC(pt).Name()
		log.Printf("%s:%d: func=%v ", file, line, funcName)
	}

	log.Printf(format, argv...)
}

func Fatalf(format string, argv ...interface{}) {
	Printf(format, argv...)
	os.Exit(1)
}

func Usage() {
	err := fmt.Errorf("ERRRRRRR!!!!!!")
	Printf("%s\n", err)
}
