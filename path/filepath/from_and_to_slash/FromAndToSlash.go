//
// FromAndToSlash.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"path/filepath"
)

const sep = filepath.Separator

type PathTest struct {
	path, result string
}

var slashtests = []PathTest{
	{"", ""},
	{"/", string(sep)},
	{"/a/b", string([]byte{sep, 'a', sep, 'b'})},
	{"a//b", string([]byte{'a', sep, sep, 'b'})},
}

func main() {
	for _, test := range slashtests {
		if s := filepath.FromSlash(test.path); s != test.result {
			fmt.Printf("FromSlash(%q) = %q, want %q\n", test.path, s, test.result)
		}
		if s := filepath.ToSlash(test.result); s != test.path {
			fmt.Printf("ToSlash(%q) = %q, want %q\n", test.result, s, test.path)
		}
	}
}
