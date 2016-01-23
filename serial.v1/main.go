//
// main.go
// Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"bytes"
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
	Ic uint64 `json:"id"`
}

var (
	num = uint64(0)
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
	http.HandleFunc("/id", http.HandlerFunc(getID))
	http.ListenAndServe(":8080", nil)
}

func getID(res http.ResponseWriter, req *http.Request) {
	id := Id{num}
	log.Println("id", id)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(res).Encode(id); err != nil {
		log.Println("Encode error")
	}
	num = 0
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
	var buffer bytes.Buffer
	i := 0

	for {
		i = 0
		buffer.Truncate(0)
		for i != 10 && i < 10 {
			// n = the read length
			n, err := port.Read(buf)
			if err != nil {
				log.Println("serial read error:", err)
			}
			time.Sleep(time.Second)
			log.Println(n, "接收到刷卡信息--> ", string(buf[:n]))
			buffer.Write(buf[:n])
			i += n
		}
		if i != 10 {
			continue
		}
		num = parserID(buffer.String())
	}
}

func parserID(res string) uint64 {
	var s uint64
	data := []byte(res)
	if len(res) != 10 {
		s = 0
		return s
	}
	log.Println("before change data:", data)
	//change D1EC80AC to AC80ECD1
	data[0], data[6] = data[6], data[0]
	data[1], data[7] = data[7], data[1]
	data[2], data[4] = data[4], data[2]
	data[3], data[5] = data[5], data[3]
	log.Println("after change data:", data)

	card := string(data[:8])
	log.Println("before parserID:", card)
	if s, err := strconv.ParseUint(card, 16, 64); err == nil {
		log.Printf("parserID: %T, %v", s, s)
	}

	return s
}
