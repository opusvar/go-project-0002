package main

import f "fmt"
import t "time"

func main() {
	var publisher, writer, artist, title, genre string
	var year, pageNumber, age uint
	var grade float32
	title = "Mr. GoToSleep"
	writer = "Tracy Hatchet"
	artist = "Jewel Tampson"  
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	genre = "fiction"
	grade = 6.5
	age = uint(t.Now().Year()) - year
	f.Println(title, "published by", publisher, "year", year, "age", age, "written by", writer, "drawn by", artist, "genre", genre, "pages", pageNumber, "Grade", grade)
	// next comic book, publishers are the same
	title = "Epic Vol. 1"
    writer = "Ryan N. Shaw"
	artist = "Phobe Paperclips"
	year = 2013
	genre = "Compliation"
	pageNumber = 160
	grade = 9.0
	f.Println(title, "published by", publisher, "year", year, "written by", writer, "drawn by", artist, "pages", pageNumber, "Grade", grade)
}