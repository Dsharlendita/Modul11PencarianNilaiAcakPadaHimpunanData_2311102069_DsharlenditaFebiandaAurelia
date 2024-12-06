package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	fmt.Println("Jumlah suara yang masuk:", totalSuara)
	fmt.Println("Jumlah suara sah:", suaraSah)

	// Mengurutkan suara berdasarkan jumlah suara dan nomor calon
	type calon struct {
		nomor int
		suara int
	}

	calonList := []calon{}
	for i := 1; i <= 20; i++ {
		if jumlahSuara[i] > 0 {
			calonList = append(calonList, calon{nomor: i, suara: jumlahSuara[i]})
		}
	}

	// Urutkan berdasarkan suara (descending), jika sama, berdasarkan nomor (ascending)
	sort.Slice(calonList, func(i, j int) bool {
		if calonList[i].suara == calonList[j].suara {
			return calonList[i].nomor < calonList[j].nomor
		}
		return calonList[i].suara > calonList[j].suara
	})

	if len(calonList) > 0 {
		ketua := calonList[0]
		fmt.Printf("Ketua RT: Calon %d dengan %d suara\n", ketua.nomor, ketua.suara)

		if len(calonList) > 1 {
			wakil := calonList[1]
			fmt.Printf("Wakil Ketua RT: Calon %d dengan %d suara\n", wakil.nomor, wakil.suara)
		} else {
			fmt.Println("Tidak ada wakil ketua karena hanya ada satu calon dengan suara sah.")
		}
	} else {
		fmt.Println("Tidak ada suara sah, tidak ada ketua dan wakil yang terpilih.")
	}
}
