//
// main.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := new(tabwriter.Writer)

	// Foramt in tab-separated coulums with a tat stop of 8
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.Debug)
	fmt.Fprintln(w, "NAME\tDESCRIPTION\tSIZE\t.")
	fmt.Fprintln(w, "library/ubuntu\ta offical ubuntu docker images\t100MB\t.")
	fmt.Fprintln(w, "www.example.com/library/ubuntu\tan another offical ububuntu docker images\t101MB\t.")
	fmt.Fprintln(w)
	w.Flush()

	// Format right-aligned in space-seprated columns of minimal width 15 and
	// at least one blank of padding
	w.Init(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "NAME\tDESCRIPTION\tSIZE\t.")
	fmt.Fprintln(w, "library/ubuntu\ta offical ubuntu docker images\t100MB\t.")
	fmt.Fprintln(w, "www.example.com/library/ubuntu\tan another offical ububuntu docker images\t101MB\t.")
	fmt.Fprintln(w)
	w.Flush()
}
