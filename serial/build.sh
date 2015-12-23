#! /bin/sh
#
# build.sh
# Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
#
# Distributed under terms of the MIT license.
#


GOOS=windows GOARCH=386 go build -o SerialServer.exe main.go
