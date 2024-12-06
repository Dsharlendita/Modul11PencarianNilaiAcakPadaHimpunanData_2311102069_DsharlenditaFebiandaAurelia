package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	baca := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan daftar suara (pisahkan dengan spasi, akhiri dengan angka 0):")

	input, _ := baca.ReadString('\n')
	input = strings.TrimSpace(input)

	daftarSuara := strings.Fields(input)

	totalSuara := 0
	suaraSah := 0
	suaraTidakSah := 0

	jumlahSuara := make([]int, 21) 

	for _, suara := range daftarSuara {
		angka, err := strconv.Atoi(suara)
		if err != nil {
			suaraTidakSah++ 
			continue
		}

		if angka == 0 {
			break
		}

		totalSuara++ 

		if angka >= 1 && angka <= 20 {
			jumlahSuara[angka]++ 
			suaraSah++
		} else {
			suaraTidakSah++ 
		}
	}

	fmt.Println("Jumlah suara yang masuk: ", totalSuara)
	fmt.Println("Jumlah suara sah: ", suaraSah)

	fmt.Println("Hasil perolehan suara:")
	for i := 1; i <= 20; i++ {
		if jumlahSuara[i] > 0 {
			fmt.Printf("Calon %d: %d suara\n", i, jumlahSuara[i])
		}
	}
}
