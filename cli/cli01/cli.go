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
	cli.NewApp().Run(os.Args)
}
