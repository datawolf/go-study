//
// Get.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//
// ./Get http://releases.rancher.com/os/releases.yml

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"io/ioutil"
)

func main() {
	//os.Setenv("HTTPS_PROXY", "http://username:password@proxyhost:port")
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	reader := response.Body
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println(string(body))
	os.Exit(0)
}
