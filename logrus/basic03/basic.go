//
// basic.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"os"
	"github.com/Sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instance
var log = logrus.New()

func main() {
	// The API for setting attributes is a little different than the package level
	// exported logger.
	log.Out = os.Stderr

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size": 10,
	}).Info("A group of walrus emerges from the ocean")

}
