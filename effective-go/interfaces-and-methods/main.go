//
// main.go
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

// Simple counter server.
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	go update(ctr)
	http.ListenAndServe(":8080", nil)
}

func update(c *Counter) {
	for {
		time.Sleep(5 * time.Second)
		c.n++
		fmt.Println("updated")
	}
}
