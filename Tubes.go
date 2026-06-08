package main

import (
	"fmt"
)

// +++ APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL +++

var nama [100]string
var visiMisi [100]string
var nomorUrut [100]int
var suara [100]int

var jumlahKandidat int

func tambahKandidat() {

	fmt.Print("\nJumlah kandidat yang ingin ditambahkan: ")
	fmt.Scan(&jumlahKandidat)

	for eel := 0; eel < jumlahKandidat; eel++ {

		fmt.Println("\nData Kandidat ke-", eel+1)

		fmt.Print("Nama Kandidat : ")
		fmt.Scan(&nama[eel])

		fmt.Print("Nomor Urut    : ")
		fmt.Scan(&nomorUrut[eel])

		fmt.Print("Visi Misi     : ")
		fmt.Scan(&visiMisi[eel])

		suara[eel] = 0
	}
}

func tampilKandidat() {

	fmt.Println("\n=== DATA KANDIDAT ===")

	if jumlahKandidat == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}

	for eel := 0; eel < jumlahKandidat; eel++ {

		fmt.Println("\nKandidat ke-", eel+1)
		fmt.Println("Nama        :", nama[eel])
		fmt.Println("Nomor Urut  :", nomorUrut[eel])
		fmt.Println("Visi Misi   :", visiMisi[eel])
		fmt.Println("Jumlah Suara:", suara[eel])
	}
}

func updateKandidat() {

	var cari int
	ketemu := false

	fmt.Print("\nMasukkan nomor urut kandidat yang ingin diedit: ")
	fmt.Scan(&cari)

	for eel := 0; eel < jumlahKandidat; eel++ {

		if nomorUrut[eel] == cari {

			fmt.Println("\nData ditemukan!")

			fmt.Print("Nama Baru       : ")
			fmt.Scan(&nama[eel])

			fmt.Print("Nomor Urut Baru : ")
			fmt.Scan(&nomorUrut[eel])

			fmt.Print("Visi Misi Baru  : ")
			fmt.Scan(&visiMisi[eel])

			fmt.Println("Data berhasil diperbarui!")

			ketemu = true
			break
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan!")
	}
}

func deleteKandidat() {

	var cari int
	ketemu := false

	fmt.Print("\nMasukkan nomor urut kandidat yang ingin dihapus: ")
	fmt.Scan(&cari)

	for eel := 0; eel < jumlahKandidat; eel++ {

		if nomorUrut[eel] == cari {

			for j := eel; j < jumlahKandidat-1; j++ {

				nama[j] = nama[j+1]
				nomorUrut[j] = nomorUrut[j+1]
				visiMisi[j] = visiMisi[j+1]
				suara[j] = suara[j+1]
			}

			jumlahKandidat--

			fmt.Println("Data berhasil dihapus!")

			ketemu = true
			break
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan!")
	}
}

func voting() {

	var pilih int
	var lagi string

	for {

		fmt.Println("\n=== PILIH KANDIDAT ===")

		for eel := 0; eel < jumlahKandidat; eel++ {

			fmt.Println(nomorUrut[eel], ".", nama[eel])
		}

		fmt.Print("Masukkan nomor urut kandidat: ")
		fmt.Scan(&pilih)

		ditemukan := false

		for eel := 0; eel < jumlahKandidat; eel++ {

			if pilih == nomorUrut[eel] {

				suara[eel]++
				ditemukan = true

				fmt.Println("Suara berhasil diberikan ke", nama[eel])
			}
		}

		if !ditemukan {

			fmt.Println("Nomor kandidat tidak ditemukan!")
		}

		fmt.Print("Voting lagi? (y/t): ")
		fmt.Scan(&lagi)

		if lagi != "y" && lagi != "Y" {

			break
		}
	}
}

func totalSuara() int {

	total := 0

	for eel := 0; eel < jumlahKandidat; eel++ {

		total += suara[eel]
	}

	return total
}

func statistik() {

	total := totalSuara()

	fmt.Println("\n=== HASIL SUARA ===")

	for eel := 0; eel < jumlahKandidat; eel++ {

		var persen float64 = 0

		if total != 0 {

			persen = float64(suara[eel]) / float64(total) * 100
		}

		fmt.Printf("%s = %d suara (%.2f%%)\n",
			nama[eel],
			suara[eel],
			persen)
	}

	fmt.Println("Total Suara Masuk:", total)
}

func sequentialSearch() {

	var cari int
	ketemu := false

	fmt.Print("\nMasukkan nomor urut yang dicari: ")
	fmt.Scan(&cari)

	for eel := 0; eel < jumlahKandidat; eel++ {

		if nomorUrut[eel] == cari {

			fmt.Println("\nData ditemukan!")
			fmt.Println("Nama Kandidat :", nama[eel])
			fmt.Println("Visi Misi     :", visiMisi[eel])
			fmt.Println("Jumlah Suara  :", suara[eel])

			ketemu = true
		}
	}

	if !ketemu {

		fmt.Println("Data tidak ditemukan!")
	}
}

func binarySearch() {

	var cari int

	fmt.Print("\nMasukkan nomor urut yang dicari: ")
	fmt.Scan(&cari)

	low := 0
	high := jumlahKandidat - 1
	ketemu := false

	for low <= high {

		mid := (low + high) / 2

		if nomorUrut[mid] == cari {

			fmt.Println("\nData ditemukan!")
			fmt.Println("Nama Kandidat :", nama[mid])
			fmt.Println("Visi Misi     :", visiMisi[mid])
			fmt.Println("Jumlah Suara  :", suara[mid])

			ketemu = true
			break

		} else if nomorUrut[mid] < cari {

			low = mid + 1

		} else {

			high = mid - 1
		}
	}

	if !ketemu {

		fmt.Println("Data tidak ditemukan!")
	}
}

func selectionSort() {

	for eel := 0; eel < jumlahKandidat-1; eel++ {

		max := eel

		for j := eel + 1; j < jumlahKandidat; j++ {

			if suara[j] > suara[max] {

				max = j
			}
		}

		suara[eel], suara[max] = suara[max], suara[eel]
		nama[eel], nama[max] = nama[max], nama[eel]
		nomorUrut[eel], nomorUrut[max] = nomorUrut[max], nomorUrut[eel]
		visiMisi[eel], visiMisi[max] = visiMisi[max], visiMisi[eel]
	}

	fmt.Println("\nData berhasil diurutkan berdasarkan suara terbanyak!")
}

func insertionSort() {

	for eel := 1; eel < jumlahKandidat; eel++ {

		keyNomor := nomorUrut[eel]
		keyNama := nama[eel]
		keyVisi := visiMisi[eel]
		keySuara := suara[eel]

		j := eel - 1

		for j >= 0 && nomorUrut[j] > keyNomor {

			nomorUrut[j+1] = nomorUrut[j]
			nama[j+1] = nama[j]
			visiMisi[j+1] = visiMisi[j]
			suara[j+1] = suara[j]

			j--
		}

		nomorUrut[j+1] = keyNomor
		nama[j+1] = keyNama
		visiMisi[j+1] = keyVisi
		suara[j+1] = keySuara
	}

	fmt.Println("\nData berhasil diurutkan berdasarkan nomor urut!")
}

func main() {

	var pilihan int

	for {

		fmt.Println("\n+++ APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL +++")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Voting")
		fmt.Println("4. Statistik Suara")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("7. Selection Sort (Suara)")
		fmt.Println("8. Insertion Sort (Nomor Urut)")
		fmt.Println("9. Edit Kandidat")
		fmt.Println("10. Hapus Kandidat")
		fmt.Println("11. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahKandidat()

		case 2:
			tampilKandidat()

		case 3:
			voting()

		case 4:
			statistik()

		case 5:
			sequentialSearch()

		case 6:
			binarySearch()

		case 7:
			selectionSort()

		case 8:
			insertionSort()

		case 9:
			updateKandidat()

		case 10:
			deleteKandidat()

		case 11:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}
