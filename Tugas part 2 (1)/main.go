package main

import "fmt"

func KonversiNilai(nilai int) string {

	if nilai >= 80 && nilai <= 100 {
		return "A"
	} else if nilai >= 65 && nilai <= 79 {
		return "B+"
	} else if nilai >= 50 && nilai <= 64 {
		return "B"

	} else if nilai >= 35 && nilai <= 49 {
		return "C"
	} else if nilai >= 0 && nilai <= 34 {
		return "D"
	} else {
		return "Nilai Tidak valid"
	}

}

func main() {

	var nilai int
	fmt.Print("masukkan nilai: ")
	fmt.Scanln(&nilai)

	nilaiHuruf := KonversiNilai(nilai)

	fmt.Print("Nilai akan menjadi: ", nilai, (nilaiHuruf))

}
