package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const NMAX = 1000
const jumlahKategori = 8

type NFT struct {
	Nama          string
	Kategori      string
	Harga         float64
	JumlahPemilik int
}

type Transaksi struct {
	NamaNFT string
	Jenis   string
	Harga   float64
	Tanggal string
}

var daftarTransaksi [NMAX]Transaksi
var jumlahTransaksi int

var daftarNFT [NMAX]NFT
var jumlahNFT int

var kategoriNFT = [jumlahKategori]string{
	"Art",
	"Gaming",
	"Memberships",
	"Music",
	"PFPS",
	"Photography",
	"Collectibles",
	"Virtual Worlds",
}

func main() {
	var username string
	start(&username)
	mainmenu(username)
}

func start(username *string) {
	fmt.Println("====================================================================")
	fmt.Println("===   Aplikasi Manajemen Portofolio NFT dan Karya Digital   ===")
	fmt.Println("===            Created Satria Dinata dan Hilwa              ===")
	fmt.Println("===            Algoritma Pemrograman 2025                   ===")

	fmt.Println("====================================================================")
	fmt.Print("Masukkan username: ")
	fmt.Scan(username)
}

func mainmenu(username string) {
	var pilihan int
	fmt.Println("====================================================================")
	fmt.Println("\n=== Menu Utama ===")
	fmt.Println("1. NFT")
	fmt.Println("2. Transaksi")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	fmt.Println("====================================================================")

	switch pilihan {
	case 1:
		menuNFT()
	case 2:
		menuRiwayat()
	case 3:
		fmt.Println(">> Logout. Sampai jumpa,", username)
		return
	}
	fmt.Println("====================================================================")
}

func menuNFT() {
	var subPilihan int
	var username string
	for {
		fmt.Println("====================================================================")
		fmt.Println("\n=== Menu Daftar NFT ===")
		fmt.Println("1. Lihat Kategori ")
		fmt.Println("2. Lihat Daftar NFT")
		fmt.Println("3. Lihat tren Pasar NFT")
		fmt.Println("4. Tambah NFT")
		fmt.Println("5. Ubah NFT")
		fmt.Println("6. Hapus NFT")
		fmt.Println("7. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&subPilihan)
		fmt.Println("====================================================================")

		switch subPilihan {
		case 1:
			tampilkanKategori(jumlahKategori)
		case 2:
			if jumlahNFT == 0 {
				fmt.Println(">> Maaf anda belum memiliki NFT")
			} else {
				tampilkanDaftarNFT(jumlahNFT)
			}
		case 3:
			trendPasarNFT()
		case 4:
			tambahNFT()
		case 5:
			if jumlahNFT == 0 {
				fmt.Println(">> Maaf anda belum memiliki NFT")
			} else {
				ubahNFT()
			}
		case 6:
			if jumlahNFT == 0 {
				fmt.Println(">> Maaf anda belum memiliki NFT")
			} else {
				hapusNFT()
			}
		case 7:
			mainmenu(username)
			return
		}
	}
}

func tampilkanKategori(n int) {
	var i int
	fmt.Println("\n=== Daftar Kategori NFT ===")
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, kategoriNFT[i])
	}
}

func tampilkanDaftarNFT(n int) {
	var pilihan int
	for {
		fmt.Println("\n>> Lihat Daftar NFT")
		fmt.Println("1. Cari berdasarkan Nama")
		fmt.Println("2. Cari berdasarkan Kategori")
		fmt.Println("3. Urutkan berdasarkan Harga")
		fmt.Println("4. Urutkan berdasarkan Jumlah Pemilik")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			cariNFTByNama(n)
		case 2:
			cariNFTByKategori(n)
		case 3:
			menuUrutHarga(n)
		case 4:
			menuUrutPemilik(n)
		case 5:
			return
		default:
			fmt.Println(">> Pilihan tidak valid.")
		}
	}
}

func menuUrutHarga(n int) {
	if n == 0 {
		fmt.Println(">> Tidak ada NFT untuk diurutkan.")
		return
	}
	selectionSortHarga(n)
	cetakDaftar(n)
}

func menuUrutPemilik(n int) {
	if n == 0 {
		fmt.Println(">> Tidak ada NFT untuk diurutkan.")
		return
	}
	insertionSortPemilik(n)
	cetakDaftar(n)
}

func cetakDaftar(n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s - Kategori: %s - Harga: Rp%.2f - Pemilik: %d\n",
			i+1, daftarNFT[i].Nama, daftarNFT[i].Kategori, daftarNFT[i].Harga, daftarNFT[i].JumlahPemilik)
	}
}

func cariNFTByNama(n int) {
	var ditemukan bool
	var nama string
	var i int

	fmt.Print("Masukkan nama NFT yang dicari: ")
	fmt.Scanln(&nama)

	ditemukan = false
	for i = 0; i < n; i++ {
		if strings.EqualFold(daftarNFT[i].Nama, nama) {
			fmt.Println(">> NFT ditemukan:")
			fmt.Printf("- Nama     : %s\n", daftarNFT[i].Nama)
			fmt.Printf("- Kategori : %s\n", daftarNFT[i].Kategori)
			fmt.Printf("- Harga    : Rp%.2f\n", daftarNFT[i].Harga)
			fmt.Printf("- Pemilik  : %d\n", daftarNFT[i].JumlahPemilik)
			ditemukan = true
			break
		}
	}
	if !ditemukan {
		fmt.Println(">> NFT tidak ditemukan.")
	}
}

func urutkanKategori(n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(daftarNFT[j].Kategori) < strings.ToLower(daftarNFT[minIdx].Kategori) {
				minIdx = j
			}
		}
		daftarNFT[i], daftarNFT[minIdx] = daftarNFT[minIdx], daftarNFT[i]
	}
}

func cariNFTByKategori(n int) {
	var kategori string
	fmt.Print("Masukkan kategori NFT yang dicari: ")
	fmt.Scanln(&kategori)

	urutkanKategori(n)

	low := 0
	high := n - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		katMid := strings.ToLower(daftarNFT[mid].Kategori)
		target := strings.ToLower(kategori)

		if katMid == target {
			fmt.Println(">> NFT ditemukan dalam kategori:", kategori)
			i := mid
			for i >= 0 && strings.ToLower(daftarNFT[i].Kategori) == target {
				i--
			}
			i++
			for i < n && strings.ToLower(daftarNFT[i].Kategori) == target {
				fmt.Printf("- %s | Harga: Rp%.2f | Pemilik: %d\n",
					daftarNFT[i].Nama, daftarNFT[i].Harga, daftarNFT[i].JumlahPemilik)
				i++
			}
			found = true
			break
		} else if katMid < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println(">> Tidak ditemukan NFT dalam kategori tersebut.")
	}
}

func selectionSortHarga(n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarNFT[j].Harga < daftarNFT[minIdx].Harga {
				minIdx = j
			}
		}
		daftarNFT[i], daftarNFT[minIdx] = daftarNFT[minIdx], daftarNFT[i]
	}
	fmt.Println(">> Berhasil diurutkan berdasarkan Harga (Selection Sort)")
}

func insertionSortPemilik(n int) {
	for i := 1; i < n; i++ {
		key := daftarNFT[i]
		j := i - 1
		for j >= 0 && daftarNFT[j].JumlahPemilik > key.JumlahPemilik {
			daftarNFT[j+1] = daftarNFT[j]
			j--
		}
		daftarNFT[j+1] = key
	}
	fmt.Println(">> Berhasil diurutkan berdasarkan Jumlah Pemilik (Insertion Sort)")
}

func trendPasarNFT() {
	var hargaLama [NMAX]float64
	var persen float64
	var selisih float64
	var status string

	if jumlahNFT == 0 {
		fmt.Println(">> Belum ada NFT dalam portofolio.")
		return
	}

	fmt.Println("\n=== Tren Pasar NFT ===")

	for i := 0; i < jumlahNFT; i++ {
		hargaLama[i] = daftarNFT[i].Harga
	}

	for i := 0; i < jumlahNFT; i++ {
		persen = rand.Float64()*0.2 - 0.1
		daftarNFT[i].Harga += daftarNFT[i].Harga * persen
	}

	for i := 0; i < jumlahNFT; i++ {
		selisih = daftarNFT[i].Harga - hargaLama[i]
		if selisih >= 0 {
			status = "naik"
		} else {
			status = "turun"
		}
		persen = (selisih / hargaLama[i]) * 100

		fmt.Printf("%d. %s - Sebelumnya: Rp%.2f, Sekarang: Rp%.2f (%s %.2f%%)\n",
			i+1, daftarNFT[i].Nama, hargaLama[i], daftarNFT[i].Harga, status, persen)
	}
}

func tambahNFT() {
	var nft NFT
	var pilihKategori int

	fmt.Println("\n====================================================================")

	if jumlahNFT >= NMAX {
		fmt.Println(">> Kapasitas NFT penuh")
		return
	}

	fmt.Print("Masukkan nama NFT (tanpa spasi): ")
	fmt.Scanln(&nft.Nama)
	tampilkanKategori(jumlahKategori)

	fmt.Print("Pilih kategori (1-8): ")
	fmt.Scanln(&pilihKategori)

	if pilihKategori >= 1 && pilihKategori <= jumlahKategori {
		nft.Kategori = kategoriNFT[pilihKategori-1]
	} else {
		fmt.Println(">> Kategori tidak valid")
		return
	}

	fmt.Print("Masukkan harga NFT: ")
	fmt.Scanln(&nft.Harga)
	if nft.Harga < 0 {
		fmt.Println(">> Harga tidak boleh negatif")
		return
	}

	daftarNFT[jumlahNFT] = nft
	jumlahNFT++
	fmt.Println(">> NFT berhasil ditambahkan")

	daftarTransaksi[jumlahTransaksi] = Transaksi{
		NamaNFT: nft.Nama,
		Jenis:   "Beli",
		Harga:   nft.Harga,
		Tanggal: "17-05-2025",
	}
	jumlahTransaksi++
	fmt.Println(">> Transaksi berhasil dicatat")
}

func ubahNFT() {
	var pilihKategori int
	var pilihan int

	fmt.Println("\n====================================================================")
	tampilkanDaftarNFT(jumlahNFT)

	fmt.Print("Pilih nomor NFT yang akan diubah: ")
	fmt.Scanln(&pilihan)

	if pilihan >= 1 && pilihan <= jumlahNFT {
		fmt.Print("Masukkan nama baru: ")
		fmt.Scanln(&daftarNFT[pilihan-1].Nama)

		tampilkanKategori(jumlahKategori)

		fmt.Print("Pilih kategori baru (1-8): ")
		fmt.Scanln(&pilihKategori)

		if pilihKategori >= 1 && pilihKategori <= jumlahKategori {
			daftarNFT[pilihan-1].Kategori = kategoriNFT[pilihKategori-1]
		}

		fmt.Print("Masukkan harga baru: ")
		fmt.Scanln(&daftarNFT[pilihan-1].Harga)
		fmt.Println(">> NFT berhasil diubah")
	} else {
		fmt.Println(">> Nomor NFT tidak valid")
	}
}

func hapusNFT() {
	var pilihan, i int
	var namaTerhapus, tanggal string
	var hargaTerhapus float64

	fmt.Println("\n====================================================================")
	tampilkanDaftarNFT(jumlahNFT)

	fmt.Print("Pilih nomor NFT yang akan dihapus: ")
	fmt.Scanln(&pilihan)

	if pilihan >= 1 && pilihan <= jumlahNFT {
		namaTerhapus = daftarNFT[pilihan-1].Nama
		hargaTerhapus = daftarNFT[pilihan-1].Harga

		for i = pilihan - 1; i < jumlahNFT-1; i++ {
			daftarNFT[i] = daftarNFT[i+1]
		}
		jumlahNFT--

		tanggal = time.Now().Format("02-01-2006")
		daftarTransaksi[jumlahTransaksi] = Transaksi{
			NamaNFT: namaTerhapus,
			Jenis:   "Jual",
			Harga:   hargaTerhapus,
			Tanggal: tanggal,
		}
		jumlahTransaksi++

		fmt.Println(">> NFT berhasil dihapus")
		fmt.Println(">> Transaksi berhasil dicatat")
	} else {
		fmt.Println(">> Nomor NFT tidak valid")
	}
}

func menuRiwayat() {
	var pilihan int
	for {
		fmt.Println("\n====================================================================")
		fmt.Println("=== Menu Transaksi ===")
		fmt.Println("1. Lihat Riwayat Transaksi")
		fmt.Println("2. Tambah Transaksi Baru")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatRiwayat()
		case 2:
			tambahTransaksi()
		case 3:
			return
		default:
			fmt.Println(">> Pilihan tidak valid.")
		}
	}
}

func lihatRiwayat() {
	var i int
	var t Transaksi

	if jumlahTransaksi == 0 {
		fmt.Println(">> Belum ada transaksi yang tercatat.")
		return
	}
	fmt.Println("\n=== Riwayat Transaksi ===")
	for i = 0; i < jumlahTransaksi; i++ {
		t = daftarTransaksi[i]
		fmt.Printf("%d. [%s] %s - Rp%.2f pada %s\n", i+1, t.Jenis, t.NamaNFT, t.Harga, t.Tanggal)
	}
}

func tambahTransaksi() {
	var t Transaksi

	if jumlahTransaksi >= NMAX {
		fmt.Println(">> Kapasitas transaksi penuh.")
		return
	}

	fmt.Print("Masukkan nama NFT: ")
	fmt.Scanln(&t.NamaNFT)

	fmt.Print("Jenis transaksi (Beli/Jual): ")
	fmt.Scanln(&t.Jenis)

	if t.Jenis != "Beli" && t.Jenis != "Jual" {
		fmt.Println(">> Jenis transaksi tidak valid. Gunakan 'Beli' atau 'Jual'.")
		return
	}

	fmt.Print("Masukkan harga transaksi: ")
	fmt.Scanln(&t.Harga)

	fmt.Print("Masukkan tanggal transaksi (contoh: 17-05-2025): ")
	fmt.Scanln(&t.Tanggal)

	daftarTransaksi[jumlahTransaksi] = t
	jumlahTransaksi++
	fmt.Println(">> Transaksi berhasil dicatat.")
}