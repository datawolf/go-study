//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	// To set a key/value pair, use os.Setenv.
	// To get a vaule for a key, use os.Getenv. This will return an empty string if the key is not present in the environment
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// use os.Environ to list all key/value pairs in the environment.
	// this returns a slice of strings in the form KEY=value.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(e)
		fmt.Println("        ", pair[0])
		fmt.Println("        ", pair[1])
	}
}
