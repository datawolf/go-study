//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"time"

	"gopkg.in/ini.v1"
)

type Embeded struct {
	Dates  []time.Time `delim:"|"`
	Places []string
	None   []int
}

type Author struct {
	Name      string `ini:"NAME"`
	Male      bool
	Age       int
	GPA       float64
	NeverMind string `ini:"-"`
	*Embeded
}

func main() {
	a := &Author{"Unknwon", true, 21, 2.8, "",
		&Embeded{
			[]time.Time{time.Now(), time.Now()},
			[]string{"Shanghai", "Boston"},
			[]int{},
		}}

	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, a)
	if err != nil {
		print("ini reflectFrom error:", err)
	}

	cfg.SaveTo("author.ini")
}
