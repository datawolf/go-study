//
// make.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	unbuffered := make(chan int)
	log.Printf("unbuffered: %v, type %T, len: %d, cap: %d", unbuffered, unbuffered, len(unbuffered), cap(unbuffered))

	buffered := make(chan int, 10)
	log.Printf("buffered: %v, type %T, len: %d, cap: %d", buffered, buffered, len(buffered), cap(buffered))

	m := make(map[string]int)
	log.Printf("m: %v, len: %d", m, len(m))

	slice := make([]byte, 5)
	log.Printf("slice: %v, len: %d, cap: %d", slice, len(slice), cap(slice))

	slice2 := make([]byte, 0, 10)
	log.Printf("slice2: %v, len: %d, cap: %d", slice2, len(slice2), cap(slice2))
}

/* the result:
2015/12/16 17:24:37 unbuffered: 0xc82001c0c0, type chan int, len: 0, cap: 0
2015/12/16 17:24:37 buffered: 0xc82007e000, type chan int, len: 0, cap: 10
2015/12/16 17:24:37 m: map[], len: 0
2015/12/16 17:24:37 slice: [0 0 0 0 0], len: 5, cap: 5
2015/12/16 17:24:37 slice2: [], len: 0, cap: 10
*/
