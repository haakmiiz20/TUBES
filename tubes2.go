package main

import (
	"fmt"
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
	status     string
}
type arrTrans [NMAX]transaksi

const akunMax = 100

type akun struct {
	username string
	password string
	saldo    float64
}

var jumlahTrans int
var riwayat arrTrans

func main() {
	var kripto arrKripto
	var jumlahAset, option int
	var daftarAkun [akunMax]akun
	var jumlahAkun int
	var akunAktif *akun
	jumlahAkun = 0
	jumlahAset = 0
	jumlahTrans = 0
	akunAktif = nil
	for akunAktif == nil {
		var pilih int
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||----------------- Menu Awal ----------------||")
		fmt.Println("|| 1. Login                                   ||")
		fmt.Println("|| 2. Register                                ||")
		fmt.Println("|| 0. Keluar                                  ||")
		fmt.Print("|| Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			akunAktif = login(&daftarAkun, jumlahAkun)
		case 2:
			register(&daftarAkun, &jumlahAkun)
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

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
			beliAset(&kripto, jumlahAset, akunAktif)
		case 5:
			jualAset(&kripto, jumlahAset, akunAktif)
		case 6:
			selectionSort(&kripto, jumlahAset)
			daftarAset(kripto, jumlahAset)
		case 7:
			tampilSaldo(akunAktif)
		case 8:
			cetakriwayat()
		case 9:
			isiSaldo(akunAktif)
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
func menu(option *int) {
	//{ I.S: variabel optoon belum menerima input angka dari pengguna
	//F.S: variabel option menerima input angka dari pengguna yang dipilih dari menu}
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
	fmt.Println("|| 9. Isi Saldo Akun                          ||")
	fmt.Println("|| 0. Keluar                                  ||")
	fmt.Println("||--------------------------------------------||")
	fmt.Println("||--------------------------------------------||")
	fmt.Print("|| Option: ")
	fmt.Scan(option)
}
func register(daftarAkun *[akunMax]akun, jumlahAkun *int) {
	var uname, pass string
	fmt.Println("||--------------- Register Akun --------------||")
	fmt.Print("|| Buat Username: ")
	fmt.Scan(&uname)
	fmt.Print("|| Buat Password: ")
	fmt.Scan(&pass)

	for i := 0; i < *jumlahAkun; i++ {
		if daftarAkun[i].username == uname {
			fmt.Println("|| Username sudah digunakan.")
			return
		}
	}

	daftarAkun[*jumlahAkun] = akun{username: uname, password: pass, saldo: 0}
	*jumlahAkun++
	fmt.Println("|| Akun berhasil dibuat.")
}

func login(daftarAkun *[akunMax]akun, jumlahAkun int) *akun {
	var uname, pass string
	fmt.Println("||---------------- Login Akun ----------------||")
	fmt.Print("|| Username: ")
	fmt.Scan(&uname)
	fmt.Print("|| Password: ")
	fmt.Scan(&pass)

	for i := 0; i < jumlahAkun; i++ {
		if daftarAkun[i].username == uname && daftarAkun[i].password == pass {
			fmt.Println("|| Login berhasil.")
			return &daftarAkun[i]
		}
	}
	fmt.Println("|| Login gagal. Username atau password salah.")
	return nil
}
func isiSaldo(akunAktif *akun) {
	var tambah float64
	fmt.Println("||----------------- Isi Saldo ----------------||")
	fmt.Print("|| Masukkan jumlah saldo: ")
	fmt.Scan(&tambah)

	if tambah > 0 {
		akunAktif.saldo += tambah
		fmt.Printf("|| Saldo berhasil ditambah sebesar %.0f\n", tambah)
	} else {
		fmt.Println("|| Jumlah tidak valid.")
	}
}

func selectionSort(a *arrKripto, jumlahAset int) {
	//{ I.S: array a berisi data aset kripto yang tidak terurut
	//F.S: array a terurut menurun (descending) berdasarkan atribut harga}
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
func insertionSort(a *arrKripto, jumlahAset int) {
	//{ I.S array a berisi data aset kripto dalam urutan tidak tertentu
	//  F.S array a terurut menaik (ascending) berdasarkan atribut harga.}
	var i, pass int
	var temp asetKripto

	pass = 1
	for pass <= jumlahAset-1 {
		i = pass
		temp.harga = a[pass].harga
		for i > 0 && temp.harga < a[i-1].harga {
			a[i].harga = a[i-1].harga
			i = i - 1
		}
		a[i].harga = temp.harga
		pass = pass + 1
	}
}
func binarySearch(a arrKripto, jumlahAset int, nama string) int {
	//{ I.S array a tidak terurut berdasarkan namaAset, nama adalah input dari pengguna.
	//  F.S array a terurut menaik berdasarkan namaAset, program mencetak informasi aset jika ditemukan, atau pesan tidak ditemukan}
	var i, pass int
	var temp asetKripto

	pass = 1
	for pass <= jumlahAset-1 {
		i = pass
		temp.namaAset = a[pass].namaAset
		for i > 0 && temp.namaAset < a[i-1].namaAset {
			a[i].namaAset = a[i-1].namaAset
			i = i - 1
		}
		a[i].namaAset = temp.namaAset
		pass = pass + 1
	}

	var low, high, mid int
	low = 0
	high = jumlahAset - 1

	for low <= high {
		mid = (low + high) / 2
		if a[mid].namaAset == nama {
			fmt.Println("||--------------------------------------------||")
			fmt.Printf(" Aset ditemukan: %s | Harga: %.0f | Nilai Pasar: %.0f ||\n", a[mid].namaAset, a[mid].harga, a[mid].nilaiAset)
			return mid
		} else if a[mid].namaAset < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("||--------------------------------------------||")
	fmt.Println("||         Aset tidak ditemukan               ||")
	return -1
}
func tambahAset(kripto *arrKripto, jumlahAset *int) {

	//{ I.S array kripto berisi sejumlah aset kripto dan jumlahAset menyatakan jumlah aset saat ini.
	//F.S Jika nama belum ada, data aset baru ditambahkan ke array kripto dan jumlahAset bertambah satu.}
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

	//{ I.S array kripto berisi aset kripto, dan pengguna memasukkan data aset yang ingin diubah.
	//F.S Jika ditemukan, data aset tersebut diubah dengan data baru yang dimasukkan pengguna.}
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

	//{ I.S array kripto berisi aset kripto, dan pengguna memasukkan nama aset yang akan dihapus.
	//F.S Jika ditemukan, aset kripto akan dihapus dari array kripto dan array jumlahAset berkurang.}
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
func beliAset(kripto *arrKripto, jumlahAset int, akunAktif *akun) {

	//{ I.S pengguna memasukkan nama aset yang ingin dibeli dan jumlah pembelian. Saldo dan data aset sudah ada.
	//F.S Jika saldo cukup dan aset ditemukan, saldo berkurang, transaksi dicatat dalam riwayat}
	var namaBeli, jenis, status string
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
			if akunAktif.saldo >= totalHarga {
				akunAktif.saldo = akunAktif.saldo - totalHarga
				fmt.Println("||--------------------------------------------||")
				fmt.Printf("|| Berhasil membeli %.0f unit %s   \n", jumlah, namaBeli)
				fmt.Printf("|| dengan total harga %.0f    \n", totalHarga)
				status = "Berhasil"
			} else {
				fmt.Println("||--------------------------------------------||")
				fmt.Println("||           Saldo tidak mencukupi            ||")
				status = "Gagal"
			}
		}
	}
	if !find {
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||            Aset tidak ditemukan            ||")
	}
	riwayattransaksi(jenis, namaBeli, jumlah, totalHarga, status)
}
func jualAset(kripto *arrKripto, jumlahAset int, akunAktif *akun) {

	//{ I.S pengguna memasukkan nama aset yang ingin dijual dan jumlah penjualan.
	//F.S Jika aset ditemukan, saldo bertambah, transaksi dicatat dalam riwayat}
	var namaJual, jenis, status string
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
			akunAktif.saldo = akunAktif.saldo + totalHarga
			fmt.Println("||--------------------------------------------||")
			fmt.Printf("|| Berhasil menjual %.0f unit %s \n", jumlah, namaJual)
			fmt.Printf("|| dengan total harga %.0f \n", totalHarga)
			status = "Berhasil"
		}
	}
	if !find {
		fmt.Println("Aset tidak ditemukan.")
	}
	riwayattransaksi(jenis, namaJual, jumlah, totalHarga, status)
}
func daftarAset(a arrKripto, jumlahAset int) {
	//{ I.S Array a berisi data aset.
	//F.S menampilkan daftar aset yag sudah dibeli dan dijual oleh pengguna.}
	var i int

	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Daftar aset terurut berdasarkan harga :    ||")
	fmt.Println("|| Nama Aset    Harga    NilaiAset            ||")
	for i = 0; i < jumlahAset; i++ {
		fmt.Printf("|| %s %12.0f %8.0f\n", a[i].namaAset, a[i].harga, a[i].nilaiAset)
	}
	fmt.Println()
}
func tampilSaldo(akunAktif *akun) {
	//{ I.S saldo sudah dideklrarasikan.
	//F.S menampilkan nilai saldo saat ini ke layar.}
	fmt.Println("||--------------------------------------------||")
	fmt.Printf("|| Saldo saat ini :  %.0f \n", akunAktif.saldo)
}
func riwayattransaksi(jenis, namaAset string, jumlah, totalHarga float64, status string) {
	//{ I.S jumlahTrans berisi jumlah transaksi yang telah dilakukan.
	//F.S Menambahkan data transaksi baru ke array riwayat, dan jumlahTrans bertambah satu.}
	if jumlahTrans < NMAX {
		riwayat[jumlahTrans].jenis = jenis
		riwayat[jumlahTrans].namaAset = namaAset
		riwayat[jumlahTrans].jumlah = jumlah
		riwayat[jumlahTrans].totalHarga = totalHarga
		riwayat[jumlahTrans].status = status
	}
	jumlahTrans = jumlahTrans + 1
}
func cetakriwayat() {
	//{ I.S Array riwayat berisi transaksi yang sudah dilakukan (bisa kosong).
	//  F.S Menampilkan semua data transaksi ke layar, atau pesan kosong jika belum ada transaksi.}
	var i int
	if jumlahTrans == 0 {
		fmt.Println("||--------------------------------------------||")
		fmt.Println("||               Tidak ada riwayat        -   ||")
		return
	}
	fmt.Println("||----------- Riwayat Transaksi --------------||")
	fmt.Println("||--------------------------------------------||")
	fmt.Println("|| Jenis    NamaAset     Jumlah    TotalHarga    Status")
	for i = 0; i < jumlahTrans; i++ {
		fmt.Printf("|| %s %5s %12.0f %9.0f %8s\n", riwayat[i].jenis, riwayat[i].namaAset, riwayat[i].jumlah, riwayat[i].totalHarga, riwayat[i].status)
	}
}
