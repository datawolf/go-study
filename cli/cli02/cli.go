//
// cli.go
// Copyright (C) 2015 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "Make an explosive entrance"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	app.Run(os.Args)
}
