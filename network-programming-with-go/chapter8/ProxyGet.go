//
// ProxyGet.go
// Copyright (C) 2015 wanglong <wanglong@wanglong-Lenovo-Product>
//
// Distributed under terms of the MIT license.
//
// ./ProxyGet http://proxy-host:port    http://releases.rancher.com/os/releases.yml
package main

import (
	"os"
	"net/http"
	"net/http/httputil"
	"fmt"
	"io"
	"net/url"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
		os.Exit(1)
	}

	proxyString := os.Args[1]
	proxyURL, err := url.Parse(proxyString)
	checkError(err)

	rawURL := os.Args[2]
	url, err := url.Parse(rawURL)
	checkError(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)

	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	response, err := client.Do(request)
	checkError(err)
	fmt.Println("Read ok")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("Response OK")

	reader := response.Body
    body, err := ioutil.ReadAll(reader)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }
    fmt.Println(string(body))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

