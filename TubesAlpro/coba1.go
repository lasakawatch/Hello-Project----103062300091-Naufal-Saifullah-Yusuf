package main

import (
	"fmt"
	"strconv"
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
	Materi   map[string]string   // Kunci: Judul Materi, Nilai: Konten
	Kuis     map[string][]string // Kunci: Judul Kuis, Nilai: Pertanyaan
	Forum    map[string][]string // Kunci: Judul Forum, Nilai: Postingan
}

// Daftar mendaftarkan pengguna baru ke dalam sistem LMS.
func (lms *LMS) Daftar(username, password, peran string) {
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
				},
				Penilaian: make(map[string]map[string]int),
			}
		} else {
			lms.Pengguna[username] = &Pengguna{
				Username:    username,
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
	if pengguna, ok := lms.Pengguna[username]; ok {
		switch p := pengguna.(type) {
		case *Pengguna:
			if p.Password == password {
				fmt.Printf("Selamat datang %s!\n", username)
				return p
			}
			fmt.Println("Password salah.")
			return nil
		case *Guru:
			if p.Pengguna.Password == password {
				fmt.Printf("Selamat datang %s!\n", username)
				return p
			}
			fmt.Println("Password salah.")
			return nil
		default:
			fmt.Println("Error: Tipe pengguna tidak valid.")
			return nil
		}
	}
	fmt.Println("Pengguna tidak ditemukan. Silakan daftar terlebih dahulu.")
	return nil
}

// MenuGuru menampilkan menu untuk peran guru.
func (lms *LMS) MenuGuru(pengguna interface{}) {
	guru, ok := pengguna.(*Guru)
	if !ok {
		fmt.Println("Error: Pengguna bukan Guru.")
		return
	}
	for {
		fmt.Println("\nMenu Guru:")
		fmt.Println("1. Buat Materi")
		fmt.Println("2. Edit Materi")
		fmt.Println("3. Hapus Materi")
		fmt.Println("4. Buat Kuis")
		fmt.Println("5. Edit Kuis")
		fmt.Println("6. Hapus Kuis")
		fmt.Println("7. Buat Forum")
		fmt.Println("8. Data Siswa")
		fmt.Println("9. Keluar")

		fmt.Print("Masukkan pilihan Anda: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			lms.BuatMateri(guru.Pengguna)
		case "2":
			lms.EditMateri(guru.Pengguna)
		case "3":
			lms.HapusMateri(guru.Pengguna)
		case "4":
			lms.BuatKuis(guru.Pengguna)
		case "5":
			lms.EditKuis(guru.Pengguna)
		case "6":
			lms.HapusKuis(guru.Pengguna)
		case "7":
			lms.BuatForum(guru.Pengguna)
		case "8":
			lms.DataSiswa(guru)
		case "9":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// MenuSiswa menampilkan menu untuk peran siswa.
func (lms *LMS) MenuSiswa(pengguna *Pengguna) {
	for {
		fmt.Println("\nMenu Siswa:")
		fmt.Println("1. Materi")
		fmt.Println("2. Kuis")
		fmt.Println("3. Forum")
		fmt.Println("4. Keluar")

		fmt.Print("Masukkan pilihan Anda: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			lms.LihatMateri(pengguna)
		case "2":
			lms.KerjakanKuis(pengguna)
		case "3":
			lms.IkutForum(pengguna)
		case "4":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// BuatMateri memungkinkan guru untuk membuat materi baru.
func (lms *LMS) BuatMateri(pengguna *Pengguna) {
	fmt.Print("Masukkan judul materi: ")
	var judul string
	fmt.Scanln(&judul)

	fmt.Print("Masukkan konten materi: ")
	var konten string
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
		var pertanyaanStr string
		fmt.Scanln(&pertanyaanStr)

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
		var pertanyaan []string
		for {
			fmt.Print("Masukkan pertanyaan baru (atau ketik 'selesai' untuk selesai): ")
			var pertanyaanStr string
			fmt.Scanln(&pertanyaanStr)

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

// DataSiswa menampilkan data siswa kepada guru.
func (lms *LMS) DataSiswa(guru *Guru) {
	fmt.Println("\nSiswa Terdaftar:")
	for username, siswa := range lms.Pengguna {
		if _, ok := siswa.(*Pengguna); ok {
			fmt.Printf("- %s\n", username)
		}
	}

	fmt.Print("Masukkan username siswa yang ingin Anda lihat: ")
	var usernameSiswa string
	fmt.Scanln(&usernameSiswa)

	if siswa, ok := lms.Pengguna[usernameSiswa]; ok {
		if s, ok := siswa.(*Pengguna); ok {
			fmt.Printf("\nSiswa: %s\n", usernameSiswa)

			// Menampilkan Jawaban Kuis
			if len(s.JawabanKuis) > 0 {
				fmt.Println("\nJawaban Kuis:")
				for judulKuis, jawaban := range s.JawabanKuis {
					fmt.Printf("  - %s:\n", judulKuis)
					for i, jwb := range jawaban {
						fmt.Printf("    %d. %s\n", i+1, jwb)
					}
				}
			} else {
				fmt.Println("Tidak ada jawaban kuis.")
			}

			// Menampilkan Nilai Kuis
			if len(s.NilaiKuis) > 0 {
				fmt.Println("\nNilai Kuis:")
				for judulKuis, nilai := range s.NilaiKuis {
					fmt.Printf("  - %s: %d\n", judulKuis, nilai)
				}
			} else {
				fmt.Println("Tidak ada nilai kuis.")
			}

			// Menampilkan Postingan Forum
			if len(s.Postingan) > 0 {
				fmt.Println("\nPostingan Forum:")
				for judulForum, postingan := range s.Postingan {
					fmt.Printf("  - %s:\n", judulForum)
					for _, post := range postingan {
						fmt.Printf("    %s\n", post)
					}
				}
			} else {
				fmt.Println("Tidak ada postingan forum.")
			}

			// Menampilkan Penilaian Kuis
			if _, ok := guru.Penilaian[usernameSiswa]; ok {
				fmt.Println("\nPenilaian Kuis:")
				for judulKuis, nilai := range guru.Penilaian[usernameSiswa] {
					fmt.Printf("  - %s: %d\n", judulKuis, nilai)
				}
			} else {
				fmt.Println("Tidak ada penilaian kuis.")
			}

			// Menawarkan untuk memberikan penilaian
			fmt.Println("\nApakah Anda ingin memberikan penilaian untuk kuis?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")

			var pilihan string
			fmt.Scanln(&pilihan)
			if pilihan == "1" {
				fmt.Println("\nDaftar Kuis:")
				for judulKuis := range s.JawabanKuis {
					fmt.Printf("- %s\n", judulKuis)
				}

				fmt.Print("Masukkan judul kuis yang ingin Anda nilai: ")
				var judulKuis string
				fmt.Scanln(&judulKuis)

				fmt.Print("Masukkan nilai: ")
				var nilaiStr string
				fmt.Scanln(&nilaiStr)

				nilai, _ := strconv.Atoi(nilaiStr)

				if guru.Penilaian[usernameSiswa] == nil {
					guru.Penilaian[usernameSiswa] = make(map[string]int)
				}

				guru.Penilaian[usernameSiswa][judulKuis] = nilai
				fmt.Printf("Penilaian untuk %s pada kuis %s berhasil diberikan!\n", usernameSiswa, judulKuis)
			}

			return
		}
	}

	fmt.Println("Siswa tidak ditemukan.")
}

// LihatMateri memungkinkan siswa untuk melihat materi.
func (lms *LMS) LihatMateri(pengguna *Pengguna) {
	fmt.Println("\nMateri Tersedia:")
	for judul := range lms.Materi {
		fmt.Printf("- %s\n", judul)
	}

	fmt.Print("Masukkan judul materi yang ingin Anda lihat: ")
	var judul string
	fmt.Scanln(&judul)

	if konten, ok := lms.Materi[judul]; ok {
		fmt.Println("\nKonten Materi:")
		fmt.Println(konten)
		return
	}

	fmt.Println("Materi tidak ditemukan.")
}

// KerjakanKuis memungkinkan siswa untuk mengerjakan kuis.
func (lms *LMS) KerjakanKuis(pengguna *Pengguna) {
	fmt.Println("\nKuis Tersedia:")
	for judul := range lms.Kuis {
		fmt.Printf("- %s\n", judul)
	}

	fmt.Print("Masukkan judul kuis yang ingin Anda kerjakan: ")
	var judul string
	fmt.Scanln(&judul)

	if pertanyaan, ok := lms.Kuis[judul]; ok {
		fmt.Println("\nPertanyaan Kuis:")
		var jawaban []string
		for i, pertanyaan := range pertanyaan {
			fmt.Printf("%d. %s\n", i+1, pertanyaan)
			fmt.Print("Jawaban Anda: ")
			var jwb string
			fmt.Scanln(&jwb)
			jawaban = append(jawaban, jwb)
		}
		pengguna.JawabanKuis[judul] = jawaban

		fmt.Println("Jawaban kuis berhasil dikumpulkan!")
		return
	}

	fmt.Println("Kuis tidak ditemukan.")
}

// IkutForum memungkinkan siswa untuk berpartisipasi dalam forum.
func (lms *LMS) IkutForum(pengguna *Pengguna) {
	fmt.Println("\nForum Tersedia:")
	for judul := range lms.Forum {
		fmt.Printf("- %s\n", judul)
	}

	fmt.Print("Masukkan judul forum yang ingin Anda ikuti: ")
	var judul string
	fmt.Scanln(&judul)

	if postingan, ok := lms.Forum[judul]; ok {
		fmt.Println("\nPostingan Forum:")
		for _, post := range postingan {
			fmt.Println(post)
		}

		fmt.Print("Masukkan postingan Anda (atau ketik 'keluar' untuk kembali): ")
		var post string
		fmt.Scanln(&post)
		if post == "keluar" {
			return
		}

		lms.Forum[judul] = append(postingan, post)
		pengguna.Postingan[judul] = append(pengguna.Postingan[judul], post)

		fmt.Println("Postingan berhasil ditambahkan!")
		return
	}

	fmt.Println("Forum tidak ditemukan.")
}

func main() {
	lms := LMS{
		Pengguna: make(map[string]interface{}),
		Materi:   make(map[string]string),
		Kuis:     make(map[string][]string),
		Forum:    make(map[string][]string),
	}

	for {
		fmt.Println("\nSelamat datang di LMS Universitas Telkom!")
		fmt.Println("1. Guru")
		fmt.Println("2. Siswa")
		fmt.Println("3. Keluar")

		fmt.Print("Masukkan pilihan Anda: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			var peran string = "guru"
			for {
				fmt.Print("Masuk atau Daftar (M/D)? ")
				var aksi string
				fmt.Scanln(&aksi)

				if aksi == "M" {
					fmt.Print("Masukkan username: ")
					var username string
					fmt.Scanln(&username)

					fmt.Print("Masukkan password: ")
					var password string
					fmt.Scanln(&password)

					pengguna := lms.Masuk(username, password)
					if pengguna != nil {
						lms.MenuGuru(pengguna)
					}
					break
				} else if aksi == "D" {
					fmt.Print("Masukkan username: ")
					var username string
					fmt.Scanln(&username)

					fmt.Print("Masukkan password: ")
					var password string
					fmt.Scanln(&password)

					lms.Daftar(username, password, peran)
					break
				}
				fmt.Println("Aksi tidak valid.")
			}
		case "2":
			var peran string = "siswa"
			for {
				fmt.Print("Masuk atau Daftar (M/D)? ")
				var aksi string
				fmt.Scanln(&aksi)

				if aksi == "M" {
					fmt.Print("Masukkan username: ")
					var username string
					fmt.Scanln(&username)

					fmt.Print("Masukkan password: ")
					var password string
					fmt.Scanln(&password)

					pengguna := lms.Masuk(username, password)
					if pengguna != nil {
						if p, ok := pengguna.(*Pengguna); ok {
							lms.MenuSiswa(p)
						}
					}
					break
				} else if aksi == "D" {
					fmt.Print("Masukkan username: ")
					var username string
					fmt.Scanln(&username)

					fmt.Print("Masukkan password: ")
					var password string
					fmt.Scanln(&password)

					lms.Daftar(username, password, peran)
					break
				}
				fmt.Println("Aksi tidak valid.")
			}
		case "3":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}