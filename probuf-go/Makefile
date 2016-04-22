#
# Makefile
# wanglong, 2016-04-22 15:41
#

all: go

go:	add_person  list_people

clean:
	rm -fr add_person  list_people *.pb.go protoc_middleman
	rm -fr tutorial

protoc_middleman: proto/addressbook.proto
	mkdir tutorial # make diectory for go package
	protoc -I=./proto --go_out=tutorial  proto/addressbook.proto
	@touch protoc_middleman

add_person: add_person.go protoc_middleman
	go build -o add_person add_person.go

list_people: list_person.go protoc_middleman
	go build -o list_people list_person.go
