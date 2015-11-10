//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"os"
	"fmt"
)

// Note that unlike e.g. C, Go does not use an integer return
// value from main to indicate exit status. If you would like
// to exit with a non-zero status you should use os.Exit

func main() {
	// defer will not run when using os.Exit, so this fmt.Println will never 
	// be called.
	defer	fmt.Println("!")

	// Exit with status 3
	os.Exit(3)
}
