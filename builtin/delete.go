//
// delete.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	m := make(map[string]int)
	log.Println(m)

	m["one"] = 1
	log.Println(m)

	m["two"] = 2
	log.Println(m)

	delete(m, "one")
	log.Println(m)

	delete(m, "one")
	log.Println(m)

	m = nil
	delete(m, "two")
}

/* result
2015/12/16 17:51:14 map[]
2015/12/16 17:51:14 map[one:1]
2015/12/16 17:51:14 map[one:1 two:2]
2015/12/16 17:51:14 map[two:2]
2015/12/16 17:51:14 map[two:2]
*/
