//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/lnmx/serial"
	"gopkg.in/ini.v1"
)

type Id struct {
	ic string `json:"id"`
}

var (
	num = make(chan string, 1)
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

func serialServer() {
	http.HandleFunc("/id", getID)
	http.ListenAndServe(":8080", nil)
}

func getID(res http.ResponseWriter, req *http.Request) {
	temp_num := ""
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1*time.Second)
		timeout <- true
	}()

	select {
	case temp_num = <-num:
		log.Println("get id ok")
	case <-timeout:
		temp_num = "404"
		log.Println("get id timeout")
	}

	id := Id{
		ic: temp_num,
	}

	log.Println("id:", id)
	log.Println("temp_num:", temp_num)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(id); err != nil {
		log.Println("Encode error")
	}
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
	timeout := make(chan bool, 1)
	for {
		n, err := port.Read(buf)
		if err != nil {
			log.Println("serial read error:", err)
		}

		go func() {
			time.Sleep(1*time.Second)
			timeout <- true
		}()

		select {
		case num <- string(buf[:n]):
			log.Println("setupserail ok")
		case <-timeout:
			log.Println("setupserail timeout")
		}
		// n = the read length
		if n > 0 {
			log.Println(n, "接收到刷卡信息--> ", string(buf[:n]))
		}
	}
}
