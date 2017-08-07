package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		p := fmt.Println //ekrana yazdırma fonksiyonunu bir değişkene ata
		s := fmt.Scanln	//klavyeden veri okuma fonksiyonunu bir değişkene ata
		m := map[string]string{
			"Monday": "Pazartesi", "Tuesday": "Salı",
			"Wednesday": "Çarşamba", "Thursday": "Perşembe",
			"Friday": "Cuma", "Saturday": "Cumartesi", "Sunday": "Pazar"}

		var input string
		layout := "02-1-2006" // tarih şablonu

		p("Tarih girin: ")
		s(&input)

		str := input

		t, err := time.Parse(layout, str) //string olarak girilen tarihi yukarıdaki şablona göre parse et

		if err != nil{ //Hata durumda hatayı ekrana yazdır
			p(err)
		}

		var day = t.Weekday().String() //parse edilen tarihin haftanın hangi günü olduğunu değişkene ata

		p(m[day])

		p("*********************")
	}
}