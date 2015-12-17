//
// panic_recover.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"errors"
	"log"
)

func handlePanic(f func()) {
	defer func() {
		if r := recover(); r != nil {
			if str, ok := r.(string); ok {
				log.Printf("got a string error: %s", str)
				return
			}

			if err, ok := r.(error); ok {
				log.Printf("got an error error: %s", err.Error())
				return
			}
			log.Printf("got a different kind of error: %v", r)
		}
	}()

	f()
}

func main() {
	handlePanic(func() {
		panic("string error")
	})
	handlePanic(func() {
		panic(errors.New("error error"))
	})
	handlePanic(func() {
		panic(10)
	})
}

/*
2015/12/17 19:16:28 got a string error: string error
2015/12/17 19:16:28 got an error error: error error
2015/12/17 19:16:28 got a different kind of error: 10
*/
