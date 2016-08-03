//
// regexp1.go
// Copyright (C) 2016 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"regexp"
	"fmt"
)

var (
	re				= regexp.MustCompile(`a(x*)b(y|z)c`)
	lineError		= regexp.MustCompile(`^YAML error line (?P<line>[[:digit:]]+): (?P<msg>.*)$`)
	line1Error		= regexp.MustCompile(`YAML error line (?P<line>[[:digit:]]+): (?P<msg>.*)`)
)

func main() {
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	matches := lineError.FindStringSubmatch("YAML error line 12: can not find the symbol")

	for i:= 0; i < len(matches); i++ {
		fmt.Printf("matches[%v] = %v\n", i, matches[i])
	}

	matches1 := line1Error.FindStringSubmatch("WangLongYAML error line 12: can not find the symbol")

	for i:= 0; i < len(matches1); i++ {
		fmt.Printf("matches1[%v] = %v\n", i, matches1[i])
	}
}
