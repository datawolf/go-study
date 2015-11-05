//
// library.go
// Copyright (C) 2015 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Name		string
	Authors		[]string
	Publisher	string
	Language	string
	PublishDate	string

	Categories	[]string
	CoverUrl	string

	Pages		int
	Price		float64
	Note		[]string
}

type Library struct {
	Name	string
	Date	string
	Create	string
	Books	[]Book
}


func main() {
	Shu := []Book {
		{
			Name: "Book Name",
			Authors:	[]string {"auther01", "auther02", "author03",},
			Publisher:	"出版社",
			Language:	"Chinese",
			PublishDate:	"2014-05-01",

			Categories:	[]string{"psychology", "心理学通俗读物",},
			CoverUrl:	"http://item.jd.com/11476283.html",

			Pages:		2212,
			Price:		35.50,
			Note:		[]string{"Note 1", "Note2", },
		},
		{
			Name: "正向思考力",
			Authors:	[]string {"Sue Hadfield", "哈德菲尔德", "欧阳瑾",},
			Publisher:	"人民邮电出版社",
			Language:	"Chinese",
			PublishDate:	"2014-05-01",

			Categories:	[]string{"psychology", "心理学通俗读物",},
			CoverUrl:	"http://item.jd.com/11476283.html",

			Pages:		238,
			Price:		35.00,
			Note:		[]string{"抑郁、暴躁、不自信、焦虑、抱怨、否定一切……如果你感觉自己的生活如坠入消极悲观的陷阱，负能量一直在拖你的后腿，那么是时候拥抱积极乐观的思考方式，掌控自己的情绪，改变自己的命运了。"},
		},
	}

	Lib := Library {
		Name: "My Library",
		Date: "2015-11-05",
		Create:	"Wang Long",
		Books: Shu,
	}

	b, err := json.Marshal(Lib)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}
