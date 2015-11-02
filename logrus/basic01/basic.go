//
// basic.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	log "github.com/Sirupsen/logrus"
)

// The simplest way to user Logrus is simply the package-level exported logger.
func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
