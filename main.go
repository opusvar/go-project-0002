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
	f.Println(title, "written by", writer, "drawn by", artist)

}