//
// main1.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"github.com/lnmx/serial"
	"log"
	"time"
)

/*
var (
	c = make(chan []byte)
)*/

func SerialHandler(ws *websocket.Conn) {
	log.Println("js connected")
	msg := make([]byte, 512)
	msg = []byte("hello world")
	for {
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("wait for serial....")
//		msg =  <- c
		m, err := ws.Write(msg[:len(msg)])
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Send: %s", msg[:m])
		time.Sleep(3 * time.Second)
	}
}

func main() {
	http.Handle("/serial", websocket.Handler(SerialHandler))
	log.Println("Set up Serial")
	//go setupSerial()
	log.Println("Set up Serial Server")
	err := http.ListenAndServe(":8005", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func setupSerial() {
	device := "COM11"
	baud := 115200

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
		if n > 0 {
			log.Println(n, ">", string(buf[:n]))
//			c <- buf
		}
		time.Sleep(2 * time.Second)
	}
}
