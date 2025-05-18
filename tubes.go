package main

import "fmt"

const NMAX int = 1000

type asetKripto struct {
	namaAset  string
	harga     float64
	nilaiAset float64
}
type arrKripto [NMAX]asetKripto
type transaksi struct {
	jenis      string
	namaAset   string
	jumlah     float64
	totalHarga float64
}
type arrTrans [NMAX]transaksi

var saldo float64 = 1000000
var jumlahTrans int
var riwayat arrTrans

func main() {

	var kripto arrKripto
	var jumlahAset, option int

	jumlahAset = 0
	jumlahTrans = 0
	for {
		menu(&option)
		switch option {
		case 1:
			tambahAset(&kripto, &jumlahAset)

		case 2:
			ubahAset(&kripto, jumlahAset)

		case 3:
			hapusAset(&kripto, &jumlahAset)

		case 4:
			beliAset(&kripto, jumlahAset)

		case 5:
			jualAset(&kripto, jumlahAset)

		case 6:
			selectionSort(&kripto, jumlahAset)
			daftarAset(kripto, jumlahAset)

		case 7:
			tampilSaldo()
		case 8:
			cetakriwayat()
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func menu(option *int) {
	fmt.Println("===== Menu =====")
	fmt.Println("1. Tambah Aset Kripto")
	fmt.Println("2. Ubah Aset Kripto")
	fmt.Println("3. Hapus Aset Kripto")
	fmt.Println("4. Beli Aset Kripto")
	fmt.Println("5. Jual Aset Kripto")
	fmt.Println("6. Lihat Daftar Aset")
	fmt.Println("7. Lihat Saldo")
	fmt.Println("8. Lihat Riwayat Transaksi")
	fmt.Println("0. Keluar")
	fmt.Print("Option: ")
	fmt.Scan(option)
}
func selectionSort(a *arrKripto, jumlahAset int) {
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
		*jumlahAset = *jumlahAset + 1
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

			fmt.Print("Masukkan nama baru:")
			fmt.Scan(&namaBaru)
			fmt.Print("Masukkan harga baru:")
			fmt.Scan(&hargaBaru)
			fmt.Print("Masukkan nilai pasar baru:")
			fmt.Scan(&nilaiBaru)

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

	fmt.Println("Masukkan nama aset yang ingin dihapus: ")
	fmt.Scan(&nama)

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

	var namaBeli, jenis string
	var jumlah, totalHarga float64
	var find bool
	find = false
	jenis = "Beli"

	fmt.Print("Masukkan nama aset yang ingin dibeli : ")
	fmt.Scan(&namaBeli)
	fmt.Print("Masukkan jumlah aset yang ingin dibeli : ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlahAset; i++ {
		if (*kripto)[i].namaAset == namaBeli {
			find = true
			totalHarga = jumlah * kripto[i].harga
			if saldo >= totalHarga {
				saldo = saldo - totalHarga
				fmt.Printf("Berhasil membeli %.2f unit %s dengan total harga %.2f", jumlah, namaBeli, totalHarga)
			} else {
				fmt.Println("Saldo tidak mencukupi")
			}
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
	riwayattransaksi(jenis, namaBeli, jumlah, totalHarga)
}
func jualAset(kripto *arrKripto, jumlahAset int) {

	var namaJual, jenis string
	var jumlah, totalHarga float64
	var find bool
	find = false
	jenis = "Jual"

	fmt.Print("Masukkan nama aset yang ingin dijual : ")
	fmt.Scan(&namaJual)
	fmt.Print("Masukkan jumlah aset yang ingin dijual : ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlahAset; i++ {
		if (*kripto)[i].namaAset == namaJual {
			find = true
			totalHarga = jumlah * kripto[i].harga
			saldo = saldo + totalHarga
			fmt.Printf("Berhasil menjual %.2f unit %s dengan total harga %.2f\n", jumlah, namaJual, totalHarga)
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
	riwayattransaksi(jenis, namaJual, jumlah, totalHarga)
}
func daftarAset(a arrKripto, jumlahAset int) {
	var i int

	fmt.Println("Daftar aset terurut berdasarkan harga : ")
	fmt.Printf("%20s %6s %6s\n", "Nama Aset", "Harga", "NilaiAset")
	for i = 0; i < jumlahAset; i++ {
		fmt.Printf("%14s %12.2f %6.2f\n", a[i].namaAset, a[i].harga, a[i].nilaiAset)
	}
	fmt.Println()
}
func tampilSaldo() {
	fmt.Printf("Saldo saat ini :  %.0f\n", saldo)
}
func riwayattransaksi(jenis, namaAset string, jumlah, totalHarga float64) {
	if jumlahTrans < NMAX {
		riwayat[jumlahTrans].jenis = jenis
		riwayat[jumlahTrans].namaAset = namaAset
		riwayat[jumlahTrans].jumlah = jumlah
		riwayat[jumlahTrans].totalHarga = totalHarga
	}
	jumlahTrans = jumlahTrans + 1
}
func cetakriwayat() {
	var i int
	if jumlahTrans == 0 {
		fmt.Print("Tidak ada riwayat\n")
		return
	}
	fmt.Print("---Riwayat Transaksi---\n")
	fmt.Printf("%10s %10s %10s %10s\n", "Jenis", "Nama Aset", "Jumlah", "Total Harga")
	for i = 0; i < jumlahTrans; i++ {
		fmt.Printf("%10s %10s %6.0f %10.0f\n", riwayat[i].jenis, riwayat[i].namaAset, riwayat[i].jumlah, riwayat[i].totalHarga)

	}

}
