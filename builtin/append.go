//
// append.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	// empty slice, with capacity of 10
	ints := make([]int, 0, 10)
	log.Printf("ints: %v", ints)

	ints2 := append(ints, 1, 2, 3)

	log.Printf("ints2: %v", ints2)
	log.Printf("slice was at %p, it's probably still at %p", ints, ints2)

	moreInts := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	ints3 := append(ints2, moreInts...)

	log.Printf("ints3: %v", ints3)
	log.Printf("Slice was at %p, and it moved to %p", ints2, ints3)

	ints4 := []int{1, 2, 3}
	log.Printf("ints4: %v", ints4)

	// The idiomatic may to append to a slice,
	// just assign to the same variable again
	ints4 = append(ints4, 4, 5, 6)
	log.Printf("ints4: %v", ints4)
}

/* result

2015/12/16 18:56:25 ints: []
2015/12/16 18:56:25 ints2: [1 2 3]
2015/12/16 18:56:25 slice was at 0xc82007a0f0, it's probably still at 0xc82007a0f0
2015/12/16 18:56:25 ints3: [1 2 3 4 5 6 7 8 9 10 11 12 13 14]
2015/12/16 18:56:25 Slice was at 0xc82007a0f0, and it moved to 0xc82007e000
2015/12/16 18:56:25 ints4: [1 2 3]
2015/12/16 18:56:25 ints4: [1 2 3 4 5 6]

*/
