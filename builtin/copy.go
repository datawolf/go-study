//
// copy.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	otherInts := []int{11, 12, 13, 14, 15, 16}

	log.Printf("ints: %v", ints)
	log.Printf("otherInts: %v", otherInts)

	copied := copy(ints[:3], otherInts)
	log.Printf("Copied %d ints from otherInts to ints", copied)
	log.Printf("ints: %v", ints)
	log.Printf("otherInts: %v", otherInts)

	hello := "Hello World!"
	bytes := make([]byte, len(hello))

	copy(bytes, hello)

	log.Printf("bytes: %v", bytes)
	log.Printf("hello: %v", hello)

}

/*
2015/12/16 19:01:55 ints: [1 2 3 4 5 6]
2015/12/16 19:01:55 otherInts: [11 12 13 14 15 16]
2015/12/16 19:01:55 Copied 3 ints from otherInts to ints
2015/12/16 19:01:55 ints: [11 12 13 4 5 6]
2015/12/16 19:01:55 otherInts: [11 12 13 14 15 16]
2015/12/16 19:01:55 bytes: [72 101 108 108 111 32 87 111 114 108 100 33]
2015/12/16 19:01:55 hello: Hello World!
*/
