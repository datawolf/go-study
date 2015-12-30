//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

func main() {
	cfg, err := ini.Load("person.ini")
	if err != nil {
		log.Println("ini Load error :", err)
	}

	p := new(Person)
	err = cfg.MapTo(p)

	log.Println("------------person info----------------")
	log.Println(p.Name)
	log.Println(p.Age)
	log.Println(p.Male)
	log.Println(p.Born)
	log.Println(p.Created)
	log.Println(p.Content)
	log.Println(p.Cities)

}
