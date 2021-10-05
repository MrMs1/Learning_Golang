package mylib

import (
	"fmt"
	"runtime"
	"strings"
)

func StartLog() {
	funcName := GetFuncName()
	fmt.Printf("%s %v START %s\n", strings.Repeat("*", 10), funcName, strings.Repeat("*", 10))
}

func EndLog() {
	funcName := GetFuncName()
	fmt.Printf("%s %v END %s\n\n", strings.Repeat("*", 10), funcName, strings.Repeat("*", 10))
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
