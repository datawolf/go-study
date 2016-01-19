//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lnmx/serial"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("Load config error: ", err)
		os.Exit(-1)
	}

	device := cfg.Section("serial").Key("DEVICE").String()
	baud, _ := strconv.Atoi(cfg.Section("serial").Key("BAUD").String())
	log.Println("device", device, "baud", baud)

	go serialServer()
	setupSerial(device, baud)
}

func serialServer(){
	http.Handle("/id", getID)
	http.ListenAndServe(":8080", nil)
}

func getID(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "id=1232332")
}

func setupSerial(device string, baud int) {
	// open the serial
	log.Println("open", device, "at", baud)
	port, err := serial.Open(device, baud)
	if err != nil {
		fmt.Println("open failed:", err)
	}
	defer port.Close()
	log.Println("ready")

	buf := make([]byte, 512)

	for {
		n, err := port.Read(buf)
		if err != nil {
			log.Println("serial read error:", err)
		}
		// n = the read length
		if n > 0 {
			log.Println(n, "接收到刷卡信息--> ", string(buf[:n]))
		}
	}
}
