package mylib

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func Ioutil() {
	StartLog()
	defer EndLog()

	// ファイル読み込み
	// content, err := ioutil.ReadFile("main.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatal(err)
	// }

	// バッファのデータ読み込み
	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
