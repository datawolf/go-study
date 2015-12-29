//
// expvar.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"expvar"
	"flag"
	"log"
	"time"
)

var (
	times		= flag.Int("times", 1, "times to say hello")
	name		= flag.String("name", "World", "Thing to say hello to")
	helloTimes  = expvar.NewInt("hello")
)

func init() {
	expvar.Publish("time", expvar.Func(now))
}

func now() interface{} {
	return time.Now().Format(time.RFC3339Nano)
}

func hello(times int , name string) {
	helloTimes.Add(int64(times))
	for i:= 0; i < times; i++ {
		log.Printf("Hello, %s!", name)
	}
}

func printVars() {
	log.Println("expvars:")
	expvar.Do(func(kv expvar.KeyValue) {
		switch kv.Key {
		case "memstats":
			// Do nothing, this is a big output
		default:
			log.Printf("\t%s -> %s", kv.Key, kv.Value)
		}
	})
}

func main() {
	flag.Parse()
	printVars()
	hello(*times, *name)
	printVars()
	hello(*times, *name)
	printVars()
}
/*
2015/12/17 19:36:04 expvars:
2015/12/17 19:36:04 	cmdline -> ["./expvar"]
2015/12/17 19:36:04 	hello -> 0
2015/12/17 19:36:04 	time -> "2015-12-17T19:36:04.644527219+08:00"
2015/12/17 19:36:04 Hello, World!
2015/12/17 19:36:04 expvars:
2015/12/17 19:36:04 	cmdline -> ["./expvar"]
2015/12/17 19:36:04 	hello -> 1
2015/12/17 19:36:04 	time -> "2015-12-17T19:36:04.644627071+08:00"
2015/12/17 19:36:04 Hello, World!
2015/12/17 19:36:04 expvars:
2015/12/17 19:36:04 	cmdline -> ["./expvar"]
2015/12/17 19:36:04 	hello -> 2
2015/12/17 19:36:04 	time -> "2015-12-17T19:36:04.644695308+08:00"
*/
