//
// reflect4.go
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
/*
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())
*/
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println("v :", v.Interface())
	fmt.Println("x :", x)
}
