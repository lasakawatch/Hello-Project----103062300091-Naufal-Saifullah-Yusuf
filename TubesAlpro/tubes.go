package main

import (
	"bufio"     // Paket bufio digunakan untuk membaca input dari pengguna
	"fmt"       // Paket fmt digunakan untuk mencetak keluaran ke konsol
	"math/rand" // Paket rand digunakan untuk menghasilkan angka acak
	"os"        // Paket os digunakan untuk berinteraksi dengan sistem operasi
	"sort"      // Paket sort digunakan untuk mengurutkan data
	"strconv"   // Paket strconv digunakan untuk konversi tipe data
	"strings"   // Paket strings digunakan untuk manipulasi string
	"time"      // Paket time digunakan untuk berinteraksi dengan waktu

	"github.com/olekukonko/tablewriter" // Paket tablewriter digunakan untuk membuat tabel di konsol
)

// Mendefinisikan setiap variabel pada struktur pengguna
type Pengguna struct { // Struktur pengguna, mewakili pengguna dalam sistem
	Username    string              // Username pengguna dengan tipe data string
	Password    string              // Password pengguna dengan tipe data string
	Peran       string              // Peran pengguna (misalnya, siswa, guru) dengan tipe data string
	JawabanKuis map[string][]string // Peta jawaban kuis pengguna. Kunci: judul kuis, Nilai: daftar jawaban kuis
	Postingan   map[string][]string // Peta postingan pengguna di forum. Kunci: judul topik, Nilai: daftar postingan dalam topik tersebut
	NilaiKuis   map[string]int      // Peta nilai kuis pengguna. Kunci: judul kuis, Nilai: nilai kuis
	Tugas       map[string]string   // Peta tugas yang dikirimkan oleh pengguna. Kunci: judul tugas, Nilai: konten tugas
	NilaiTugas  map[string]int      // Peta nilai tugas pengguna. Kunci: judul tugas, Nilai: nilai tugas
	NilaiUTS    int                 // Nilai UTS pengguna dengan tipe data integer
}

// Mendefinisikan setiap variabel pada struktur Guru
type Guru struct { // Struktur Guru, mewakili data seorang guru
	*Pengguna                            // Struct tertanam Pengguna, menunjukkan Guru adalah tipe khusus dari Pengguna
	Penilaian  map[string]map[string]int // Peta untuk menyimpan penilaian siswa. Kunci pertama: mata pelajaran atau kursus, Nilai: peta kedua. Kunci peta kedua: username siswa, Nilai: nilai siswa
	NilaiTugas map[string]map[string]int // Peta untuk menyimpan nilai tugas siswa. Kunci: username siswa, Nilai: peta kedua. Kunci peta kedua: judul tugas, Nilai: nilai untuk tugas tersebut
}

// LMS ...
type LMS struct { // Struktur LMS
	Pengguna map[string]interface{} // Peta untuk menyimpan data pengguna. Kunci: username, Nilai: data pengguna yang bisa berupa tipe apa saja (menggunakan interface{})
	Materi   map[string]string      // Peta untuk menyimpan materi pembelajaran. Kunci: judul materi, Nilai: konten materi
	Kuis     map[string][]string    // Peta untuk menyimpan kuis. Kunci: judul kuis, Nilai: daftar pertanyaan dalam kuis tersebut
	Forum    map[string][]string    // Peta untuk menyimpan forum diskusi. Kunci: judul topik forum, Nilai: daftar posting dalam topik tersebut
	Tugas    map[string]string      // Peta untuk menyimpan tugas. Kunci: judul tugas, Nilai: konten tugas
}

// printTable ...
func printTable(header []string, data [][]string) {
	// Membuat objek tablewriter baru dengan output ke output standar
	table := tablewriter.NewWriter(os.Stdout)

	// Mengatur header tabel
	table.SetHeader(header)

	// Menambahkan data ke dalam tabel
	table.AppendBulk(data)

	// Merender tabel dan mencetaknya ke output standar
	table.Render()
}

// printMenu
func printMenu(menuItems []string) {
	fmt.Println() // Mencetak baris kosong untuk memberikan jarak visual
	for i, item := range menuItems {
		// Mencetak nomor menu dan item menu
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Pilihan Anda: ") // Menampilkan pesan untuk meminta input pilihan
}

// Daftar
func (lms *LMS) Daftar(username, password, peran string) {
	username = strings.ReplaceAll(username, " ", "") // Menghapus spasi dari username
	fmt.Println("Mendaftarkan pengguna...")

	// Memeriksa apakah username sudah ada dalam sistem
	if _, ok := lms.Pengguna[username]; !ok {
		// Jika username belum ada, membuat pengguna baru sesuai peran
		if peran == "guru" {
			// Jika peran adalah guru, membuat objek Guru
			lms.Pengguna[username] = &Guru{
				Pengguna: &Pengguna{
					Username:    username,
					Password:    password,
					Peran:       peran,
					JawabanKuis: make(map[string][]string),
					Postingan:   make(map[string][]string),
					NilaiKuis:   make(map[string]int),
					Tugas:       make(map[string]string),
					NilaiTugas:  make(map[string]int),
					NilaiUTS:    rand.Intn(41) + 60, // Nilai UTS random (60-100)
				},
				Penilaian:  make(map[string]map[string]int),
				NilaiTugas: make(map[string]map[string]int),
			}
		} else {
			// Jika peran bukan guru, membuat objek Pengguna biasa
			lms.Pengguna[username] = &Pengguna{
				Username:    username,
				Password:    password,
				Peran:       peran,
				JawabanKuis: make(map[string][]string),
				Postingan:   make(map[string][]string),
				NilaiKuis:   make(map[string]int),
				Tugas:       make(map[string]string),
				NilaiTugas:  make(map[string]int),
				NilaiUTS:    rand.Intn(41) + 60, // Nilai UTS random (60-100)
			}
		}
		fmt.Printf("Pengguna %s berhasil terdaftar!\n", username)
	} else {
		fmt.Println("Username sudah ada.") // Jika username sudah ada dalam sistem
	}
}

// Masuk ...
func (lms *LMS) Masuk(username, password string) interface{} {
	username = strings.ReplaceAll(username, " ", "") // Menghapus spasi dari username
	fmt.Println("Memeriksa kredensial...")

	// Memeriksa apakah pengguna dengan username yang diberikan ada dalam sistem
	if pengguna, ok := lms.Pengguna[username]; ok {
		switch p := pengguna.(type) {
		case *Pengguna:
			// Jika pengguna adalah siswa, memeriksa password dan perannya
			if p.Password == password && p.Peran == "siswa" {
				fmt.Printf("Selamat datang %s!\n", username)
				return p // Mengembalikan objek Pengguna jika kredensial valid
			}
		case *Guru:
			// Jika pengguna adalah guru, memeriksa password dan perannya
			if p.Password == password && p.Peran == "guru" {
				fmt.Printf("Selamat datang %s!\n", username)
				return p // Mengembalikan objek Guru jika kredensial valid
			}
		}
		fmt.Println("Password salah atau peran tidak sesuai.")
		return nil // Mengembalikan nilai nil jika kredensial tidak valid
	}
	fmt.Println("Pengguna tidak ditemukan. Silakan daftar terlebih dahulu.")
	return nil // Mengembalikan nilai nil jika pengguna tidak ditemukan
}

// MenuGuru ...
func (lms *LMS) MenuGuru(guru *Guru) {
	menuItems := []string{
		"Buat Materi",
		"Edit Materi",
		"Hapus Materi",
		"Buat Kuis",
		"Edit Kuis",
		"Hapus Kuis",
		"Buat Tugas",
		"Edit Tugas",
		"Hapus Tugas",
		"Buat Forum",
		"Data Siswa",
		"Keluar",
	}

	for {
		fmt.Println("==========================================")
		printMenu(menuItems) // Memanggil fungsi printMenu untuk menampilkan menu
		var pilihan int
		fmt.Scanln(&pilihan) // Meminta pengguna untuk memasukkan pilihan

		switch pilihan {
		// Memanggil metode yang sesuai berdasarkan pilihan pengguna
		case 1:
			lms.BuatMateri(guru.Pengguna)
		case 2:
			lms.EditMateri(guru.Pengguna)
		case 3:
			lms.HapusMateri(guru.Pengguna)
		case 4:
			lms.BuatKuis(guru.Pengguna)
		case 5:
			lms.EditKuis(guru.Pengguna)
		case 6:
			lms.HapusKuis(guru.Pengguna)
		case 7:
			lms.BuatTugas(guru.Pengguna)
		case 8:
			lms.EditTugas(guru.Pengguna)
		case 9:
			lms.HapusTugas(guru.Pengguna)
		case 10:
			lms.BuatForum(guru.Pengguna)
		case 11:
			lms.DataSiswa(guru)
		case 12:
			fmt.Println("Keluar...")
			fmt.Println("==========================================")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// MenuSiswa ...
func (lms *LMS) MenuSiswa(pengguna *Pengguna) {
	menuItems := []string{
		"Materi",
		"Kuis",
		"Tugas",
		"Lihat Forum",
		"Ikut Forum",
		"Lihat Nilai Kuis",
		"Lihat Nilai Tugas",
		"Lihat Nilai UTS",
		"Keluar",
	}

	for {
		fmt.Println("==========================================")
		printMenu(menuItems) // Menampilkan menu menggunakan fungsi printMenu
		var pilihan string   // Menggunakan string untuk menerima input
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1": // Materi
			lms.LihatMateri(pengguna)
		case "2": // Kuis
			lms.LihatKuis(pengguna)
			fmt.Print("Pilih judul kuis yang ingin Anda kerjakan: ")
			var judulKuis string
			fmt.Scanln(&judulKuis)
			lms.KerjakanKuis(pengguna, judulKuis) // Memanggil KerjakanKuis
		case "3": // Tugas
			lms.LihatTugas(pengguna)
			fmt.Print("Pilih judul tugas yang ingin Anda serahkan: ")
			var judulTugas string
			fmt.Scanln(&judulTugas)
			lms.SerahkanTugas(pengguna, judulTugas) // Memanggil SerahkanTugas
		case "4": // Lihat Forum
			lms.LihatForum(pengguna)
		case "5": // Ikut Forum
			lms.IkutForum(pengguna)
		case "6": // Lihat Nilai Kuis
			lms.LihatNilaiKuis(pengguna)
		case "7": // Lihat Nilai Tugas
			lms.LihatNilaiTugas(pengguna)
		case "8": // Lihat Nilai UTS
			fmt.Printf("Nilai UTS Anda: %d\n", pengguna.NilaiUTS)
		case "9": // Keluar
			fmt.Println("Keluar...")
			fmt.Println("==========================================")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// BuatMateri ...
func (lms *LMS) BuatMateri(pengguna *Pengguna) {
	var judul, konten string
	fmt.Print("Masukkan judul materi: ") // Meminta pengguna untuk memasukkan judul materi
	fmt.Scanln(&judul)
	fmt.Print("Masukkan konten materi: ") // Meminta pengguna untuk memasukkan konten materi
	fmt.Scanln(&konten)

	lms.Materi[judul] = konten             // Menambahkan materi baru ke map Materi dengan judul sebagai kunci dan konten sebagai nilai
	fmt.Println("Materi berhasil dibuat!") // Menampilkan pesan bahwa materi berhasil dibuat
}

// EditMateri ...
func (lms *LMS) EditMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi yang ingin Anda edit: ") // Meminta pengguna untuk memasukkan judul materi yang ingin diedit
	var judul string
	fmt.Scanln(&judul)

	// Mencari materi dengan judul (menggunakan pencarian sequential)
	if index := sequentialSearch(lms.getMateriJudul(), judul); index != -1 { // Memeriksa apakah materi dengan judul tersebut ada
		fmt.Print("Masukkan konten baru: ") // Meminta pengguna untuk memasukkan konten baru
		var kontenBaru string
		fmt.Scanln(&kontenBaru)

		lms.Materi[judul] = kontenBaru           // Mengupdate konten materi dengan konten baru
		fmt.Println("Materi berhasil diupdate!") // Menampilkan pesan bahwa materi berhasil diupdate
		return
	}

	fmt.Println("Materi tidak ditemukan.") // Menampilkan pesan bahwa materi tidak ditemukan jika judul tidak ada
}

// HapusMateri ...
func (lms *LMS) HapusMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi yang ingin Anda hapus: ") // Meminta pengguna untuk memasukkan judul materi yang ingin dihapus
	var judul string
	fmt.Scanln(&judul)

	// Mencari materi dengan judul (menggunakan pencarian sequential)
	if index := sequentialSearch(lms.getMateriJudul(), judul); index != -1 { // Memeriksa apakah materi dengan judul tersebut ada
		delete(lms.Materi, judul)               // Menghapus materi dengan judul tersebut
		fmt.Println("Materi berhasil dihapus!") // Menampilkan pesan bahwa materi berhasil dihapus
		return
	}

	fmt.Println("Materi tidak ditemukan.") // Menampilkan pesan bahwa materi tidak ditemukan jika judul tidak ada
}

// BuatKuis ...
func (lms *LMS) BuatKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis: ") // Meminta pengguna untuk memasukkan judul kuis
	var judul string
	fmt.Scanln(&judul)

	var pertanyaan []string
	for {
		fmt.Print("Masukkan pertanyaan (atau ketik 'selesai' untuk selesai): ") // Meminta pengguna untuk memasukkan pertanyaan
		pertanyaanStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		pertanyaanStr = strings.TrimSpace(pertanyaanStr)

		if pertanyaanStr == "selesai" { // Jika pengguna mengetik 'selesai', keluar dari loop
			break
		}

		pertanyaan = append(pertanyaan, pertanyaanStr) // Menambahkan pertanyaan ke slice pertanyaan
	}

	lms.Kuis[judul] = pertanyaan         // Menyimpan kuis beserta pertanyaannya
	fmt.Println("Kuis berhasil dibuat!") // Menampilkan pesan bahwa kuis berhasil dibuat
}

// EditKuis ...
func (lms *LMS) EditKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis yang ingin Anda edit: ") // Meminta pengguna untuk memasukkan judul kuis yang ingin diubah
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Kuis[judul]; ok { // Memeriksa apakah judul kuis tersebut ada di dalam map Kuis
		var pertanyaan []string
		for {
			fmt.Print("Masukkan pertanyaan (atau ketik 'selesai' untuk selesai): ") // Meminta pengguna untuk memasukkan pertanyaan baru
			pertanyaanStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			pertanyaanStr = strings.TrimSpace(pertanyaanStr)

			if pertanyaanStr == "selesai" { // Jika pengguna mengetik 'selesai', keluar dari loop
				break
			}

			pertanyaan = append(pertanyaan, pertanyaanStr) // Menambahkan pertanyaan baru ke slice pertanyaan
		}
		lms.Kuis[judul] = pertanyaan           // Mengupdate pertanyaan kuis
		fmt.Println("Kuis berhasil diupdate!") // Menampilkan pesan bahwa kuis berhasil diupdate
		return
	}

	fmt.Println("Kuis tidak ditemukan.") // Menampilkan pesan bahwa kuis tidak ditemukan
}

// HapusKuis ...
func (lms *LMS) HapusKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis yang ingin Anda hapus: ") // Meminta pengguna untuk memasukkan judul kuis yang ingin dihapus
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Kuis[judul]; ok { // Memeriksa apakah judul kuis tersebut ada di dalam map Kuis
		delete(lms.Kuis, judul)               // Menghapus kuis dengan judul yang dimasukkan pengguna
		fmt.Println("Kuis berhasil dihapus!") // Menampilkan pesan bahwa kuis berhasil dihapus
		return
	}

	fmt.Println("Kuis tidak ditemukan.") // Menampilkan pesan bahwa kuis tidak ditemukan
}

// BuatTugas ...
func (lms *LMS) BuatTugas(pengguna *Pengguna) {
	var judul, konten string
	fmt.Print("Masukkan judul tugas: ") // Meminta pengguna untuk memasukkan judul tugas baru
	fmt.Scanln(&judul)
	fmt.Print("Masukkan konten tugas: ") // Meminta pengguna untuk memasukkan konten tugas
	fmt.Scanln(&konten)

	lms.Tugas[judul] = konten             // Menambahkan tugas baru ke dalam map Tugas dengan judul sebagai kunci dan konten sebagai nilai
	fmt.Println("Tugas berhasil dibuat!") // Menampilkan pesan bahwa tugas berhasil dibuat
}

// EditTugas ...
func (lms *LMS) EditTugas(pengguna *Pengguna) {
	fmt.Print("Masukkan judul tugas yang ingin Anda edit: ") // Meminta pengguna untuk memasukkan judul tugas yang ingin diedit
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Tugas[judul]; ok { // Memeriksa apakah tugas dengan judul yang dimasukkan pengguna ada di map Tugas
		fmt.Print("Masukkan konten baru: ") // Meminta pengguna untuk memasukkan konten tugas yang baru
		var kontenBaru string
		fmt.Scanln(&kontenBaru)

		lms.Tugas[judul] = kontenBaru           // Mengupdate konten tugas dengan konten baru
		fmt.Println("Tugas berhasil diupdate!") // Menampilkan pesan bahwa tugas berhasil diupdate
		return
	}

	fmt.Println("Tugas tidak ditemukan.") // Menampilkan pesan bahwa tugas tidak ditemukan jika judul yang dimasukkan tidak ada dalam map Tugas
}

// HapusTugas ...
func (lms *LMS) HapusTugas(pengguna *Pengguna) {
	fmt.Print("Masukkan judul tugas yang ingin Anda hapus: ") // Meminta pengguna untuk memasukkan judul tugas yang ingin dihapus
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Tugas[judul]; ok { // Memeriksa apakah tugas dengan judul yang dimasukkan pengguna ada di map Tugas
		delete(lms.Tugas, judul)               // Menghapus tugas dengan judul yang dimasukkan
		fmt.Println("Tugas berhasil dihapus!") // Menampilkan pesan bahwa tugas berhasil dihapus
		return
	}

	fmt.Println("Tugas tidak ditemukan.") // Menampilkan pesan bahwa tugas tidak ditemukan jika judul yang dimasukkan tidak ada dalam map Tugas
}

// BuatForum ...
func (lms *LMS) BuatForum(pengguna *Pengguna) {
	var judul, postingan string
	fmt.Print("Masukkan judul forum: ") // Meminta pengguna untuk memasukkan judul forum yang ingin dibuat
	fmt.Scanln(&judul)
	fmt.Print("Masukkan postingan pertama: ")                 // Meminta pengguna untuk memasukkan postingan pertama dalam forum
	postingan, _ = bufio.NewReader(os.Stdin).ReadString('\n') // Membaca inputan pengguna termasuk spasi dan baris baru
	postingan = strings.TrimSpace(postingan)                  // Menghapus spasi di awal dan akhir postingan

	lms.Forum[judul] = []string{postingan} // Menyimpan postingan pertama dalam map Forum dengan judul forum sebagai kunci
	fmt.Println("Forum berhasil dibuat!")  // Menampilkan pesan bahwa forum berhasil dibuat
}

// LihatForum ...
func (lms *LMS) LihatForum(pengguna *Pengguna) {
	header := []string{"Judul Forum"} // Membuat header tabel untuk menampilkan judul forum
	var data [][]string               // Variabel untuk menyimpan data forum dalam bentuk slice

	// Mengambil setiap judul forum dan memasukkannya ke dalam data
	for judul := range lms.Forum {
		data = append(data, []string{judul})
	}

	// Menampilkan tabel dengan judul forum yang tersedia
	printTable(header, data)

	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ") // Meminta pengguna untuk memasukkan judul forum yang ingin diikuti
	var judul string
	fmt.Scanln(&judul)

	// Memeriksa apakah judul forum yang dimasukkan oleh pengguna ada dalam daftar forum
	if postingan, ok := lms.Forum[judul]; ok {
		fmt.Println("Postingan di forum ini:") // Menampilkan judul forum yang dipilih dan postingannya
		for _, p := range postingan {
			fmt.Println(p)
		}

		fmt.Print("Masukkan postingan Anda: ")                         // Meminta pengguna untuk memasukkan postingan baru dalam forum
		postinganBaru, _ := bufio.NewReader(os.Stdin).ReadString('\n') // Membaca inputan pengguna termasuk spasi dan baris baru
		postinganBaru = strings.TrimSpace(postinganBaru)               // Menghapus spasi di awal dan akhir postingan baru

		// Menambahkan postingan baru ke dalam forum dengan judul yang dipilih
		lms.Forum[judul] = append(lms.Forum[judul], postinganBaru)
		fmt.Println("Postingan Anda berhasil ditambahkan!") // Menampilkan pesan bahwa postingan berhasil ditambahkan
		return
	}

	fmt.Println("Forum tidak ditemukan.") // Menampilkan pesan bahwa forum tidak ditemukan jika judul forum tidak ada dalam daftar forum
}

// DataSiswa ...
func (lms *LMS) DataSiswa(guru *Guru) {
	fmt.Println("\nData Siswa:") // Menampilkan header untuk data siswa

	// Mengurutkan username siswa secara alfabetis
	var usernames []string
	for username := range lms.Pengguna {
		if _, ok := lms.Pengguna[username].(*Pengguna); ok {
			usernames = append(usernames, username)
		}
	}
	sort.Strings(usernames)

	headerSiswa := []string{"Username", "Nilai UTS"} // Header untuk tabel data siswa
	var dataSiswa [][]string                         // Variabel untuk menyimpan data siswa dalam bentuk slice
	for _, username := range usernames {             // Mengambil setiap username siswa dan nilai UTS-nya
		if siswa, ok := lms.Pengguna[username].(*Pengguna); ok {
			dataSiswa = append(dataSiswa, []string{username, strconv.Itoa(siswa.NilaiUTS)})
		}
	}
	printTable(headerSiswa, dataSiswa) // Menampilkan tabel data siswa

	fmt.Print("Masukkan username siswa untuk melihat detail atau memberi nilai: ") // Meminta inputan pengguna untuk memilih username siswa
	var usernameSiswa string
	fmt.Scanln(&usernameSiswa)

	if siswa, ok := lms.Pengguna[usernameSiswa]; ok { // Memeriksa apakah username siswa yang dimasukkan ada dalam daftar pengguna
		if _, ok := siswa.(*Pengguna); ok { // Memeriksa apakah pengguna dengan username tersebut adalah seorang siswa
			lms.tampilkanDataSiswa(guru, usernameSiswa) // Menampilkan detail atau memberi nilai kepada siswa yang dipilih
		} else {
			fmt.Println("Pengguna tersebut bukan siswa.") // Menampilkan pesan jika pengguna bukanlah seorang siswa
		}
	} else {
		fmt.Println("Siswa tidak ditemukan.") // Menampilkan pesan jika siswa tidak ditemukan dalam daftar pengguna
	}
}

// tampilkanDataSiswa ...
func (lms *LMS) tampilkanDataSiswa(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna) // Mengambil data siswa dari pengguna dengan username yang diberikan

	fmt.Println("\nDetail Siswa:")               // Menampilkan header untuk detail siswa
	fmt.Printf("Username: %s\n", siswa.Username) // Menampilkan username siswa

	fmt.Println("\nJawaban Kuis:")  // Menampilkan jawaban kuis siswa
	if len(siswa.JawabanKuis) > 0 { // Memeriksa apakah siswa memiliki jawaban kuis
		for judulKuis, jawaban := range siswa.JawabanKuis { // Mengambil setiap jawaban kuis siswa
			fmt.Printf("- %s:\n", judulKuis) // Menampilkan judul kuis
			for i, jwb := range jawaban {    // Menampilkan setiap jawaban kuis
				fmt.Printf("  %d. %s\n", i+1, jwb)
			}
		}
	} else {
		fmt.Println("  (Tidak ada jawaban kuis)") // Menampilkan pesan jika siswa tidak memiliki jawaban kuis
	}

	fmt.Println("\nNilai Kuis:")                                 // Menampilkan nilai kuis siswa
	if penilaianSiswa, ok := guru.Penilaian[usernameSiswa]; ok { // Memeriksa apakah guru memberikan penilaian kuis kepada siswa
		headerNilai := []string{"Kuis", "Nilai"}  // Header untuk tabel nilai kuis
		var dataNilai [][]string                  // Variabel untuk menyimpan data nilai kuis siswa dalam bentuk slice
		for kuis, nilai := range penilaianSiswa { // Mengambil setiap nilai kuis siswa
			dataNilai = append(dataNilai, []string{kuis, strconv.Itoa(nilai)})
		}
		printTable(headerNilai, dataNilai) // Menampilkan tabel nilai kuis siswa
	} else {
		fmt.Println("  (Belum ada nilai)") // Menampilkan pesan jika belum ada nilai kuis
	}

	fmt.Println("\nJawaban Tugas:") // Menampilkan jawaban tugas siswa
	if len(siswa.Tugas) > 0 {       // Memeriksa apakah siswa memiliki jawaban tugas
		for judulTugas, jawaban := range siswa.Tugas { // Mengambil setiap jawaban tugas siswa
			fmt.Printf("- %s:\n", judulTugas) // Menampilkan judul tugas
			fmt.Printf("  %s\n", jawaban)     // Menampilkan jawaban tugas
		}
	} else {
		fmt.Println("  (Tidak ada jawaban tugas)") // Menampilkan pesan jika siswa tidak memiliki jawaban tugas
	}

	fmt.Println("\nNilai Tugas:")                                 // Menampilkan nilai tugas siswa
	if penilaianSiswa, ok := guru.NilaiTugas[usernameSiswa]; ok { // Memeriksa apakah guru memberikan penilaian tugas kepada siswa
		headerNilai := []string{"Tugas", "Nilai"}  // Header untuk tabel nilai tugas
		var dataNilai [][]string                   // Variabel untuk menyimpan data nilai tugas siswa dalam bentuk slice
		for tugas, nilai := range penilaianSiswa { // Mengambil setiap nilai tugas siswa
			dataNilai = append(dataNilai, []string{tugas, strconv.Itoa(nilai)})
		}
		printTable(headerNilai, dataNilai) // Menampilkan tabel nilai tugas siswa
	} else {
		fmt.Println("  (Belum ada nilai)") // Menampilkan pesan jika belum ada nilai tugas
	}

	fmt.Print("\nApakah Anda ingin memberikan penilaian untuk kuis? (y/n): ") // Meminta inputan pengguna untuk memberikan penilaian kuis
	var beriPenilaian string
	fmt.Scanln(&beriPenilaian)

	if strings.ToLower(beriPenilaian) == "y" { // Jika pengguna ingin memberikan penilaian kuis
		lms.berikanPenilaian(guru, usernameSiswa) // Memanggil fungsi untuk memberikan penilaian kuis
	}

	fmt.Print("\nApakah Anda ingin memberikan penilaian untuk tugas? (y/n): ") // Meminta inputan pengguna untuk memberikan penilaian tugas
	var beriPenilaianTugas string
	fmt.Scanln(&beriPenilaianTugas)

	if strings.ToLower(beriPenilaianTugas) == "y" { // Jika pengguna ingin memberikan penilaian tugas
		lms.berikanPenilaianTugas(guru, usernameSiswa) // Memanggil fungsi untuk memberikan penilaian tugas
	}
}

// berikanPenilaianTugas ...
func (lms *LMS) berikanPenilaianTugas(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna) // Mengambil data siswa dari pengguna dengan username yang diberikan

	fmt.Println("\nDaftar Tugas yang Telah Dikerjakan:") // Menampilkan daftar tugas yang telah dikerjakan oleh siswa
	i := 1
	for judulTugas := range siswa.Tugas {
		fmt.Printf("%d. %s\n", i, judulTugas)
		i++
	}

	fmt.Print("Pilih nomor tugas untuk diberi nilai: ") // Meminta inputan nomor tugas yang akan diberi nilai
	var pilihanTugas int
	fmt.Scanln(&pilihanTugas)

	i = 1
	var judulTugasTerpilih string
	for judulTugas := range siswa.Tugas { // Mengambil judul tugas yang sesuai dengan nomor pilihan
		if i == pilihanTugas {
			judulTugasTerpilih = judulTugas
			break
		}
		i++
	}

	if judulTugasTerpilih != "" { // Memeriksa apakah judul tugas terpilih tidak kosong
		fmt.Print("Masukkan nilai untuk tugas '" + judulTugasTerpilih + "': ") // Meminta inputan nilai untuk tugas terpilih
		var nilai int
		fmt.Scanln(&nilai)

		if guru.NilaiTugas[usernameSiswa] == nil { // Memeriksa apakah nilai tugas siswa sudah ada
			guru.NilaiTugas[usernameSiswa] = make(map[string]int) // Jika belum ada, inisialisasi map nilai tugas siswa
		}
		guru.NilaiTugas[usernameSiswa][judulTugasTerpilih] = nilai // Menyimpan nilai tugas siswa pada map nilai tugas guru
		fmt.Println("Penilaian berhasil diberikan!")               // Menampilkan pesan penilaian berhasil diberikan
	} else {
		fmt.Println("Pilihan tugas tidak valid.") // Menampilkan pesan jika pilihan tugas tidak valid
	}
}

// berikanPenilaian ...
func (lms *LMS) berikanPenilaian(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna) // Mengambil data siswa dari pengguna dengan username yang diberikan

	fmt.Println("\nDaftar Kuis yang Telah Dikerjakan:") // Menampilkan daftar kuis yang telah dikerjakan oleh siswa
	i := 1
	for judulKuis := range siswa.JawabanKuis {
		fmt.Printf("%d. %s\n", i, judulKuis)
		i++
	}

	fmt.Print("Pilih nomor kuis untuk diberi nilai: ") // Meminta inputan nomor kuis yang akan diberi nilai
	var pilihanKuis int
	fmt.Scanln(&pilihanKuis)

	i = 1
	var judulKuisTerpilih string
	for judulKuis := range siswa.JawabanKuis { // Mengambil judul kuis yang sesuai dengan nomor pilihan
		if i == pilihanKuis {
			judulKuisTerpilih = judulKuis
			break
		}
		i++
	}

	if judulKuisTerpilih != "" { // Memeriksa apakah judul kuis terpilih tidak kosong
		fmt.Print("Masukkan nilai untuk kuis '" + judulKuisTerpilih + "': ") // Meminta inputan nilai untuk kuis terpilih
		var nilai int
		fmt.Scanln(&nilai)

		if guru.Penilaian[usernameSiswa] == nil { // Memeriksa apakah nilai kuis siswa sudah ada
			guru.Penilaian[usernameSiswa] = make(map[string]int) // Jika belum ada, inisialisasi map nilai kuis guru
		}
		guru.Penilaian[usernameSiswa][judulKuisTerpilih] = nilai // Menyimpan nilai kuis siswa pada map nilai kuis guru
		fmt.Println("Penilaian berhasil diberikan!")             // Menampilkan pesan penilaian berhasil diberikan
	} else {
		fmt.Println("Pilihan kuis tidak valid.") // Menampilkan pesan jika pilihan kuis tidak valid
	}
}

// LihatMateri ...
func (lms *LMS) LihatMateri(pengguna *Pengguna) {
	header := []string{"Judul Materi", "Konten"} // Judul kolom pada tabel
	var data [][]string

	for judul, konten := range lms.Materi { // Mengisi data materi ke dalam slice data
		data = append(data, []string{judul, konten})
	}

	printTable(header, data) // Mencetak tabel dengan judul dan konten materi
}

// LihatKuis ...
func (lms *LMS) LihatKuis(pengguna *Pengguna) {
	header := []string{"Judul Kuis"} // Judul kolom pada tabel
	var data [][]string

	for judul := range lms.Kuis { // Mengisi data judul kuis ke dalam slice data
		data = append(data, []string{judul})
	}

	printTable(header, data) // Mencetak tabel dengan judul kuis
}

// LihatTugas ...
func (lms *LMS) LihatTugas(pengguna *Pengguna) {
	header := []string{"Judul Tugas"} // Judul kolom pada tabel
	var data [][]string

	for judul := range lms.Tugas { // Mengisi data judul tugas ke dalam slice data
		data = append(data, []string{judul})
	}

	printTable(header, data) // Mencetak tabel dengan judul tugas
}

// KerjakanKuis ...
func (lms *LMS) KerjakanKuis(pengguna *Pengguna, judulKuis string) {
	if pertanyaan, ok := lms.Kuis[judulKuis]; ok { // Mengecek apakah judul kuis ditemukan
		var jawaban []string
		for _, p := range pertanyaan { // Looping untuk setiap pertanyaan pada kuis
			fmt.Println(p) // Cetak pertanyaan
			var jawabanStr string
			fmt.Scanln(&jawabanStr)               // Menerima input jawaban dari pengguna
			jawaban = append(jawaban, jawabanStr) // Menambahkan jawaban ke slice jawaban
		}

		pengguna.JawabanKuis[judulKuis] = jawaban // Simpan jawaban siswa dalam struktur Pengguna
		fmt.Println("Jawaban kuis berhasil disimpan!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// SerahkanTugas ...
func (lms *LMS) SerahkanTugas(pengguna *Pengguna, judulTugas string) {
	if _, ok := lms.Tugas[judulTugas]; ok { // Mengecek apakah judul tugas ditemukan
		fmt.Print("Masukkan jawaban tugas Anda: ")
		var jawabanTugas string
		fmt.Scanln(&jawabanTugas) // Menerima input jawaban tugas dari pengguna

		pengguna.Tugas[judulTugas] = jawabanTugas // Simpan jawaban tugas dalam struktur Pengguna
		fmt.Println("Jawaban tugas berhasil disimpan!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

// IkutForum ...
func (lms *LMS) IkutForum(pengguna *Pengguna) {
	lms.LihatForum(pengguna) // Menampilkan judul forum yang tersedia

	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ")
	var judul string
	fmt.Scanln(&judul) // Menerima input judul forum dari pengguna

	if postingan, ok := lms.Forum[judul]; ok { // Mengecek apakah judul forum ditemukan
		fmt.Println("Postingan di forum ini:")
		for _, p := range postingan {
			fmt.Println(p) // Menampilkan postingan di forum yang dipilih
		}

		fmt.Print("Masukkan postingan Anda: ")
		postinganBaru, _ := bufio.NewReader(os.Stdin).ReadString('\n') // Menerima input postingan baru dari pengguna
		postinganBaru = strings.TrimSpace(postinganBaru)

		lms.Forum[judul] = append(lms.Forum[judul], postinganBaru) // Menambahkan postingan baru ke dalam forum
		fmt.Println("Postingan Anda berhasil ditambahkan!")
		return
	}

	fmt.Println("Forum tidak ditemukan.")
}

// LihatNilaiKuis ...
func (lms *LMS) LihatNilaiKuis(pengguna *Pengguna) {
	for _, p := range lms.Pengguna {
		if guru, ok := p.(*Guru); ok { // Memeriksa apakah pengguna adalah guru
			if penilaianSiswa, ok := guru.Penilaian[pengguna.Username]; ok { // Memeriksa apakah ada penilaian siswa oleh guru
				header := []string{"Judul Kuis", "Nilai"}
				var data [][]string
				for judulKuis, nilai := range penilaianSiswa {
					data = append(data, []string{judulKuis, strconv.Itoa(nilai)}) // Menyiapkan data penilaian kuis siswa
				}
				printTable(header, data) // Menampilkan tabel penilaian kuis siswa
				return
			}
		}
	}
	fmt.Println("Nilai kuis belum tersedia.")
}

// LihatNilaiTugas ...
func (lms *LMS) LihatNilaiTugas(pengguna *Pengguna) {
	for _, p := range lms.Pengguna {
		if guru, ok := p.(*Guru); ok { // Memeriksa apakah pengguna adalah guru
			if penilaianSiswa, ok := guru.NilaiTugas[pengguna.Username]; ok { // Memeriksa apakah ada penilaian tugas siswa oleh guru
				header := []string{"Judul Tugas", "Nilai"}
				var data [][]string
				for judulTugas, nilai := range penilaianSiswa {
					data = append(data, []string{judulTugas, strconv.Itoa(nilai)}) // Menyiapkan data penilaian tugas siswa
				}
				printTable(header, data) // Menampilkan tabel penilaian tugas siswa
				return
			}
		}
	}
	fmt.Println("Nilai tugas belum tersedia.")
}

// sequentialSearch ...
func sequentialSearch(data []string, target string) int {
	for i, item := range data {
		if item == target { // Mengecek apakah item pada indeks ke-i sama dengan target
			return i // Mengembalikan indeks jika ditemukan
		}
	}
	return -1 // Mengembalikan -1 jika tidak ditemukan
}

// binarySearch ...
func binarySearch(data []string, target string) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2  // Menentukan indeks tengah dari rentang pencarian
		if data[mid] == target { // Jika nilai tengah sama dengan target, mengembalikan indeks tengah
			return mid
		} else if data[mid] < target { // Jika nilai tengah kurang dari target, perbesar rentang ke kanan
			low = mid + 1
		} else { // Jika nilai tengah lebih dari target, perkecil rentang ke kiri
			high = mid - 1
		}
	}
	return -1 // Mengembalikan -1 jika target tidak ditemukan dalam data
}

// selectionSort ...
func selectionSort(data []string, ascending bool) {
	n := len(data)
	for i := 0; i < n-1; i++ { // Looping sebanyak n-1 kali
		minIndex := i                // Menandai indeks minimum pada awal iterasi
		for j := i + 1; j < n; j++ { // Looping untuk mencari nilai minimum
			if ascending { // Jika ascending, cari nilai minimum
				if data[j] < data[minIndex] {
					minIndex = j
				}
			} else { // Jika descending, cari nilai maksimum
				if data[j] > data[minIndex] {
					minIndex = j
				}
			}
		}
		// Tukar nilai minimum dengan nilai pada indeks i
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}

// insertionSort ...
func insertionSort(data []string, ascending bool) {
	n := len(data)
	for i := 1; i < n; i++ { // Looping dari elemen kedua hingga elemen terakhir
		key := data[i] // Simpan nilai elemen yang akan disisipkan
		j := i - 1     // Mulai pengecekan dari indeks sebelumnya

		// Geser elemen-elemen yang lebih besar (untuk ascending) atau lebih kecil (untuk descending) dari key
		for j >= 0 && (ascending && data[j] > key || !ascending && data[j] < key) {
			data[j+1] = data[j] // Geser elemen ke kanan
			j--                 // Pindah ke elemen sebelumnya
		}
		data[j+1] = key // Tempatkan key pada posisi yang tepat
	}
}

// Get a slice of material titles (keys from the Materi map)
func (lms *LMS) getMateriJudul() []string {
	var titles []string
	for title := range lms.Materi { // Melakukan iterasi atas judul-judul dari map Materi
		titles = append(titles, title) // Menambahkan setiap judul ke dalam slice titles
	}
	return titles // Mengembalikan slice yang berisi semua judul Materi
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inisialisasi generator nomor acak

	// Membuat instansi baru dari LMS
	lms := LMS{
		Pengguna: make(map[string]interface{}), // Inisialisasi map untuk pengguna
		Materi:   make(map[string]string),      // Inisialisasi map untuk materi
		Kuis:     make(map[string][]string),    // Inisialisasi map untuk kuis
		Forum:    make(map[string][]string),    // Inisialisasi map untuk forum
		Tugas:    make(map[string]string),      // Inisialisasi map untuk tugas
	}

	// Menambahkan data awal untuk guru
	lms.Daftar("pakguru", "1234", "guru")

	// Data awal siswa
	daftarSiswa := []string{
		"Arimbi Try Wardani",
		"Kisya Reinatana Yama",
		"Stiefanny Dwi Chandra",
		"M Rafadi Kurniawan",
		"Defari Akbar Anggara",
		"Muhamad Audi Radittia Prasetyo",
		"Patrick Nicholas",
		"Muhammad Farid Irham",
		"Aresky Brilyan",
		"Sombolinggi",
		"Elsa Melisa Silaen",
		"Aflah Zaki Siregar",
		"Gustira Haryani",
		"Satrio Aji Nugroho",
		"Metha Anastasya",
		"Rizza Hafizh Al-Iqhwal",
		"M. Rizky Hadi Prawiro",
		"Irfan Nafis Maulana",
		"Fachrul Rozi",
		"Andi Naufal Dwi Putra",
		"Abiyyu Aryasena",
		"Aditya ataby",
		"adha rahmadani",
		"hasan naqib sa'bani",
		"M. Akbar rafsanjani",
		"Nomensen Melkisedek Pardosi",
		"Aldyansyah Wisnu Saputra",
		"Bintang Darma Sakti",
		"Azka Dhaffinanda Rahman",
		"Muhammad Hario Ifanny El Jr gania",
		"luthfi mawlanza gania",
		"Revaldi praditya gania",
		"Faiz satrio gania",
		"Fazli Radika",
		"Muhammad Afriza Hidayat",
		"Naufal Saifullah Yusuf",
		"Raehan Aldiansyah",
	}

	// Mendaftarkan semua siswa
	for i, siswa := range daftarSiswa {
		lms.Daftar(strings.ReplaceAll(siswa, " ", ""), fmt.Sprintf("password%d", i+1), "siswa")
	}

	// Data awal materi
	lms.Materi["Pembelajaran Golang"] = "Golang, atau Go, adalah bahasa pemrograman yang dikembangkan oleh Google. Bahasa ini dirancang oleh Robert Griesemer, Rob Pike, dan Ken Thompson dan pertama kali dirilis pada tahun 2009."
	lms.Materi["Java"] = "Java adalah bahasa pemrograman yang dikembangkan oleh Sun Microsystems, yang sekarang menjadi bagian dari Oracle.  Java dikenal karena portabilitasnya, yang memungkinkan program Java untuk dijalankan di berbagai platform.  Java sangat populer untuk pengembangan aplikasi enterprise, aplikasi Android, aplikasi desktop, dan sistem terdistribusi."
	lms.Materi["Python"] = "Python adalah bahasa pemrograman yang dirancang oleh Guido van Rossum dan pertama kali dirilis pada tahun 1991. Python adalah bahasa yang interpretatif, berorientasi objek, dan mudah dipelajari. Python sangat populer untuk pengembangan web, data science, machine learning, scripting, otomatisasi, dan banyak lagi."
	lms.Materi["JavaScript"] = "JavaScript adalah bahasa pemrograman yang dikembangkan oleh Brendan Eich untuk Netscape Communications.  JavaScript awalnya dirancang untuk membuat web menjadi lebih interaktif, dan sekarang merupakan bahasa yang sangat penting untuk pengembangan web, aplikasi mobile, dan aplikasi desktop.  JavaScript memiliki banyak framework populer, seperti React, Angular, dan Vue."
	lms.Materi["C++"] = "C++ adalah bahasa pemrograman yang dikembangkan oleh Bjarne Stroustrup.  C++ adalah bahasa pemrograman yang berorientasi objek, memiliki kontrol tingkat rendah, dan dikenal karena performanya yang tinggi.  C++ sangat populer untuk pengembangan perangkat lunak sistem, game, aplikasi yang membutuhkan performa tinggi, dan perangkat lunak embedded."

	// Data awal kuis
	lms.Kuis["golang"] = []string{
		"Siapa yang merancang golang?",
		"pada tahun berapa golang dirilis?",
		"Perusahaan apa yang mengembangkan golang?",
	}

	// Data awal tugas
	lms.Tugas["Tugas Golang"] = "Buatlah program Golang sederhana untuk menampilkan 'Hello, World!'"

	var pilihan int
	for {
		fmt.Println("==========================================")
		fmt.Println("         SELAMAT DATANG")
		fmt.Println("   LMS Telkom University Jakarta")
		fmt.Println("==========================================")
		fmt.Println("1. Guru")
		fmt.Println("2. Siswa")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan Anda: ")

		// Membaca input dari pengguna dan menangani error
		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Error membaca input:", err)
			continue
		}

		switch pilihan {
		case 1:
			fmt.Println("1. Daftar Guru")
			fmt.Println("2. Masuk Guru")
			fmt.Print("Pilihan Anda: ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case 1:
				var username, password string
				fmt.Print("Masukkan username: ")
				fmt.Scanln(&username)
				fmt.Print("Masukkan password: ")
				fmt.Scanln(&password)
				lms.Daftar(username, password, "guru")
			case 2:
				var username, password string
				fmt.Print("Masukkan username: ")
				fmt.Scanln(&username)
				fmt.Print("Masukkan password: ")
				fmt.Scanln(&password)
				pengguna := lms.Masuk(username, password)
				if pengguna != nil {
					if guru, ok := pengguna.(*Guru); ok {
						fmt.Println("==========================================")
						fmt.Println("         SELAMAT DATANG")
						fmt.Println("             GURU")
						fmt.Println("==========================================")
						lms.MenuGuru(guru)
					}
				}
			}
		case 2:
			fmt.Println("1. Daftar Siswa")
			fmt.Println("2. Masuk Siswa")
			fmt.Print("Pilihan Anda: ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case 1:
				var username, password string
				fmt.Print("Masukkan username: ")
				fmt.Scanln(&username)
				fmt.Print("Masukkan password: ")
				fmt.Scanln(&password)
				lms.Daftar(username, password, "siswa")

			case 2:
				var username, password string
				fmt.Print("Masukkan username: ")
				fmt.Scanln(&username)
				fmt.Print("Masukkan password: ")
				fmt.Scanln(&password)
				pengguna := lms.Masuk(username, password)
				if pengguna != nil {
					if siswa, ok := pengguna.(*Pengguna); ok {
						fmt.Println("==========================================")
						fmt.Println("         SELAMAT DATANG")
						fmt.Println("             SISWA")
						fmt.Println("==========================================")
						lms.MenuSiswa(siswa)
					}
				}
			}
		case 3:
			fmt.Println("==========================================")
			fmt.Println("           TERIMA KASIH")
			fmt.Println("    SAMPAI JUMPA DI KESEMPATAN LAIN")
			fmt.Println("==========================================")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
