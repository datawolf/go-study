//
// basic.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"log/syslog"
	"github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.TextFormatter) // default
	hook, err := logrus_syslog.NewSyslogHook("udp", "127.0.0.1:514", syslog.LOG_INFO, "")

	if err != nil {
		log.Hooks.Add(hook)
	}
}

func main() {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Fatal("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"temperature": -4,
	}).Info("Temperature changes")
}
