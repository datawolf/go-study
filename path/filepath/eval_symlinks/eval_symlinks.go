//
// eval_symlinks.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	link := "link_file"
	dest, err := filepath.EvalSymlinks(link)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dest)

	link = "regular_file.txt"
	dest, err = filepath.EvalSymlinks(link)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dest)
}
