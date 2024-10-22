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

// Pengguna merepresentasikan entitas pengguna dalam sistem LMS.
type Pengguna struct {
	Username    string
	Password    string
	Peran       string
	JawabanKuis map[string][]string // Kunci: Judul Kuis, Nilai: Jawaban
	Postingan   map[string][]string // Kunci: Judul Forum, Nilai: Postingan
	NilaiKuis   map[string]int      // Kunci: Judul Kuis, Nilai: Nilai
}

// Guru merupakan tipe khusus Pengguna dengan tambahan kolom PENILAIAN.
type Guru struct {
	*Pengguna
	Penilaian map[string]map[string]int // Kunci: Username Siswa, Nilai: map[Judul Kuis]Nilai
}

// LMS merepresentasikan sistem Learning Management System (LMS).
type LMS struct {
	Pengguna map[string]interface{} // Kunci: Username, Nilai: *Pengguna atau *Guru
	Materi   map[string]string      // Kunci: Judul Materi, Nilai: Konten
	Kuis     map[string][]string    // Kunci: Judul Kuis, Nilai: Pertanyaan
	Forum    map[string][]string    // Kunci: Judul Forum, Nilai: Postingan
}

// Fungsi untuk mencetak tabel dengan header dan data
func printTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
}

// Fungsi untuk mencetak menu dengan pilihan
func printMenu(menuItems []string) {
	fmt.Println()
	for i, item := range menuItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Pilihan Anda: ")
}

// Fungsi untuk menampilkan animasi spinner
func showSpinner(message string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + message
	s.Start()
	time.Sleep(2 * time.Second) // Simulasi proses
	s.Stop()
}

// Daftar mendaftarkan pengguna baru ke dalam sistem LMS.
func (lms *LMS) Daftar(username, password, peran string) {
	// Menghilangkan spasi dari username
	username = strings.ReplaceAll(username, " ", "")

	showSpinner("Mendaftarkan pengguna...")
	if _, ok := lms.Pengguna[username]; !ok {
		if peran == "guru" {
			lms.Pengguna[username] = &Guru{
				Pengguna: &Pengguna{
					Username:    username, // Menyimpan username tanpa spasi
					Password:    password,
					Peran:       peran,
					JawabanKuis: make(map[string][]string),
					Postingan:   make(map[string][]string),
					NilaiKuis:   make(map[string]int),
				},
				Penilaian: make(map[string]map[string]int),
			}
		} else {
			lms.Pengguna[username] = &Pengguna{
				Username:    username, // Menyimpan username tanpa spasi
				Password:    password,
				Peran:       peran,
				JawabanKuis: make(map[string][]string),
				Postingan:   make(map[string][]string),
				NilaiKuis:   make(map[string]int),
			}
		}
		fmt.Printf("Pengguna %s berhasil terdaftar!\n", username)
	} else {
		fmt.Println("Username sudah ada.")
	}
}

// Masuk memungkinkan pengguna untuk masuk ke sistem LMS.
func (lms *LMS) Masuk(username, password string) interface{} {
	// Menghilangkan spasi dari username
	username = strings.ReplaceAll(username, " ", "")

	showSpinner("Memeriksa kredensial...")
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

// MenuGuru menampilkan menu untuk peran guru.
func (lms *LMS) MenuGuru(guru *Guru) {
	menuItems := []string{
		"Buat Materi",
		"Edit Materi",
		"Hapus Materi",
		"Buat Kuis",
		"Edit Kuis",
		"Hapus Kuis",
		"Buat Forum",
		"Data Siswa",
		"Keluar",
	}

	for {
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
			lms.BuatForum(guru.Pengguna)
		case 8:
			lms.DataSiswa(guru)
		case 9:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// MenuSiswa menampilkan menu untuk peran siswa.
func (lms *LMS) MenuSiswa(pengguna *Pengguna) {
	menuItems := []string{
		"Materi",
		"Kuis",
		"Forum",
		"Lihat Nilai Kuis",
		"Keluar",
	}

	for {
		printMenu(menuItems)
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lms.LihatMateri(pengguna)
		case 2:
			lms.KerjakanKuis(pengguna)
		case 3:
			lms.IkutForum(pengguna)
		case 4:
			lms.LihatNilaiKuis(pengguna)
		case 5:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// BuatMateri memungkinkan guru untuk membuat materi baru.
func (lms *LMS) BuatMateri(pengguna *Pengguna) {
	var judul, konten string
	fmt.Print("Masukkan judul materi: ")
	fmt.Scanln(&judul)
	fmt.Print("Masukkan konten materi: ")
	fmt.Scanln(&konten)

	lms.Materi[judul] = konten
	fmt.Println("Materi berhasil dibuat!")
}

// EditMateri memungkinkan guru untuk mengedit materi yang ada.
func (lms *LMS) EditMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Materi[judul]; ok {
		fmt.Print("Masukkan konten baru: ")
		var kontenBaru string
		fmt.Scanln(&kontenBaru)

		lms.Materi[judul] = kontenBaru
		fmt.Println("Materi berhasil diupdate!")
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

// HapusMateri memungkinkan guru untuk menghapus materi.
func (lms *LMS) HapusMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi yang ingin Anda hapus: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Materi[judul]; ok {
		delete(lms.Materi, judul)
		fmt.Println("Materi berhasil dihapus!")
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

// BuatKuis memungkinkan guru untuk membuat kuis baru.
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

// EditKuis memungkinkan guru untuk mengedit kuis yang ada.
func (lms *LMS) EditKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis yang ingin Anda edit: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Kuis[judul]; ok {
		var pertanyaanBaru []string
		for {
			fmt.Print("Masukkan pertanyaan baru (atau ketik 'selesai' untuk selesai): ")
			pertanyaanStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			pertanyaanStr = strings.TrimSpace(pertanyaanStr)

			if pertanyaanStr == "selesai" {
				break
			}

			pertanyaanBaru = append(pertanyaanBaru, pertanyaanStr)
		}

		lms.Kuis[judul] = pertanyaanBaru
		fmt.Println("Kuis berhasil diupdate!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// HapusKuis memungkinkan guru untuk menghapus kuis.
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

// BuatForum memungkinkan guru untuk membuat forum baru.
func (lms *LMS) BuatForum(pengguna *Pengguna) {
	fmt.Print("Masukkan judul forum: ")
	var judul string
	fmt.Scanln(&judul)

	lms.Forum[judul] = []string{}
	fmt.Println("Forum berhasil dibuat!")
}

// LihatMateri memungkinkan siswa untuk melihat materi yang tersedia.
func (lms *LMS) LihatMateri(pengguna *Pengguna) {
	fmt.Println("\nDaftar Materi:")
	header := []string{"Judul", "Konten"}
	var data [][]string
	for judul, konten := range lms.Materi {
		data = append(data, []string{judul, konten})
	}
	printTable(header, data)
}

// KerjakanKuis memungkinkan siswa untuk mengerjakan kuis yang tersedia.
func (lms *LMS) KerjakanKuis(pengguna *Pengguna) {
	fmt.Print("Masukkan judul kuis yang ingin Anda kerjakan: ")
	var judul string
	fmt.Scanln(&judul)

	if pertanyaan, ok := lms.Kuis[judul]; ok {
		var jawaban []string
		for _, tanya := range pertanyaan {
			fmt.Println(tanya)
			var jawab string
			fmt.Scanln(&jawab)
			jawaban = append(jawaban, jawab)
		}
		pengguna.JawabanKuis[judul] = jawaban
		fmt.Println("Jawaban Anda telah disimpan.")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// IkutForum memungkinkan siswa untuk berpartisipasi dalam forum yang ada.
func (lms *LMS) IkutForum(pengguna *Pengguna) {
	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ")
	var judul string
	fmt.Scanln(&judul)

	if _, ok := lms.Forum[judul]; ok {
		for {
			fmt.Print("Masukkan pesan (atau ketik 'selesai' untuk selesai): ")
			pesan, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			pesan = strings.TrimSpace(pesan)

			if pesan == "selesai" {
				break
			}

			lms.Forum[judul] = append(lms.Forum[judul], fmt.Sprintf("%s: %s", pengguna.Username, pesan))
		}
		return
	}

	fmt.Println("Forum tidak ditemukan.")
}

// LihatNilaiKuis memungkinkan siswa untuk melihat nilai kuis mereka.
func (lms *LMS) LihatNilaiKuis(pengguna *Pengguna) {
	fmt.Println("\nNilai Kuis:")
	header := []string{"Judul Kuis", "Nilai"}
	var data [][]string
	for judul, nilai := range pengguna.NilaiKuis {
		data = append(data, []string{judul, strconv.Itoa(nilai)})
	}
	printTable(header, data)
}

// DataSiswa menampilkan data siswa termasuk nilai-nilai kuis.
func (lms *LMS) DataSiswa(guru *Guru) {
	fmt.Println("\nData Siswa:")

	// Menampilkan semua siswa
	headerSiswa := []string{"Username"}
	var dataSiswa []string
	for username, pengguna := range lms.Pengguna {
		if siswa, ok := pengguna.(*Pengguna); ok && siswa.Peran == "siswa" {
			dataSiswa = append(dataSiswa, username)
		}
	}

	// Mengurutkan data siswa menggunakan selection sort
	selectionSort(dataSiswa)

	var data [][]string
	for _, username := range dataSiswa {
		data = append(data, []string{username})
	}
	printTable(headerSiswa, data)

	// Meminta guru untuk memilih siswa
	fmt.Print("Masukkan username siswa untuk melihat detail atau memberi nilai: ")
	var usernameSiswa string
	fmt.Scanln(&usernameSiswa)

	if siswa, ok := lms.Pengguna[usernameSiswa]; ok {
		if _, ok := siswa.(*Pengguna); ok {
			// Menampilkan data detail dan nilai siswa
			lms.tampilkanDataSiswa(guru, usernameSiswa)
		} else {
			fmt.Println("Pengguna tersebut bukan siswa.")
		}
	} else {
		fmt.Println("Siswa tidak ditemukan.")
	}
}

// Fungsi selectionSort untuk mengurutkan slice string
func selectionSort(data []string) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

// tampilkanDataSiswa menampilkan data detail dan nilai siswa.
func (lms *LMS) tampilkanDataSiswa(guru *Guru, usernameSiswa string) {
	siswa := lms.Pengguna[usernameSiswa].(*Pengguna)
	fmt.Printf("\nDetail Siswa - %s\n", usernameSiswa)

	// Menampilkan nilai-nilai kuis siswa
	fmt.Println("Nilai Kuis:")
	headerNilai := []string{"Judul Kuis", "Nilai"}
	var dataNilai [][]string
	for judulKuis, nilai := range siswa.NilaiKuis {
		dataNilai = append(dataNilai, []string{judulKuis, strconv.Itoa(nilai)})
	}
	printTable(headerNilai, dataNilai)

	// Memberi nilai atau mengedit nilai kuis siswa
	fmt.Print("Apakah Anda ingin memberi nilai atau mengedit nilai kuis siswa ini? (y/n): ")
	var jawab string
	fmt.Scanln(&jawab)

	if strings.ToLower(jawab) == "y" {
		fmt.Print("Masukkan judul kuis: ")
		var judulKuis string
		fmt.Scanln(&judulKuis)

		if _, ok := lms.Kuis[judulKuis]; ok {
			fmt.Print("Masukkan nilai: ")
			var nilai int
			fmt.Scanln(&nilai)

			// Menyimpan nilai dalam data siswa dan data penilaian guru
			siswa.NilaiKuis[judulKuis] = nilai
			if guru.Penilaian[usernameSiswa] == nil {
				guru.Penilaian[usernameSiswa] = make(map[string]int)
			}
			guru.Penilaian[usernameSiswa][judulKuis] = nilai

			fmt.Println("Nilai berhasil disimpan!")
		} else {
			fmt.Println("Kuis tidak ditemukan.")
		}
	}
}

// fungsi utama
func main() {
	lms := LMS{
		Pengguna: make(map[string]interface{}),
		Materi:   make(map[string]string),
		Kuis:     make(map[string][]string),
		Forum:    make(map[string][]string),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nSelamat datang di Learning Management System!")
		fmt.Println("1. Daftar")
		fmt.Println("2. Masuk")
		fmt.Print("Pilihan Anda: ")
		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Print("Masukkan username: ")
			scanner.Scan()
			username := scanner.Text()
			fmt.Print("Masukkan password: ")
			scanner.Scan()
			password := scanner.Text()
			fmt.Print("Masukkan peran (guru/siswa): ")
			scanner.Scan()
			peran := scanner.Text()

			lms.Daftar(username, password, peran)
		case "2":
			fmt.Print("Masukkan username: ")
			scanner.Scan()
			username := scanner.Text()
			fmt.Print("Masukkan password: ")
			scanner.Scan()
			password := scanner.Text()

			pengguna := lms.Masuk(username, password)
			if pengguna != nil {
				if pengguna.(*Pengguna).Peran == "guru" {
					lms.MenuGuru(pengguna.(*Guru))
				} else {
					lms.MenuSiswa(pengguna.(*Pengguna))
				}
			}
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
