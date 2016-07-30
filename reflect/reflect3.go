//
// reflect3.go
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
	type MyInt int
	var x MyInt
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)

	fmt.Println("type: ", v.Type())
	fmt.Println("v.Kind: ", v.Kind())
	fmt.Println("t.Kind: ", t.Kind())
}
