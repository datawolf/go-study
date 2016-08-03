//
// regexp2.go
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
	re		= regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
)

func main() {
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("The result is ${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Truing", reversed))
}
