package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var jumlah, cari int
	fmt.Scan(&jumlah, &cari)

	isiAngka(jumlah)

	posisi := cariPosisi(jumlah, cari)

	if posisi != -1 {
		fmt.Println(posisi)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isiAngka(jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&data[i])
	}
}

func cariPosisi(jumlah, cari int) int {
	kiri, kanan := 0, jumlah-1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah] == cari {
			return tengah
		} else if data[tengah] < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}
