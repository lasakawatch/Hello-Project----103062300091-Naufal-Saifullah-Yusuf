package main

import (
	"fmt"
)

type Pengguna struct {
	Username    string
	Password    string
	Peran       string
	JawabanKuis map[string][]string // Kunci: Judul Kuis, Nilai: Jawaban
	Postingan   map[string][]string // Kunci: Judul Forum, Nilai: Postingan
}

type LMS struct {
	Pengguna map[string]*Pengguna
	Materi   map[string]string   // Kunci: Judul Materi, Nilai: Konten
	Kuis     map[string][]string // Kunci: Judul Kuis, Nilai: Pertanyaan
	Forum    map[string][]string // Kunci: Judul Forum, Nilai: Postingan
}

func (lms *LMS) Daftar(username, password, peran string) {
	if _, ok := lms.Pengguna[username]; !ok {
		lms.Pengguna[username] = &Pengguna{
			Username:    username,
			Password:    password,
			Peran:       peran,
			JawabanKuis: make(map[string][]string),
			Postingan:   make(map[string][]string),
		}
		fmt.Printf("Pengguna %s berhasil terdaftar!\n", username)
	} else {
		fmt.Println("Username sudah ada.")
	}
}

func (lms *LMS) Masuk(username, password string) *Pengguna {
	if pengguna, ok := lms.Pengguna[username]; ok {
		if pengguna.Password == password {
			fmt.Printf("Selamat datang %s!\n", username)
			return pengguna
		} else {
			fmt.Println("Password salah.")
		}
	} else {
		fmt.Println("Pengguna tidak ditemukan. Silakan daftar terlebih dahulu.")
	}
	return nil
}

func (lms *LMS) MenuGuru(pengguna *Pengguna) {
	for {
		fmt.Println("\nMenu Guru:")
		fmt.Println("1. Buat Materi")
		fmt.Println("2. Buat Kuis")
		fmt.Println("3. Buat Forum")
		fmt.Println("4. Data Siswa")
		fmt.Println("5. Keluar")

		fmt.Print("Masukkan pilihan Anda: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			lms.BuatMateri(pengguna)
		case "2":
			lms.BuatKuis(pengguna)
		case "3":
			lms.BuatForum(pengguna)
		case "4":
			lms.DataSiswa(pengguna)
		case "5":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

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

// Metode Guru
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

func (lms *LMS) BuatForum(pengguna *Pengguna) {
	fmt.Print("Masukkan judul forum: ")
	var judul string
	fmt.Scanln(&judul)

	lms.Forum[judul] = []string{} // Inisialisasi dengan daftar postingan kosong
	fmt.Println("Forum berhasil dibuat!")
}

func (lms *LMS) DataSiswa(pengguna *Pengguna) {
	fmt.Println("\nSiswa Terdaftar:")
	for username, siswa := range lms.Pengguna {
		if siswa.Peran == "siswa" {
			fmt.Printf("- %s\n", username)
		}
	}

	fmt.Print("Masukkan username siswa yang ingin Anda lihat: ")
	var usernameSiswa string
	fmt.Scanln(&usernameSiswa)

	if siswa, ok := lms.Pengguna[usernameSiswa]; ok {
		fmt.Printf("\nSiswa: %s\n", usernameSiswa)

		// Tampilkan Jawaban Kuis
		if len(siswa.JawabanKuis) > 0 {
			fmt.Println("\nJawaban Kuis:")
			for judulKuis, jawaban := range siswa.JawabanKuis {
				fmt.Printf("  - %s:\n", judulKuis)
				for i, jwb := range jawaban {
					fmt.Printf("    %d. %s\n", i+1, jwb)
				}
			}
		} else {
			fmt.Println("Tidak ada jawaban kuis.")
		}

		// Tampilkan Postingan Forum
		if len(siswa.Postingan) > 0 {
			fmt.Println("\nPostingan Forum:")
			for judulForum, postingan := range siswa.Postingan {
				fmt.Printf("  - %s:\n", judulForum)
				for _, post := range postingan {
					fmt.Printf("    %s\n", post)
				}
			}
		} else {
			fmt.Println("Tidak ada postingan forum.")
		}
	} else {
		fmt.Println("Siswa tidak ditemukan.")
	}
}

// Metode Siswa
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
	} else {
		fmt.Println("Materi tidak ditemukan.")
	}
}

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
		// Simpan jawaban untuk siswa
		pengguna.JawabanKuis[judul] = jawaban
		fmt.Println("Jawaban kuis berhasil dikumpulkan!")
	} else {
		fmt.Println("Kuis tidak ditemukan.")
	}
}

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

		// Simpan postingan untuk siswa
		pengguna.Postingan[judul] = append(pengguna.Postingan[judul], post)

		fmt.Println("Postingan berhasil ditambahkan!")
	} else {
		fmt.Println("Forum tidak ditemukan.")
	}
}

func main() {
	lms := LMS{
		Pengguna: make(map[string]*Pengguna),
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
			peran := "guru"
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
				} else {
					fmt.Println("Aksi tidak valid.")
				}
			}
		case "2":
			peran := "siswa"
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
						lms.MenuSiswa(pengguna)
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
				} else {
					fmt.Println("Aksi tidak valid.")
				}
			}
		case "3":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
