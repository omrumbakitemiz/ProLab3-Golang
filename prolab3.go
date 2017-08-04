package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println //ekrana yazdırma fonksiyonunu bir değişkene ata
	s := fmt.Scanln	//klavyeden veri okuma fonksiyonunu bir değişkene ata

	var input string
	layout := "02-01-2006" // tarih şablonu

	p("Tarih girin: ")
	s(&input)
	p("Girdiğiniz tarih: " + input)

	str := input

	t, err := time.Parse(layout, str) //string olarak girilen tarihi yukarıdaki şablona göre parse et

	if err != nil{ //Hata durumda hatayı ekrana yazdır
		p(err)
	}

	p(t.Weekday().String()) //parse edilen tarihin haftanın hangi günü olduğunu ekrana yazdır
}