package nwlib

import (
	"fmt"
	"io/ioutil"
	"lesson2/mylib"
	"net/http"
	"net/url"
)

func HTTP() {
	mylib.StartLog()
	defer mylib.EndLog()

	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// boby, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(boby))

	// URLの検証
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	// req, _ := http.NewRequest("GET", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("If-None-Match", `W/"wyxxy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	boby, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(boby))
}
