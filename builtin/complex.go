//
// complex.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import "log"

func main() {
	c1 := 1.5 + 0.5i
	c2 := complex(1.5, 0.5)

	log.Printf("c1: %v", c1)
	log.Printf("c2: %v", c2)
	log.Printf("c1 == c2: %v", c1 == c2)
	log.Printf("c1 real: %v", real(c1))
	log.Printf("c1 imag: %v", imag(c1))
	log.Printf("c1 + c2 : %v", c1+c2)
	log.Printf("c1 - c2 : %v", c1-c2)
	log.Printf("c1 * c2 : %v", c1*c2)
	log.Printf("c1 / c2 : %v", c1/c2)
	log.Printf("c1 type: %T", c1)

	c3 := complex(float32(1.5), float32(0.5))
	log.Printf("c3 type: %T", c3)
}

/*
2015/12/17 19:24:31 c1: (1.5+0.5i)
2015/12/17 19:24:31 c2: (1.5+0.5i)
2015/12/17 19:24:31 c1 == c2: true
2015/12/17 19:24:31 c1 real: 1.5
2015/12/17 19:24:31 c1 imag: 0.5
2015/12/17 19:24:31 c1 + c2 : (3+1i)
2015/12/17 19:24:31 c1 - c2 : (0+0i)
2015/12/17 19:24:31 c1 * c2 : (2+1.5i)
2015/12/17 19:24:31 c1 / c2 : (1+0i)
2015/12/17 19:24:31 c1 type: complex128
2015/12/17 19:24:31 c3 type: complex64
*/
