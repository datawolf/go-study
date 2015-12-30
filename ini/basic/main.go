//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	// load the ini file
	cfg, err := ini.Load("gitconfig.ini")
	if err != nil {
		log.Println("ini Load error: ", err)
	}

	// Get the default section
	section, err := cfg.GetSection("")
	if err != nil {
		log.Println("ini GetSection error: ", err)
	}
	log.Println(section.GetKey("name")) // key, err := cfg.Section("").GetKey("key name")
	log.Println(section.Key("name"))    // key := cfg.Section("").Key("key name")

	// To check if a key exists
	log.Println(cfg.Section("").HasKey("email")) // yes := cfg.Section("").HasKey("key name")

	// to create a  new key
	_, err = cfg.Section("").NewKey("key", "value")

	// To get a list of keys or key names
	log.Println(cfg.Section("").Keys())
	log.Println(cfg.Section("").KeyStrings())

	// To get a clone hash of keys and corresponding values
	section, err = cfg.GetSection("")
	log.Println(section.KeysHash())

	// To get a strign value
	log.Println(cfg.Section("user").Key("name").String())

	// To validate key value on the fly
	val := cfg.Section("user").Key("not_exist").Validate(func(in string) string {
		if len(in) == 0 {
			return "default"
		}
		return in
	})
	log.Println(val)
}
