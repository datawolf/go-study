//
// main.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"strconv"
	"fmt"
)

func main() {
	card  := "AC80ECD1"
	if s, err := strconv.ParseUint(card, 16,64); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}
