//
// cap.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	slice := make([]byte, 0, 5)
	log.Printf("slice: %d", cap(slice))

	channel := make(chan int, 10)
	log.Printf("channel: %d", cap(channel))

	var pointer *[15]byte
	log.Printf("pointer: %d == %d", cap(pointer), len(pointer))
}

/*
2015/12/16 19:15:01 slice: 5
2015/12/16 19:15:01 channel: 10
2015/12/16 19:15:01 pointer: 15 == 15
*/
