package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Pengguna ...
type Pengguna struct {
	Username    string
	Password    string
	Peran       string
	JawabanKuis map[string][]string
	Postingan   map[string][]string
	NilaiKuis   map[string]int
	Tugas       map[string]string // Kunci: Judul Tugas, Nilai: Konten Tugas
	NilaiTugas  map[string]int    // Kunci: Judul Tugas, Nilai: Nilai
	NilaiUTS    int               // Nilai UTS
}

// Guru ...
type Guru struct {
	*Pengguna
	Penilaian  map[string]map[string]int
	NilaiTugas map[string]map[string]int // Kunci: Username Siswa, Nilai: map[Judul Tugas]Nilai
}

// LMS ...
type LMS struct {
	Pengguna map[string]interface{}
	Materi   map[string]string
	Kuis     map[string][]string
	Forum    map[string][]string
	Tugas    map[string]string // Kunci: Judul Tugas, Nilai: Konten Tugas
}

// printTable ...
func printTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
}

// printMenu ...
func printMenu(menuItems []string) {
	fmt.Println()
	for i, item := range menuItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Pilihan Anda: ")
}

// Daftar ...
func (lms *LMS) Daftar(username, password, peran string) {
	username = strings.ReplaceAll(username, " ", "")
	fmt.Println("Mendaftarkan pengguna...")

	if _, ok := lms.Pengguna[username]; !ok {
		if peran == "guru" {
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
		fmt.Println("Username sudah ada.")
	}
}

// Masuk ...
func (lms *LMS) Masuk(username, password string) interface{} {
	username = strings.ReplaceAll(username, " ", "")
	fmt.Println("Memeriksa kredensial...")

	if pengguna, ok := lms.Pengguna[username]; ok {
		switch p := pengguna.(type) {
		case *Pengguna:
			if p.Password == password && p.Peran == "siswa" {
				fmt.Printf("Selamat datang %s!\n", username)
				return p
			}
		case *Guru:
			if p.Password == password && p.Peran == "guru" {
				fmt.Printf("Selamat datang %s!\n", username)
				return p
			}
		}
		fmt.Println("Password salah atau peran tidak sesuai.")
		return nil
	}
	fmt.Println("Pengguna tidak ditemukan. Silakan daftar terlebih dahulu.")
	return nil
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
		printMenu(menuItems)
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
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
		printMenu(menuItems)
		var pilihan string // Ubah ke string untuk menerima input
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1": // Materi
			lms.LihatMateri(pengguna)
		case "2": // Kuis
			lms.LihatKuis(pengguna)
			fmt.Print("Pilih judul kuis yang ingin Anda kerjakan: ")
			var judulKuis string
			fmt.Scanln(&judulKuis)
			lms.KerjakanKuis(pengguna, judulKuis) // Panggil KerjakanKuis
		case "3": // Tugas
			lms.LihatTugas(pengguna)
			fmt.Print("Pilih judul tugas yang ingin Anda serahkan: ")
			var judulTugas string
			fmt.Scanln(&judulTugas)
			lms.SerahkanTugas(pengguna, judulTugas) // Panggil SerahkanTugas
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
	fmt.Print("Masukkan judul materi: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan konten materi: ")
	fmt.Scanln(&konten)

	lms.Materi[judul] = konten
	fmt.Println("Materi berhasil dibuat!")
}

// EditMateri ...
func (lms *LMS) EditMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	// Mencari materi dengan judul (menggunakan pencarian sequential)
	if index := sequentialSearch(lms.getMateriJudul(), judul); index != -1 {
		fmt.Print("Masukkan konten baru: ")
		var kontenBaru string
		fmt.Scanln(&kontenBaru)

		lms.Materi[judul] = kontenBaru
		fmt.Println("Materi berhasil diupdate!")
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

// HapusMateri ...
func (lms *LMS) HapusMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	// Mencari materi dengan judul (menggunakan pencarian sequential)
	if index := sequentialSearch(lms.getMateriJudul(), judul); index != -1 {
		delete(lms.Materi, judul)
		fmt.Println("Materi berhasil dihapus!")
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

// BuatKuis ...
func (lms *LMS) BuatKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis: ")
	var judul string
	fmt.Scanln(&judul)

	var pertanyaan []string
	for {
		fmt.Print("Masukkan pertanyaan (atau ketik 'selesai' untuk selesai): ")
		pertanyaanStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		pertanyaanStr = strings.TrimSpace(pertanyaanStr)

		if pertanyaanStr == "selesai" {
			break
		}

		pertanyaan = append(pertanyaan, pertanyaanStr)
	}

	lms.Kuis[judul] = pertanyaan
	fmt.Println("Kuis berhasil dibuat!")
}

// EditKuis ...
func (lms *LMS) EditKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Kuis[judul]; ok {
		var pertanyaan []string
		for {
			fmt.Print("Masukkan pertanyaan (atau ketik 'selesai' untuk selesai): ")
			pertanyaanStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			pertanyaanStr = strings.TrimSpace(pertanyaanStr)

			if pertanyaanStr == "selesai" {
				break
			}

			pertanyaan = append(pertanyaan, pertanyaanStr)
		}
		lms.Kuis[judul] = pertanyaan
		fmt.Println("Kuis berhasil diupdate!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// HapusKuis ...
func (lms *LMS) HapusKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Kuis[judul]; ok {
		delete(lms.Kuis, judul)
		fmt.Println("Kuis berhasil dihapus!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// BuatTugas ...
func (lms *LMS) BuatTugas(pengguna *Pengguna) {
	var judul, konten string
	fmt.Print("Masukkan judul tugas: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan konten tugas: ")
	fmt.Scanln(&konten)

	lms.Tugas[judul] = konten
	fmt.Println("Tugas berhasil dibuat!")
}

// EditTugas ...
func (lms *LMS) EditTugas(pengguna *Pengguna) {
	fmt.Print("Masukkan judul tugas yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Tugas[judul]; ok {
		fmt.Print("Masukkan konten baru: ")
		var kontenBaru string
		fmt.Scanln(&kontenBaru)

		lms.Tugas[judul] = kontenBaru
		fmt.Println("Tugas berhasil diupdate!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

// HapusTugas ...
func (lms *LMS) HapusTugas(pengguna *Pengguna) {
	fmt.Print("Masukkan judul tugas yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Tugas[judul]; ok {
		delete(lms.Tugas, judul)
		fmt.Println("Tugas berhasil dihapus!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

// BuatForum ...
func (lms *LMS) BuatForum(pengguna *Pengguna) {
	var judul, postingan string
	fmt.Print("Masukkan judul forum: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan postingan pertama: ")
	postingan, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	postingan = strings.TrimSpace(postingan)

	lms.Forum[judul] = []string{postingan}
	fmt.Println("Forum berhasil dibuat!")
}

// LihatForum ...
func (lms *LMS) LihatForum(pengguna *Pengguna) {
	header := []string{"Judul Forum"}
	var data [][]string

	for judul := range lms.Forum {
		data = append(data, []string{judul})
	}

	printTable(header, data)

	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ")
	var judul string
	fmt.Scanln(&judul)

	if postingan, ok := lms.Forum[judul]; ok {
		fmt.Println("Postingan di forum ini:")
		for _, p := range postingan {
			fmt.Println(p)
		}

		fmt.Print("Masukkan postingan Anda: ")
		postinganBaru, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		postinganBaru = strings.TrimSpace(postinganBaru)

		lms.Forum[judul] = append(lms.Forum[judul], postinganBaru)
		fmt.Println("Postingan Anda berhasil ditambahkan!")
		return
	}

	fmt.Println("Forum tidak ditemukan.")
}

// DataSiswa ...
func (lms *LMS) DataSiswa(guru *Guru) {
	fmt.Println("\nData Siswa:")

	// Mengurutkan username siswa
	var usernames []string
	for username := range lms.Pengguna {
		if _, ok := lms.Pengguna[username].(*Pengguna); ok {
			usernames = append(usernames, username)
		}
	}
	sort.Strings(usernames)

	headerSiswa := []string{"Username", "Nilai UTS"}
	var dataSiswa [][]string
	for _, username := range usernames {
		if siswa, ok := lms.Pengguna[username].(*Pengguna); ok {
			dataSiswa = append(dataSiswa, []string{username, strconv.Itoa(siswa.NilaiUTS)})
		}
	}
	printTable(headerSiswa, dataSiswa)

	fmt.Print("Masukkan username siswa untuk melihat detail atau memberi nilai: ")
	var usernameSiswa string
	fmt.Scanln(&usernameSiswa)

	if siswa, ok := lms.Pengguna[usernameSiswa]; ok {
		if _, ok := siswa.(*Pengguna); ok {
			lms.tampilkanDataSiswa(guru, usernameSiswa)
		} else {
			fmt.Println("Pengguna tersebut bukan siswa.")
		}
	} else {
		fmt.Println("Siswa tidak ditemukan.")
	}
}

// tampilkanDataSiswa ...
func (lms *LMS) tampilkanDataSiswa(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna)

	fmt.Println("\nDetail Siswa:")
	fmt.Printf("Username: %s\n", siswa.Username)

	fmt.Println("\nJawaban Kuis:")
	if len(siswa.JawabanKuis) > 0 {
		for judulKuis, jawaban := range siswa.JawabanKuis {
			fmt.Printf("- %s:\n", judulKuis)
			for i, jwb := range jawaban {
				fmt.Printf("  %d. %s\n", i+1, jwb)
			}
		}
	} else {
		fmt.Println("  (Tidak ada jawaban kuis)")
	}

	fmt.Println("\nNilai Kuis:")
	if penilaianSiswa, ok := guru.Penilaian[usernameSiswa]; ok {
		headerNilai := []string{"Kuis", "Nilai"}
		var dataNilai [][]string
		for kuis, nilai := range penilaianSiswa {
			dataNilai = append(dataNilai, []string{kuis, strconv.Itoa(nilai)})
		}
		printTable(headerNilai, dataNilai)
	} else {
		fmt.Println("  (Belum ada nilai)")
	}

	fmt.Println("\nJawaban Tugas:")
	if len(siswa.Tugas) > 0 {
		for judulTugas, jawaban := range siswa.Tugas {
			fmt.Printf("- %s:\n", judulTugas)
			fmt.Printf("  %s\n", jawaban)
		}
	} else {
		fmt.Println("  (Tidak ada jawaban tugas)")
	}

	fmt.Println("\nNilai Tugas:")
	if penilaianSiswa, ok := guru.NilaiTugas[usernameSiswa]; ok {
		headerNilai := []string{"Tugas", "Nilai"}
		var dataNilai [][]string
		for tugas, nilai := range penilaianSiswa {
			dataNilai = append(dataNilai, []string{tugas, strconv.Itoa(nilai)})
		}
		printTable(headerNilai, dataNilai)
	} else {
		fmt.Println("  (Belum ada nilai)")
	}

	fmt.Print("\nApakah Anda ingin memberikan penilaian untuk kuis? (y/n): ")
	var beriPenilaian string
	fmt.Scanln(&beriPenilaian)

	if strings.ToLower(beriPenilaian) == "y" {
		lms.berikanPenilaian(guru, usernameSiswa)
	}

	fmt.Print("\nApakah Anda ingin memberikan penilaian untuk tugas? (y/n): ")
	var beriPenilaianTugas string
	fmt.Scanln(&beriPenilaianTugas)

	if strings.ToLower(beriPenilaianTugas) == "y" {
		lms.berikanPenilaianTugas(guru, usernameSiswa)
	}
}

// berikanPenilaianTugas ...
func (lms *LMS) berikanPenilaianTugas(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna)

	fmt.Println("\nDaftar Tugas yang Telah Dikerjakan:")
	i := 1
	for judulTugas := range siswa.Tugas {
		fmt.Printf("%d. %s\n", i, judulTugas)
		i++
	}

	fmt.Print("Pilih nomor tugas untuk diberi nilai: ")
	var pilihanTugas int
	fmt.Scanln(&pilihanTugas)

	i = 1
	var judulTugasTerpilih string
	for judulTugas := range siswa.Tugas {
		if i == pilihanTugas {
			judulTugasTerpilih = judulTugas
			break
		}
		i++
	}

	if judulTugasTerpilih != "" {
		fmt.Print("Masukkan nilai untuk tugas '" + judulTugasTerpilih + "': ")
		var nilai int
		fmt.Scanln(&nilai)

		if guru.NilaiTugas[usernameSiswa] == nil {
			guru.NilaiTugas[usernameSiswa] = make(map[string]int)
		}
		guru.NilaiTugas[usernameSiswa][judulTugasTerpilih] = nilai
		fmt.Println("Penilaian berhasil diberikan!")
	} else {
		fmt.Println("Pilihan tugas tidak valid.")
	}
}

// berikanPenilaian ...
func (lms *LMS) berikanPenilaian(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna)

	fmt.Println("\nDaftar Kuis yang Telah Dikerjakan:")
	i := 1
	for judulKuis := range siswa.JawabanKuis {
		fmt.Printf("%d. %s\n", i, judulKuis)
		i++
	}

	fmt.Print("Pilih nomor kuis untuk diberi nilai: ")
	var pilihanKuis int
	fmt.Scanln(&pilihanKuis)

	i = 1
	var judulKuisTerpilih string
	for judulKuis := range siswa.JawabanKuis {
		if i == pilihanKuis {
			judulKuisTerpilih = judulKuis
			break
		}
		i++
	}

	if judulKuisTerpilih != "" {
		fmt.Print("Masukkan nilai untuk kuis '" + judulKuisTerpilih + "': ")
		var nilai int
		fmt.Scanln(&nilai)

		if guru.Penilaian[usernameSiswa] == nil {
			guru.Penilaian[usernameSiswa] = make(map[string]int)
		}
		guru.Penilaian[usernameSiswa][judulKuisTerpilih] = nilai
		fmt.Println("Penilaian berhasil diberikan!")
	} else {
		fmt.Println("Pilihan kuis tidak valid.")
	}
}

// LihatMateri ...
func (lms *LMS) LihatMateri(pengguna *Pengguna) {
	header := []string{"Judul Materi", "Konten"}
	var data [][]string

	for judul, konten := range lms.Materi {
		data = append(data, []string{judul, konten})
	}

	printTable(header, data)
}

// LihatKuis ...
func (lms *LMS) LihatKuis(pengguna *Pengguna) {
	header := []string{"Judul Kuis"}
	var data [][]string

	for judul := range lms.Kuis {
		data = append(data, []string{judul})
	}

	printTable(header, data)
}

// LihatTugas ...
func (lms *LMS) LihatTugas(pengguna *Pengguna) {
	header := []string{"Judul Tugas"}
	var data [][]string

	for judul := range lms.Tugas {
		data = append(data, []string{judul})
	}

	printTable(header, data)
}

// KerjakanKuis ...
func (lms *LMS) KerjakanKuis(pengguna *Pengguna, judulKuis string) {
	if pertanyaan, ok := lms.Kuis[judulKuis]; ok {
		var jawaban []string
		for _, p := range pertanyaan {
			fmt.Println(p)
			var jawabanStr string
			fmt.Scanln(&jawabanStr)
			jawaban = append(jawaban, jawabanStr)
		}

		pengguna.JawabanKuis[judulKuis] = jawaban // Simpan jawaban siswa
		fmt.Println("Jawaban kuis berhasil disimpan!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// SerahkanTugas ...
func (lms *LMS) SerahkanTugas(pengguna *Pengguna, judulTugas string) {
	if _, ok := lms.Tugas[judulTugas]; ok {
		fmt.Print("Masukkan jawaban tugas Anda: ")
		var jawabanTugas string
		fmt.Scanln(&jawabanTugas)

		pengguna.Tugas[judulTugas] = jawabanTugas // Simpan jawaban tugas
		fmt.Println("Jawaban tugas berhasil disimpan!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

// IkutForum ...
func (lms *LMS) IkutForum(pengguna *Pengguna) {
	lms.LihatForum(pengguna) // First, display forum titles

	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ")
	var judul string
	fmt.Scanln(&judul)

	if postingan, ok := lms.Forum[judul]; ok {
		fmt.Println("Postingan di forum ini:")
		for _, p := range postingan {
			fmt.Println(p)
		}

		fmt.Print("Masukkan postingan Anda: ")
		postinganBaru, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		postinganBaru = strings.TrimSpace(postinganBaru)

		lms.Forum[judul] = append(lms.Forum[judul], postinganBaru)
		fmt.Println("Postingan Anda berhasil ditambahkan!")
		return
	}

	fmt.Println("Forum tidak ditemukan.")
}

// LihatNilaiKuis ...
func (lms *LMS) LihatNilaiKuis(pengguna *Pengguna) {
	for _, p := range lms.Pengguna {
		if guru, ok := p.(*Guru); ok {
			if penilaianSiswa, ok := guru.Penilaian[pengguna.Username]; ok {
				header := []string{"Judul Kuis", "Nilai"}
				var data [][]string
				for judulKuis, nilai := range penilaianSiswa {
					data = append(data, []string{judulKuis, strconv.Itoa(nilai)})
				}
				printTable(header, data)
				return
			}
		}
	}
	fmt.Println("Nilai kuis belum tersedia.")
}

// LihatNilaiTugas ...
func (lms *LMS) LihatNilaiTugas(pengguna *Pengguna) {
	for _, p := range lms.Pengguna {
		if guru, ok := p.(*Guru); ok {
			if penilaianSiswa, ok := guru.NilaiTugas[pengguna.Username]; ok {
				header := []string{"Judul Tugas", "Nilai"}
				var data [][]string
				for judulTugas, nilai := range penilaianSiswa {
					data = append(data, []string{judulTugas, strconv.Itoa(nilai)})
				}
				printTable(header, data)
				return
			}
		}
	}
	fmt.Println("Nilai tugas belum tersedia.")
}

// sequentialSearch ...
func sequentialSearch(data []string, target string) int {
	for i, item := range data {
		if item == target {
			return i
		}
	}
	return -1
}

// binarySearch ...
func binarySearch(data []string, target string) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// selectionSort ...
func selectionSort(data []string, ascending bool) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if data[j] < data[minIndex] {
					minIndex = j
				}
			} else {
				if data[j] > data[minIndex] {
					minIndex = j
				}
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}

// insertionSort ...
func insertionSort(data []string, ascending bool) {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && (ascending && data[j] > key || !ascending && data[j] < key) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Get a slice of material titles (keys from the Materi map)
func (lms *LMS) getMateriJudul() []string {
	var titles []string
	for title := range lms.Materi {
		titles = append(titles, title)
	}
	return titles
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inisialisasi random number generator

	lms := LMS{
		Pengguna: make(map[string]interface{}),
		Materi:   make(map[string]string),
		Kuis:     make(map[string][]string),
		Forum:    make(map[string][]string),
		Tugas:    make(map[string]string),
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
		"Elsa Melisa Silaen103062300075",
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
		"Saya yang merancang golang?",
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
