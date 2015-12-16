//
// new.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

type Actor struct {
	Name string
}

type Movie struct {
	Title  string
	Actors []*Actor
}

func main() {
	ip := new(int)
	log.Printf("ip type: %T, ip: %v, *ip: %v", ip, ip, *ip)

	m := new(Movie)
	log.Printf("m type: %T, m: %v, *m: %v", m, m, *m)
}

/* result:

2015/12/16 17:46:07 ip type: *int, ip: 0xc82000a300, *ip: 0
2015/12/16 17:46:07 m type: *main.Movie, m: &{ []}, *m: { []}

*/
