//
// Mask.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	// return the default IP mask of the addr
	mask := addr.DefaultMask()

	// return the result of masking the IP address with mask 
	network := addr.Mask(mask)

	// return the number of leading ones and total bits in the mask
	ones, bits := mask.Size()

	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}
