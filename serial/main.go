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
	"os"
	"strconv"
	"database/sql"

	"gopkg.in/ini.v1"
	"github.com/lnmx/serial"
	_ "github.com/go-sql-driver/mysql"
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

	setupSerial(device, baud)
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

	// open the databases
	db, err := sql.Open("mysql", "root:123qwe@/salary")
	if err != nil {
		log.Println("can not open databases: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("can not connect databbases: ", err)
	}

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO ss_id set id=?")
	if err != nil {
		log.Println("db prepare error:", err)
	}

	stmtUpdate, err := db.Prepare("UPDATE serial set id=?")
	if err != nil {
		log.Println("db prepare error:", err)
	}

	buf := make([]byte, 512)

	for {
		n, err := port.Read(buf)
		if err != nil {
			log.Println("serial read error:", err)
		}
		// n = the read length
		if n > 0 {
			log.Println(n, ">", string(buf[:n]))
			_, err = stmtIns.Exec(string(buf[:n]))
			if err != nil {
				log.Println("stmtIns exec error:", err)
			}
			_, err = stmtUpdate.Exec(string(buf[:n]))
			if err != nil {
				log.Println("stmtUpdate exec error:", err)
			}
			log.Println("update the db")
		}
	}
}
