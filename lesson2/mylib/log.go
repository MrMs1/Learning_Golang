package mylib

import (
	"fmt"
	"runtime"
)

func StartLog() {
	funcName := GetFuncName()
	fmt.Printf("*********** %v START ***********\n", funcName)
}

func EndLog() {
	funcName := GetFuncName()
	fmt.Printf("*********** %v END ***********\n\n", funcName)
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
