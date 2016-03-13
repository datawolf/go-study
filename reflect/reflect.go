//
// reflect.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
}
