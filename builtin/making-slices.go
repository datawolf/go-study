//
// making-slices.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main() {
	a := []byte("hello world!\n")
	printSlice("a", a)
	b := a[:2]
	printSlice("b", b)
	c := a[4:]
	printSlice("c", c)
	d := a[3:7]
	printSlice("d", d)

}

func printSlice(name string, data []byte) {
	fmt.Printf("%s len=%d cap=%d %v\n", name, len(data), cap(data), string(data))
}
