//
// exmaple2.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	i = 0
)

func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "i = %d", i)
}

func main() {
	http.Handle("/args", http.HandlerFunc(ArgServer))
	go update()
	http.ListenAndServe(":8080", nil)
}

func update() {
	for {
		i++
		time.Sleep(5 * time.Second)
		fmt.Println("update i")
	}
}
