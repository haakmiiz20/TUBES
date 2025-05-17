package main

import "fmt"

const NMAX int = 1000

type asetKripto struct {
	namaAset  string
	harga     float64
	nilaiAset float64
}
type arrKripto [NMAX]asetKripto

var saldo float64 = 1000000

func main() {

	var kripto arrKripto
	var jumlahAset int
	var option int
	jumlahAset = 0

	menu(&option)

	switch option {
	case 1:
		tambahAset(&kripto, &jumlahAset)
		menu(&option)
	case 2:
		ubahAset(&kripto, jumlahAset)
		menu(&option)
	case 3:
		hapusAset(&kripto, &jumlahAset)
		menu(&option)
	case 4:
		beliAset(&kripto, jumlahAset)
		menu(&option)
	case 5:
		jualAset(&kripto, jumlahAset)
		menu(&option)
	case 6:
		riwayat(kripto, jumlahAset)
		menu(&option)
	case 0:
		fmt.Println("Keluar dari program.")
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}

}
func menu(option *int) {
	fmt.Println("===== Menu =====")
	fmt.Println("1. Tambah Aset Kripto")
	fmt.Println("2. Ubah Aset Kripto")
	fmt.Println("3. Hapus Aset Kripto")
	fmt.Println("4. Beli Aset Kripto")
	fmt.Println("5. Jual Aset Kripto")
	fmt.Println("6. Lihat Riwayat")
	fmt.Println("0. Keluar")
	fmt.Print("Option: ")
	fmt.Scan(option)

}
func selectionSort(a arrKripto, jumlahAset int) {
	var i, index, pass int
	var temp asetKripto
	pass = 1
	for pass < jumlahAset {
		index = pass - 1
		i = pass
		for i < jumlahAset {
			if a[i].harga > a[index].harga {
				index = i
			}
			i = i + 1
		}
		temp = a[pass-1]
		a[pass-1] = a[index]
		a[index] = temp
		pass = pass + 1
	}
}

func tambahAset(kripto *arrKripto, jumlahAset *int) {

	var nama string
	var harga float64
	var nilai float64
	var find bool
	find = false

	fmt.Print("Isi nama, harga dan nilai pasar : ")
	fmt.Scan(&nama, &harga, &nilai)
	for i := 0; i <= *jumlahAset; i++ {
		if (*kripto)[i].namaAset == nama {
			fmt.Println("Aset sudah ada, tidak bisa ditambahkan.")
			find = true
		}
	}
	if !find {
		(*kripto)[*jumlahAset].namaAset = nama
		(*kripto)[*jumlahAset].harga = harga
		(*kripto)[*jumlahAset].nilaiAset = nilai
		*jumlahAset++
		fmt.Println("Aset berhasil ditambahkan.")
	}
}

func ubahAset(kripto *arrKripto, jumlahAset int) {

	var nama, namaBaru string
	var hargaBaru float64
	var nilaiBaru float64
	var find bool
	find = false

	fmt.Println("Masukkan nama aset yang ingin diubah:")
	fmt.Scan(&nama)

	for i := 0; i < jumlahAset; i++ {
		if kripto[i].namaAset == nama {

			fmt.Println("Masukkan nama baru:")
			fmt.Scanln(&namaBaru)
			fmt.Println("Masukkan harga baru:")
			fmt.Scanln(&hargaBaru)
			fmt.Println("Masukkan nilai pasar baru:")
			fmt.Scanln(&nilaiBaru)

			kripto[i].namaAset = namaBaru
			kripto[i].harga = hargaBaru
			kripto[i].nilaiAset = nilaiBaru

			fmt.Println("Aset berhasil diubah.")
			find = true
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
}

func hapusAset(kripto *arrKripto, jumlahAset *int) {

	var nama string
	var find bool

	fmt.Println("Masukkan nama aset yang ingin dihapus:")
	fmt.Scanln(&nama)

	for i := 0; i < *jumlahAset; i++ {
		if kripto[i].namaAset == nama {
			find = true
			for j := i; j < *jumlahAset-1; j++ {
				kripto[j] = kripto[j+1]
			}
			*jumlahAset -= 1
			fmt.Println("Aset berhasil dihapus.")

		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
}

func beliAset(kripto *arrKripto, jumlahAset int) {

	var namaBeli string
	var jumlah, totalHarga float64
	var find bool
	find = false

	fmt.Print("Masukkan nama aset yang ingin dibeli")
	fmt.Scan(&namaBeli)
	fmt.Print("Masukkan jumlah aset yang ingin dibeli")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlahAset; i++ {
		if (*kripto)[i].namaAset == namaBeli {
			find = true
			totalHarga = jumlah * kripto[i].harga
			if saldo >= totalHarga {
				saldo -= totalHarga
				fmt.Printf("Berhasil membeli %.2f unit %s dengan total harga %.2f", jumlah, namaBeli, totalHarga)
			} else {
				fmt.Println("Saldo tidak mencukupi")
			}
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
}

func jualAset(kripto *arrKripto, jumlahAset int) {

	var namaJual string
	var jumlah, totalHarga float64
	var find bool
	find = false

	fmt.Print("Masukkan nama aset yang ingin dijual")
	fmt.Scan(&namaJual)
	fmt.Print("Masukkan jumlah aset yang ingin dijual")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlahAset; i++ {
		if (*kripto)[i].namaAset == namaJual {
			find = true
			totalHarga = jumlah * kripto[i].harga
			saldo += totalHarga
			fmt.Printf("Berhasil menjual %.2f unit %s dengan total harga %.2f", jumlah, namaJual, totalHarga)
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
}
func riwayat(a arrKripto, jumlahAset int) {
	var i int
	selectionSort(kripto, jumlahAset)
	fmt.Println("Daftar aset terurut berdasarkan harga : ")
	fmt.Printf("%20s %6s %6s", "Nama Aset", "Harga", "NilaiAset")
	for i = 0; i < jumlahAset; i++ {
		fmt.Printf("%20s %6d %d", a[i].namaAset, a[i].harga, a[i].nilaiAset)
	}
	fmt.Println()
}
