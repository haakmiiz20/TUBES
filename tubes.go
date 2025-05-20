package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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
			fmt.Println("||--------------------------------------------||")
			fmt.Println("||            Keluar dari program             ||")
			fmt.Println("||--------------------------------------------||")
			return
		default:
			fmt.Println("||--------------------------------------------||")
			fmt.Println("||            Pilihan tidak valid             ||")
		}
	}
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}
func menu(option *int) {
	fmt.Println("||--------------------------------------------||")
	fmt.Println("||------------------- Menu -------------------||")
	fmt.Println("|| 1. Tambah Aset Kripto                      ||")
	fmt.Println("|| 2. Ubah Aset Kripto                        ||")
	fmt.Println("|| 3. Hapus Aset Kripto                       ||")
	fmt.Println("|| 4. Beli Aset Kripto                        ||")
	fmt.Println("|| 5. Jual Aset Kripto                        ||")
	fmt.Println("|| 6. Lihat Daftar Aset                       ||")
	fmt.Println("|| 7. Lihat Saldo                             ||")
	fmt.Println("|| 8. Lihat Riwayat Transaksi                 ||")
	fmt.Println("|| 0. Keluar                                  ||")
	fmt.Println("||--------------------------------------------||")
	fmt.Println("||--------------------------------------------||")

	fmt.Print("|| Option: ")
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

	fmt.Println("||--------------------------------------------||")
	fmt.Print("|| Isi nama        : ")
	fmt.Scan(&nama)
	fmt.Print("|| Isi harga       : ")
	fmt.Scan(&harga)
	fmt.Print("|| Isi nilai pasar : ")
	fmt.Scan(&nilai)
	for i := 0; i <= *jumlahAset; i++ {
		if (*kripto)[i].namaAset == nama {
			fmt.Println("||--------------------------------------------||")
			fmt.Println("||   Aset sudah ada, tidak bisa ditambahkan   ||")
			find = true
		}
	}
	if !find {
		(*kripto)[*jumlahAset].namaAset = nama
		(*kripto)[*jumlahAset].harga = harga
		(*kripto)[*jumlahAset].nilaiAset = nilai
		*jumlahAset = *jumlahAset + 1
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||         Aset berhasil ditambahkan.         ||")
	}
}
func ubahAset(kripto *arrKripto, jumlahAset int) {

	var nama, namaBaru string
	var hargaBaru float64
	var nilaiBaru float64
	var find bool
	find = false

	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Masukkan Aset yang Ingin diubah:           ||")
	fmt.Print("|| Nama Lama : ")
	fmt.Scan(&nama)

	for i := 0; i < jumlahAset; i++ {
		if kripto[i].namaAset == nama {

			fmt.Println("||--------------------------------------------||")
			fmt.Print("|| Masukkan Nama Baru        : ")
			fmt.Scan(&namaBaru)
			fmt.Print("|| Masukkan Harga Baru       : ")
			fmt.Scan(&hargaBaru)
			fmt.Print("|| Masukkan Nilai Pasar Baru : ")
			fmt.Scan(&nilaiBaru)

			kripto[i].namaAset = namaBaru
			kripto[i].harga = hargaBaru
			kripto[i].nilaiAset = nilaiBaru
			fmt.Println("||--------------------------------------------||")
			fmt.Println("||           Aset berhasil diubah             ||")
			find = true
		}
	}
	if !find {
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||            Aset tidak ditemukan            ||")
	}
}
func hapusAset(kripto *arrKripto, jumlahAset *int) {

	var nama string
	var find bool

	fmt.Println("||--------------------------------------------||")
	fmt.Print("|| Masukkan Aset yang Ingin dihapus: ")
	fmt.Scan(&nama)

	for i := 0; i < *jumlahAset; i++ {
		if kripto[i].namaAset == nama {
			find = true
			for j := i; j < *jumlahAset-1; j++ {
				kripto[j] = kripto[j+1]
			}
			*jumlahAset -= 1
			fmt.Println("||--------------------------------------------||")
			fmt.Println("||           Aset berhasil dihapus            ||")

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

	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Masukkan Aset yang Ingin dibeli :          ||")
	fmt.Print("|| Nama Aset : ")
	fmt.Scan(&namaBeli)
	fmt.Print("|| Jumlah Aset : ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlahAset; i++ {
		if (*kripto)[i].namaAset == namaBeli {
			find = true
			totalHarga = jumlah * kripto[i].harga
			if saldo >= totalHarga {
				saldo = saldo - totalHarga
				fmt.Println("||--------------------------------------------||")
				fmt.Printf("|| Berhasil membeli %.0f unit %s   \n", jumlah, namaBeli)
				fmt.Printf("|| dengan total harga %.0f    \n", totalHarga)
			} else {
				fmt.Println("||--------------------------------------------||")
				fmt.Println("||           Saldo tidak mencukupi            ||")
			}
		}
	}
	if !find {
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||            Aset tidak ditemukan            ||")
	}
	riwayattransaksi(jenis, namaBeli, jumlah, totalHarga)
}
func jualAset(kripto *arrKripto, jumlahAset int) {

	var namaJual, jenis string
	var jumlah, totalHarga float64
	var find bool
	find = false
	jenis = "Jual"
	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Masukkan Aset yang Ingin dijual :          ||")
	fmt.Print("|| Nama Aset : ")
	fmt.Scan(&namaJual)
	fmt.Print("|| Jumlah aset : ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlahAset; i++ {
		if (*kripto)[i].namaAset == namaJual {
			find = true
			totalHarga = jumlah * kripto[i].harga
			saldo = saldo + totalHarga
			fmt.Println("||--------------------------------------------||")
			fmt.Printf("|| Berhasil menjual %.0f unit %s \n", jumlah, namaJual)
			fmt.Printf("|| dengan total harga %.0f \n", totalHarga)
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
	riwayattransaksi(jenis, namaJual, jumlah, totalHarga)
}
func daftarAset(a arrKripto, jumlahAset int) {
	var i int

	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Daftar aset terurut berdasarkan harga :    ||")
	fmt.Println("|| Nama Aset    Harga    NilaiAset            ||")
	for i = 0; i < jumlahAset; i++ {
		fmt.Printf("|| %s %12.0f %8.0f\n", a[i].namaAset, a[i].harga, a[i].nilaiAset)
	}
	fmt.Println()
}
func tampilSaldo() {
	fmt.Println("||--------------------------------------------||")
	fmt.Printf("|| Saldo saat ini :  %.0f                     ||\n", saldo)
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
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||               Tidak ada riwayat        -   ||")
		return
	}
	fmt.Println("||----------- Riwayat Transaksi --------------||")
	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Jenis    NamaAset     Jumlah    TotalHarga ||")
	for i = 0; i < jumlahTrans; i++ {
		fmt.Printf("|| %s %5s %12.0f %9.0f\n", riwayat[i].jenis, riwayat[i].namaAset, riwayat[i].jumlah, riwayat[i].totalHarga)

	}

}
