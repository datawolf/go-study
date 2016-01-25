//
// indent.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"encoding/json"
	"log"
	"os"
	"bytes"
)

type Road struct {
	Name   string
	Number int
}

func main() {
	roads := []Road{
		{"五道口", 28},
		{"西二旗", 25},
	}
	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)
}
