package main

import (
	"lesson2/db"
	"lesson2/mylib"
	"lesson2/nwlib"
	"lesson2/thirdpartylib"
)

func main() {
	mylib.Time()
	mylib.Regex()
	mylib.Sort()
	mylib.Iota()
	mylib.Context()
	mylib.Ioutil()

	nwlib.HTTP()
	nwlib.UnmarshalMarshal()
	nwlib.Hmac()

	thirdpartylib.Semaphore()
	thirdpartylib.Ini()
	thirdpartylib.Talib()

	db.Dbope()
}
