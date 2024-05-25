package main

import "fmt"

func main() {

	var HargaAsli float64 = 370000
	var diskon = 0.15
	var jumlahDiskon float64
	var hargaSetelahDiskon float64

	jumlahDiskon = HargaAsli * diskon

	hargaSetelahDiskon = HargaAsli - jumlahDiskon

	fmt.Println(jumlahDiskon)
	fmt.Println(hargaSetelahDiskon)
}
