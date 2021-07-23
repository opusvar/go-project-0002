package main

import f "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber uint
	var grade float32
	title = "Mr. GoToSleep"
	writer = "Tracy Hatchet"
	artist = "Jewel Tampson"  
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	f.Println(title, "published by", publisher, "year", year, "written by", writer, "drawn by", artist, "pages", pageNumber, "Grade", grade)
	// next comic book, publishers are the same
	title = "Epic Vol. 1"
    writer = "Ryan N. Shaw"
	artist = "Phobe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0
	f.Println(title, "published by", publisher, "year", year, "written by", writer, "drawn by", artist, "pages", pageNumber, "Grade", grade)
}