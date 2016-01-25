#! /bin/sh
#
# build.sh
# Copyright (C) 2015 wanglong <wanglong@SZX1000042009>
#
# Distributed under terms of the MIT license.
#


GOOS=windows GOARCH=386 godep go build -ldflags "-H windowsgui"  -o ReadID.exe main.go
