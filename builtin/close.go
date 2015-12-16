//
// close.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	c := make(chan int, 1)
	c <- 1

	log.Println(<-c) // print 1

	c <- 2
	close(c)

	log.Println(<-c) // print 2
	log.Println(<-c) // print 0

	if i, ok := <-c; ok {
		log.Printf("Channel is open, got %d", i)
	} else {
		log.Printf("Channel is close, got %d", i)
	}

	close(c) // Panics, channel is already closed
}

/* result
2015/12/16 17:55:28 1
2015/12/16 17:55:28 2
2015/12/16 17:55:28 0
2015/12/16 17:55:28 Channel is close, got 0
panic: close of closed channel

goroutine 1 [running]:
main.main()
	/go/src/github.com/datawolf/go-study/builtin/close.go:30 +0x496
*/
