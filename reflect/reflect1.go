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
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("Kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
