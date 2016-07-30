//
// reflect2.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is uint8:", v.Kind() == reflect.Uint8)
	fmt.Println("v.Uint() type:", reflect.ValueOf(v.Uint()).Type())
	x = uint8(v.Uint())
}
