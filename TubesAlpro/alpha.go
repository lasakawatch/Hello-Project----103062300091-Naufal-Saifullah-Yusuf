package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
)

const (
	maxPengguna = 100
	maxMateri   = 100
	maxKuis     = 100
	maxTugas    = 100
	maxForum    = 100
)

var (
	penggunaList [maxPengguna]interface{}
	materiList   [maxMateri]string
	kuisList     [maxKuis][]string
	tugasList    [maxTugas]*Tugas
	forumList    [maxForum][]string
	nextPengguna int
	nextMateri   int
	nextKuis     int
	nextTugas    int
	nextForum    int
)

// Pengguna ...
type Pengguna struct {
	Username      string
	Password      string
	Peran         string
	JawabanKuis  map[string][]string
	JawabanTugas map[string]string
	Postingan     map[string][]string
	NilaiKuis     map[string]int
	NilaiTugas    map[string]int
}

// Guru ...
type Guru struct {
	Pengguna
	Penilaian   map[string]map[string]int
	NilaiTugas map[string]map[string]int
}

// Tugas ...
type Tugas struct {
	Judul    string
	Soal     string
	NilaiSiswa map[string]int
}

func printTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
}

func printMenu(menuItems []string) {
	fmt.Println()
	for i, item := range menuItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Pilihan Anda: ")
}

func showSpinner(message string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + message
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()
}

// Daftar ...
func Daftar(username, password, peran string) {
	username = strings.ReplaceAll(username, " ", "")
	showSpinner("Mendaftarkan pengguna...")

	if _, ok := sequentialSearch(username); !ok {
		if peran == "guru" {
			penggunaList[nextPengguna] = &Guru{
				Pengguna: Pengguna{
					Username:      username,
					Password:      password,
					Peran:         peran,
					JawabanKuis:  make(map[string][]string),
					JawabanTugas: make(map[string]string),
					Postingan:     make(map[string][]string),
					NilaiKuis:     make(map[string]int),
					NilaiTugas:    make(map[string]int),
				},
				Penilaian:   make(map[string]map[string]int),
				NilaiTugas: make(map[string]map[string]int),
			}
		} else {
			penggunaList[nextPengguna] = &Pengguna{
				Username:      username,
				Password:      password,
				Peran:         peran,
				JawabanKuis:  make(map[string][]string),
				JawabanTugas: make(map[string]string),
				Postingan:     make(map[string][]string),
				NilaiKuis:     make(map[string]int),
				NilaiTugas:    make(map[string]int),
			}
		}
		fmt.Printf("Pengguna %s berhasil terdaftar!\n", username)
		nextPengguna++
	} else {
		fmt.Println("Username sudah ada.")
	}
}

// Masuk ...
func Masuk(username, password string) interface{} {
	username = strings.ReplaceAll(username, " ", "")
	showSpinner("Memeriksa kredensial...")

	if index, ok := sequentialSearch(username); ok {
		switch p := penggunaList[index].(type) {
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
func MenuGuru(guru *Guru) {
	menuItems := []string{
		"Buat Materi", "Edit Materi", "Hapus Materi",
		"Buat Kuis", "Edit Kuis", "Hapus Kuis",
		"Buat Tugas", "Edit Tugas", "Hapus Tugas",
		"Buat Forum", "Data Siswa", "Keluar",
	}

	for {
		printMenu(menuItems)
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			buatMateri()
		case 2:
			editMateri()
		case 3:
			hapusMateri()
		case 4:
			buatKuis()
		case 5:
			editKuis()
		case 6:
			hapusKuis()
		case 7:
			buatTugas()
		case 8:
			editTugas()
		case 9:
			hapusTugas()
		case 10:
			buatForum()
		case 11:
			dataSiswa(guru)
		case 12:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// MenuSiswa ...
func MenuSiswa(pengguna *Pengguna) {
	menuItems := []string{
		"Materi", "Kuis", "Tugas", "Forum",
		"Lihat Nilai Kuis", "Lihat Nilai Tugas", "Keluar",
	}

	for {
		printMenu(menuItems)
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatMateri()
		case 2:
			kerjakanKuis(pengguna)
		case 3:
			kerjakanTugas(pengguna)
		case 4:
			ikutForum(pengguna)
		case 5:
			lihatNilaiKuis(pengguna)
		case 6:
			lihatNilaiTugas(pengguna)
		case 7:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ===================== MATERI =====================
func buatMateri() {
	var judul, konten string
	fmt.Print("Masukkan judul materi: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan konten materi: ")
	fmt.Scanln(&konten)

	materiList[nextMateri] = judul + ": " + konten
	fmt.Println("Materi berhasil dibuat!")
	nextMateri++
}

func editMateri() {
	fmt.Print("Masukkan judul materi yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchMateri(judul); ok {
		fmt.Print("Masukkan konten baru: ")
		var kontenBaru string
		fmt.Scanln(&kontenBaru)

		materiList[index] = judul + ": " + kontenBaru
		fmt.Println("Materi berhasil diupdate!")
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

func hapusMateri() {
	fmt.Print("Masukkan judul materi yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchMateri(judul); ok {
		for i := index; i < nextMateri-1; i++ {
			materiList[i] = materiList[i+1]
		}
		nextMateri--
		fmt.Println("Materi berhasil dihapus!")
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

func lihatMateri() {
	header := []string{"Judul Materi", "Konten"}
	var data [][]string

	for i := 0; i < nextMateri; i++ {
		parts := strings.Split(materiList[i], ": ")
		if len(parts) == 2 {
			data = append(data, []string{parts[0], parts[1]})
		}
	}

	printTable(header, data)
}

// ===================== KUIS =====================
func buatKuis() {
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

		pertanyaan = append(pertanyaan, judul+": "+pertanyaanStr)
	}

	kuisList[nextKuis] = pertanyaan
	fmt.Println("Kuis berhasil dibuat!")
	nextKuis++
}

func editKuis() {
	fmt.Print("Masukkan judul kuis yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchKuis(judul); ok {
		var pertanyaan []string
		for {
			fmt.Print("Masukkan pertanyaan (atau ketik 'selesai' untuk selesai): ")
			pertanyaanStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			pertanyaanStr = strings.TrimSpace(pertanyaanStr)

			if pertanyaanStr == "selesai" {
				break
			}

			pertanyaan = append(pertanyaan, judul+": "+pertanyaanStr)
		}
		kuisList[index] = pertanyaan
		fmt.Println("Kuis berhasil diupdate!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

func hapusKuis() {
	fmt.Print("Masukkan judul kuis yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchKuis(judul); ok {
		for i := index; i < nextKuis-1; i++ {
			kuisList[i] = kuisList[i+1]
		}
		nextKuis--
		fmt.Println("Kuis berhasil dihapus!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

func tampilkanDaftarKuis() {
	header := []string{"Judul Kuis"}
	var data [][]string

	for i := 0; i < nextKuis; i++ {
		judul := getJudulKuis(kuisList[i])
		if judul != "" {
			data = append(data, []string{judul})
		}
	}

	printTable(header, data)
}

func kerjakanKuis(pengguna *Pengguna) {
	tampilkanDaftarKuis()

	fmt.Print("Masukkan judul kuis yang ingin Anda kerjakan: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchKuis(judul); ok {
		fmt.Println("\nKuis:", judul)
		var jawaban []string
		for _, p := range kuisList[index] {
			fmt.Println(p)
			var jawabanStr string
			fmt.Scanln(&jawabanStr)
			jawaban = append(jawaban, jawabanStr)
		}

		pengguna.JawabanKuis[judul] = jawaban
		fmt.Println("Jawaban kuis berhasil disimpan!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

func getJudulKuis(pertanyaan []string) string {
	if len(pertanyaan) > 0 {
		parts := strings.Split(pertanyaan[0], ": ")
		if len(parts) == 2 {
			return parts[0]
		}
	}
	return ""
}

// ===================== TUGAS =====================
func buatTugas() {
	var judul, soal string
	fmt.Print("Masukkan judul tugas: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan soal tugas: ")
	fmt.Scanln(&soal)

	tugasList[nextTugas] = &Tugas{
		Judul:    judul,
		Soal:     soal,
		NilaiSiswa: make(map[string]int),
	}
	fmt.Println("Tugas berhasil dibuat!")
	nextTugas++
}

func editTugas() {
	fmt.Print("Masukkan judul tugas yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchTugas(judul); ok {
		fmt.Print("Masukkan soal tugas baru: ")
		var soalBaru string
		fmt.Scanln(&soalBaru)

		tugasList[index].Soal = soalBaru
		fmt.Println("Tugas berhasil diupdate!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

func hapusTugas() {
	fmt.Print("Masukkan judul tugas yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchTugas(judul); ok {
		for i := index; i < nextTugas-1; i++ {
			tugasList[i] = tugasList[i+1]
		}
		nextTugas--
		fmt.Println("Tugas berhasil dihapus!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

func tampilkanDaftarTugas() {
	header := []string{"Judul Tugas"}
	var data [][]string

	for i := 0; i < nextTugas; i++ {
		if tugasList[i] != nil {
			data = append(data, []string{tugasList[i].Judul})
		}
	}

	printTable(header, data)
}

func kerjakanTugas(pengguna *Pengguna) {
	tampilkanDaftarTugas()

	fmt.Print("Masukkan judul tugas yang ingin Anda kerjakan: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchTugas(judul); ok {
		fmt.Println("\nTugas:", judul)
		fmt.Println(tugasList[index].Soal)
		fmt.Print("Masukkan jawaban Anda: ")
		var jawaban string
		fmt.Scanln(&jawaban)

		pengguna.JawabanTugas[judul] = jawaban
		fmt.Println("Jawaban tugas berhasil disimpan!")
		return
	}

	fmt.Println("Tugas tidak ditemukan.")
}

// ===================== FORUM =====================
func buatForum() {
	var judul, postingan string
	fmt.Print("Masukkan judul forum: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan postingan pertama: ")
	postingan, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	postingan = strings.TrimSpace(postingan)

	forumList[nextForum] = []string{judul, postingan}
	fmt.Println("Forum berhasil dibuat!")
	nextForum++
}

func tampilkanDaftarForum() {
	header := []string{"Judul Forum"}
	var data [][]string

	for i := 0; i < nextForum; i++ {
		if len(forumList[i]) > 0 {
			judul := forumList[i][0]
			data = append(data, []string{judul})
		}
	}

	printTable(header, data)
}

func ikutForum(pengguna *Pengguna) {
	tampilkanDaftarForum()

	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ")
	var judul string
	fmt.Scanln(&judul)

	if index, ok := sequentialSearchForum(judul); ok {
		fmt.Println("\nForum:", judul)
		fmt.Println("Postingan di forum ini:")
		for _, p := range forumList[index][1:] {
			fmt.Println(p)
		}

		fmt.Print("Masukkan postingan Anda: ")
		postinganBaru, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		postinganBaru = strings.TrimSpace(postinganBaru)

		forumList[index] = append(forumList[index], postinganBaru)
		fmt.Println("Postingan Anda berhasil ditambahkan!")
		return
	}

	fmt.Println("Forum tidak ditemukan.")
}

// ===================== DATA SISWA =====================
func dataSiswa(guru *Guru) {
	fmt.Println("\nData Siswa:")

	headerSiswa := []string{"Username"}
	var dataSiswa [][]string
	for i := 0; i < nextPengguna; i++ {
		if p, ok := penggunaList[i].(*Pengguna); ok && p.Peran == "siswa" {
			dataSiswa = append(dataSiswa, []string{p.Username})
		}
	}
	printTable(headerSiswa, dataSiswa)

	fmt.Print("Masukkan username siswa untuk melihat detail atau memberi nilai: ")
	var usernameSiswa string
	fmt.Scanln(&usernameSiswa)

	if index, ok := sequentialSearch(usernameSiswa); ok {
		if _, ok := penggunaList[index].(*Pengguna); ok {
			tampilkanDataSiswa(guru, index)
		} else {
			fmt.Println("Pengguna tersebut bukan siswa.")
		}
	} else {
		fmt.Println("Siswa tidak ditemukan.")
	}
}

func tampilkanDataSiswa(guru *Guru, index int) {
	siswa := penggunaList[index].(*Pengguna)

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
	if penilaianSiswa, ok := guru.Penilaian[siswa.Username]; ok {
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
	if len(siswa.JawabanTugas) > 0 {
		for judulTugas, jawaban := range siswa.JawabanTugas {
			fmt.Printf("- %s:\n", judulTugas)
			fmt.Printf("  %s\n", jawaban)
		}
	} else {
		fmt.Println("  (Tidak ada jawaban tugas)")
	}

	fmt.Println("\nNilai Tugas:")
	if nilaiTugasSiswa, ok := guru.NilaiTugas[siswa.Username]; ok {
		headerNilai := []string{"Tugas", "Nilai"}
		var dataNilai [][]string
		for tugas, nilai := range nilaiTugasSiswa {
			dataNilai = append(dataNilai, []string{tugas, strconv.Itoa(nilai)})
		}
		printTable(headerNilai, dataNilai)
	} else {
		fmt.Println("  (Belum ada nilai)")
	}

	fmt.Print("\nApakah Anda ingin memberikan penilaian? (y/n): ")
	var beriPenilaian string
	fmt.Scanln(&beriPenilaian)

	if strings.ToLower(beriPenilaian) == "y" {
		berikanPenilaian(guru, index)
	}
}

func berikanPenilaian(guru *Guru, index int) {
	siswa := penggunaList[index].(*Pengguna)

	fmt.Println("\nPilih jenis penilaian:")
	fmt.Println("1. Kuis")
	fmt.Println("2. Tugas")
	fmt.Print("Pilihan Anda: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
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

			if guru.Penilaian[siswa.Username] == nil {
				guru.Penilaian[siswa.Username] = make(map[string]int)
			}
			guru.Penilaian[siswa.Username][judulKuisTerpilih] = nilai
			fmt.Println("Penilaian berhasil diberikan!")
		} else {
			fmt.Println("Pilihan kuis tidak valid.")
		}
	} else if pilihan == 2 {
		fmt.Println("\nDaftar Tugas yang Telah Dikerjakan:")
		i := 1
		for judulTugas := range siswa.JawabanTugas {
			fmt.Printf("%d. %s\n", i, judulTugas)
			i++
		}

		fmt.Print("Pilih nomor tugas untuk diberi nilai: ")
		var pilihanTugas int
		fmt.Scanln(&pilihanTugas)

		i = 1
		var judulTugasTerpilih string
		for judulTugas := range siswa.JawabanTugas {
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

			if guru.NilaiTugas[siswa.Username] == nil {
				guru.NilaiTugas[siswa.Username] = make(map[string]int)
			}
			guru.NilaiTugas[siswa.Username][judulTugasTerpilih] = nilai
			fmt.Println("Penilaian berhasil diberikan!")
		} else {
			fmt.Println("Pilihan tugas tidak valid.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func lihatNilaiKuis(penggunaLogin *Pengguna) {
	for i := 0; i < nextPengguna; i++ {
		if guru, ok := penggunaList[i].(*Guru); ok {
			if penilaianSiswa, ok := guru.Penilaian[penggunaLogin.Username]; ok {
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

func lihatNilaiTugas(penggunaLogin *Pengguna) {
	for i := 0; i < nextPengguna; i++ {
		if guru, ok := penggunaList[i].(*Guru); ok {
			if nilaiTugasSiswa, ok := guru.NilaiTugas[penggunaLogin.Username]; ok {
				header := []string{"Judul Tugas", "Nilai"}
				var data [][]string
				for judulTugas, nilai := range nilaiTugasSiswa {
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
func sequentialSearch(target string) (int, bool) {
	for i := 0; i < nextPengguna; i++ {
		if penggunaList[i] != nil {
			switch pengguna := penggunaList[i].(type) {
			case *Pengguna:
				if pengguna.Username == target {
					return i, true
				}
			case *Guru:
				if pengguna.Username == target {
					return i, true
				}
			}
		}
	}
	return -1, false
}

func sequentialSearchMateri(target string) (int, bool) {
	for i := 0; i < nextMateri; i++ {
		if strings.Contains(materiList[i], target+": ") {
			return i, true
		}
	}
	return -1, false
}

func sequentialSearchKuis(target string) (int, bool) {
	for i := 0; i < nextKuis; i++ {
		if getJudulKuis(kuisList[i]) == target {
			return i, true
		}
	}
	return -1, false
}

func sequentialSearchTugas(target string) (int, bool) {
	for i := 0; i < nextTugas; i++ {
		if tugasList[i] != nil && tugasList[i].Judul == target {
			return i, true
		}
	}
	return -1, false
}

func sequentialSearchForum(target string) (int, bool) {
	for i := 0; i < nextForum; i++ {
		if len(forumList[i]) > 0 && forumList[i][0] == target {
			return i, true
		}
	}
	return -1, false
}

func main() {
	penggunaList[0] = &Guru{
		Pengguna: Pengguna{
			Username:      "pakguru",
			Password:      "1234",
			Peran:         "guru",
			JawabanKuis:  make(map[string][]string),
			JawabanTugas: make(map[string]string),
			Postingan:     make(map[string][]string),
			NilaiKuis:     make(map[string]int),
			NilaiTugas:    make(map[string]int),
		},
		Penilaian:   make(map[string]map[string]int),
		NilaiTugas: make(map[string]map[string]int),
	}
	nextPengguna = 1

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

	for i, siswa := range daftarSiswa {
		Daftar(strings.ReplaceAll(siswa, " ", ""), fmt.Sprintf("password%d", i+1), "siswa")
	}

	var pilihan int
	for {
		fmt.Println("\n--- Selamat datang di LMS Universitas Telkom! ---")
		fmt.Println("1. Guru")
		fmt.Println("2. Siswa")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilihan)

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
				Daftar(username, password, "guru")
			case 2:
				var username, password string
				fmt.Print("Masukkan username: ")
				fmt.Scanln(&username)
				fmt.Print("Masukkan password: ")
				fmt.Scanln(&password)
				pengguna := Masuk(username, password)
				if pengguna != nil {
					if guru, ok := pengguna.(*Guru); ok {
						MenuGuru(guru)
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
				Daftar(username, password, "siswa")
			case 2:
				var username, password string
				fmt.Print("Masukkan username: ")
				fmt.Scanln(&username)
				fmt.Print("Masukkan password: ")
				fmt.Scanln(&password)
				pengguna := Masuk(username, password)
				if pengguna != nil {
					if siswa, ok := pengguna.(*Pengguna); ok {
						MenuSiswa(siswa)
					}
				}
			}
		case 3:
			fmt.Println("Keluar dari sistem...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}