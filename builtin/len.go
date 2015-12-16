//
// len.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	slice := make([]byte, 10)
	log.Printf("slice: %d", len(slice))

	str := "hello world\n"
	log.Printf("string: %d", len(str))

	m := make(map[string]int)
	m["hello"] = 1
	log.Printf("map: %d", len(m))

	channel := make(chan int, 5)
	log.Printf("channel: %d", len(channel))
	channel <- 1
	log.Printf("channel: %d", len(channel))

	var pointer *[5]byte
	log.Printf("pointer: %d", len(pointer))

}

/*
2015/12/16 19:10:05 slice: 10
2015/12/16 19:10:05 string: 12
2015/12/16 19:10:05 map: 1
2015/12/16 19:10:05 channel: 0
2015/12/16 19:10:05 channel: 1
2015/12/16 19:10:05 pointer: 5
*/
