//
// reflect5.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, typeOfT.Field(i).Type, f.Interface())
	}

	fmt.Println("")

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")

	fmt.Println("t is now: ", t)
}
