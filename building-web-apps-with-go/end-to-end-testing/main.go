//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func HelloWorld(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(rw, "Hello World")
}

func App() http.Handler {
	n := negroni.Classic()

	m := func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		fmt.Fprint(rw, "Before...")
		next(rw, req)
		fmt.Fprint(rw, "...After")
	}
	n.Use(negroni.HandlerFunc(m))

	r := httprouter.New()

	r.GET("/", HelloWorld)
	n.UseHandler(r)

	return n
}

func main() {
	http.ListenAndServe(":8080", App())
}
