package main // Mendeklarasikan paket utama


import ( // Mengimpor paket yang dibutuhkan
    "fmt"  // Mengimpor paket fmt untuk format I/O
    "sort" // Mengimpor paket sort untuk pengurutan slice
    "time" // Mengimpor paket time untuk menangani waktu
)


const jumlahSlotMaks = 10 // Mendefinisikan kapasitas maksimum slot parkir


type Kendaraan struct { // Mendefinisikan struktur untuk data kendaraan
    PlatNomor  string    // Nomor plat kendaraan
    Jenis      string    // Jenis kendaraan (mobil/motor)
    WaktuMasuk time.Time // Waktu masuk kendaraan
}


var slotParkir [jumlahSlotMaks]Kendaraan // Mendeklarasikan array untuk slot parkir dengan kapasitas maksimum


// Data admin
var admins = map[string]string{ // Mendeklarasikan map untuk menyimpan username dan password admin
    "abiyyu": "abiyyu04",     // Admin dengan username "abiyyu" dan password "abiyyu04"
    "andi":   "andicemungut", // Admin dengan username "andi" dan password "andicemungut"
}


// Menjalankan fungsi utama
func main() {
    // Melakukan loop tak terbatas untuk menampilkan menu utama
    for {
        // Menampilkan menu utama
        fmt.Println("Selamat Datang di Aplikasi Parkir")
        fmt.Println("====================================")
        fmt.Println("1. Menu User")
        fmt.Println("2. Menu Admin")
        fmt.Println("3. Keluar")
        fmt.Println("====================================")


        // Meminta user untuk memasukkan pilihan menu
        var pilihanMenuUtama int
        fmt.Print("Masukkan pilihan menu: ")
        fmt.Scan(&pilihanMenuUtama)


        // Melakukan switch case berdasarkan pilihan menu utama
        switch pilihanMenuUtama {
        case 1: // Jika user memilih menu 1 (Menu User)
            menuUser()
        case 2: // Jika user memilih menu 2 (Menu Admin)
            if authenticateAdmin() { // Memeriksa autentikasi admin
                menuAdmin()
            } else {
                fmt.Println("Autentikasi gagal. Username atau password salah.")
            }
        case 3: // Jika user memilih menu 3 (Keluar)
            fmt.Println("Terima kasih telah menggunakan aplikasi parkir.")
            return
        default: // Jika user memasukkan pilihan yang tidak valid
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}


// Fungsi untuk autentikasi admin
func authenticateAdmin() bool {
    var username, password string // Mendeklarasikan variabel untuk username dan password
    fmt.Print("Masukkan username: ")
    fmt.Scan(&username) // Meminta user memasukkan username
    fmt.Print("Masukkan password: ")
    fmt.Scan(&password) // Meminta user memasukkan password


    // Memeriksa apakah username dan password sesuai dengan data admin
    if pass, exists := admins[username]; exists && pass == password { //Mengecek username dan password menggunakan boolean
        return true //Mengembalikan jika true
    }
    return false // Mengem balikan jika false
}


// Fungsi untuk menampilkan menu user
func menuUser() {
    for {
        // Menampilkan menu user
        fmt.Println("====================================")
        fmt.Println("Menu User:")
        fmt.Println("1. Parkir Kendaraan")
        fmt.Println("2. Keluarkan Kendaraan")
        fmt.Println("3. Mencari Lahan Parkir Yang Kosong")
        fmt.Println("4. Kembali ke Menu Utama")
        fmt.Println("====================================")


        var pilihanUser int // Mendeklarasikan variabel untuk pilihan user
        fmt.Print("Masukkan pilihan menu user: ")
        fmt.Scan(&pilihanUser) // Meminta user memasukkan pilihan menu user


        // Melakukan switch case berdasarkan pilihan menu user
        switch pilihanUser {
        case 1:
            parkirKendaraan() // Memanggil fungsi untuk parkir kendaraan
        case 2:
            keluarkanKendaraan() // Memanggil fungsi untuk mengeluarkan kendaraan
        case 3:
            daftarSlotKosong() // Memanggil fungsi untuk menampilkan daftar slot kosong
        case 4:
            return // Kembali ke menu utama
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.") // Menampilkan pesan jika pilihan tidak valid
        }
    }
}


// Fungsi untuk menampilkan menu admin
func menuAdmin() {
    for {
        // Menampilkan menu admin
        fmt.Println("====================================")
        fmt.Println("Menu Admin:")
        fmt.Println("1. Tampilkan Status Parkir")
        fmt.Println("2. Urutkan Slot Parkir (Ascending)")
        fmt.Println("3. Tambah Data Parkir")
        fmt.Println("4. Hapus Data Parkir")
        fmt.Println("5. Ubah Data Parkir")
        fmt.Println("6. Total Pendapatan")
        fmt.Println("7. Kembali ke Menu Utama")
        fmt.Println("====================================")


        var pilihanAdmin int // Mendeklarasikan variabel untuk pilihan admin
        fmt.Print("Masukkan pilihan menu admin: ")
        fmt.Scan(&pilihanAdmin) // Meminta admin memasukkan pilihan menu admin


        // Melakukan switch case berdasarkan pilihan menu admin
        switch pilihanAdmin {
        case 1:
            tampilkanStatusParkir() // Memanggil fungsi untuk menampilkan status parkir
        case 2:
            urutkanSlotParkirAscending() // Memanggil fungsi untuk mengurutkan slot parkir secara ascending
            tampilkanStatusParkir()      // Menampilkan status parkir setelah pengurutan
        case 3:
            tambahDataParkir() // Memanggil fungsi untuk menambah data parkir
        case 4:
            hapusDataParkir() // Memanggil fungsi untuk menghapus data parkir
        case 5:
            ubahDataParkir() // Memanggil fungsi untuk mengubah data parkir
        case 6:
            totalPendapatan() // Memanggil fungsi untuk menampilkan total pendapatan
        case 7:
            return // Kembali ke menu utama
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.") // Menampilkan pesan jika pilihan tidak valid
        }
    }
}


// Fungsi untuk mengurutkan slot parkir secara ascending
func urutkanSlotParkirAscending() {
    // Menyimpan data yang tidak kosong
    // Deklarasi slice untuk menyimpan kendaraan yang slot parkirnya tidak kosong
    var slotsNonKosong []Kendaraan


    // Iterasi melalui semua slot parkir
    for i := 0; i < jumlahSlotMaks; i++ {
        // Jika slot parkir tidak kosong (ada kendaraan)
        if slotParkir[i].PlatNomor != "" {
            // Tambahkan kendaraan ke slice slotsNonKosong
            slotsNonKosong = append(slotsNonKosong, slotParkir[i])
        }
    }


    // Mengurutkan slot yang tidak kosong secara ascending berdasarkan nomor plat
    sort.Slice(slotsNonKosong, func(i, j int) bool {
        // Membandingkan plat nomor untuk menentukan urutan
        return slotsNonKosong[i].PlatNomor < slotsNonKosong[j].PlatNomor
    })


    // Mengosongkan semua slot
    for i := 0; i < jumlahSlotMaks; i++ {
        // Mengatur slot parkir menjadi kosong
        slotParkir[i] = Kendaraan{}
    }


    // Mengisi kembali slot dengan data yang telah diurutkan
    for i, kendaraan := range slotsNonKosong {
        // Menempatkan kendaraan yang sudah diurutkan ke slot parkir
        slotParkir[i] = kendaraan
    }


    // Menampilkan pesan sukses setelah slot parkir berhasil diurutkan
    fmt.Println("Slot Parkir berhasil diurutkan secara Ascending.")
}


// Fungsi untuk menampilkan status parkir
func tampilkanStatusParkir() {
    fmt.Println("Status Parkir:") // Menampilkan header status parkir
    for i := 0; i < jumlahSlotMaks; i++ {
        if slotParkir[i].PlatNomor == "" {
            fmt.Printf("Slot %d: Kosong\n", i+1) // Menampilkan slot kosong
        } else {
            fmt.Printf("Slot %d: %s (%s)\n", i+1, slotParkir[i].PlatNomor, slotParkir[i].Jenis) // Menampilkan slot terisi
        }
    }
}


// Fungsi untuk menghitung total pendapatan
func totalPendapatan() {
    totalMobil := 0 // Mendeklarasikan variabel untuk total pendapatan mobil
    totalMotor := 0 // Mendeklarasikan variabel untuk total pendapatan motor


    // Loop melalui semua slot parkir untuk menghitung total pendapatan
    for i := 0; i < jumlahSlotMaks; i++ {
        // Jika slot tidak kosong, lakukan perhitungan biaya
        if slotParkir[i].PlatNomor != "" {
            // Hitung biaya parkir berdasarkan waktu masuk dan jenis kendaraan
            biaya := hitungBiaya(slotParkir[i].WaktuMasuk, time.Now(), slotParkir[i].Jenis)
            // Jika jenis kendaraan adalah mobil, tambahkan biaya ke total pendapatan mobil
            if slotParkir[i].Jenis == "mobil" {
                totalMobil += biaya
                // Jika jenis kendaraan adalah motor, tambahkan biaya ke total pendapatan motor
            } else if slotParkir[i].Jenis == "motor" {
                totalMotor += biaya


            }
        }
    }


    fmt.Printf("Total Pendapatan Parkir Mobil: Rp%d\n", totalMobil) // Menampilkan total pendapatan parkir mobil
    fmt.Printf("Total Pendapatan Parkir Motor: Rp%d\n", totalMotor) // Menampilkan total pendapatan parkir motor
    fmt.Printf("Total Pendapatan: Rp%d\n", totalMobil+totalMotor)   // Menampilkan total pendapatan keseluruhan
}


// Fungsi untuk parkir kendaraan
func parkirKendaraan() {
    slotKosong := cariSlotKosong() // Mencari slot kosong
    if slotKosong == -1 {
        fmt.Println("Parkir penuh! Pilih slot lain.") // Menampilkan pesan jika parkir penuh
        return
    }


    var platNomor string
    var jenis string
    fmt.Print("Masukkan nomor plat kendaraan: ")
    fmt.Scan(&platNomor) // Meminta user memasukkan nomor plat kendaraan
    fmt.Print("Jenis Kendaraan (mobil/motor): ")
    fmt.Scan(&jenis) // Meminta user memasukkan jenis kendaraan
    kendaraan := Kendaraan{
        PlatNomor:  platNomor,
        Jenis:      jenis,
        WaktuMasuk: time.Now(), // Menyimpan waktu masuk kendaraan
    }
    slotParkir[slotKosong] = kendaraan // Menyimpan data kendaraan ke slot parkir
    fmt.Printf("Kendaraan dengan nomor plat %s diparkir di slot %d\n", platNomor, slotKosong+1)
}


// Fungsi untuk mengeluarkan kendaraan
func keluarkanKendaraan() {
    var platNomor string
    fmt.Print("Masukkan nomor plat kendaraan yang ingin dikeluarkan: ")
    fmt.Scan(&platNomor) // Meminta user memasukkan nomor plat kendaraan yang ingin dikeluarkan


    // Mengurutkan slot parkir sebelum melakukan pencarian
    urutkanSlotParkirAscending()


    index := -1 // Mendeklarasikan variabel untuk menyimpan indeks slot
    for i := 0; i < jumlahSlotMaks; i++ {
        if slotParkir[i].PlatNomor == platNomor {
            index = i
            break
        }
    }


    if index == -1 {
        fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.") // Menampilkan pesan jika kendaraan tidak ditemukan
        return
    }


    waktuMasuk := slotParkir[index].WaktuMasuk
    waktuKeluar := time.Now()                                              // Mendapatkan waktu keluar kendaraan
    biaya := hitungBiaya(waktuMasuk, waktuKeluar, slotParkir[index].Jenis) // Menghitung biaya parkir
    durasiParkir := waktuKeluar.Sub(waktuMasuk)                            // Menghitung durasi parkir
    fmt.Printf("Kendaraan dengan nomor plat %s dikeluarkan dari slot %d.\n", slotParkir[index].PlatNomor, index+1)
    fmt.Printf("Waktu Masuk: %s\n", waktuMasuk.Format("15:04:05"))   // Menampilkan waktu masuk
    fmt.Printf("Waktu Keluar: %s\n", waktuKeluar.Format("15:04:05")) // Menampilkan waktu keluar
    fmt.Printf("Durasi Parkir: %s\n", durasiParkir)                  // Menampilkan durasi parkir
    fmt.Printf("Biaya parkir: Rp%d\n", biaya)                        // Menampilkan biaya parkir
    slotParkir[index] = Kendaraan{}                                  // Mengosongkan slot parkir
}


// Fungsi untuk menghitung biaya parkir
func hitungBiaya(waktuMasuk, waktuKeluar time.Time, jenis string) int {
    const tarifAwalMotor = 2000 // Tarif awal parkir motor
    const tarifAwalMobil = 5000 // Tarif awal parkir mobil
    const tarifPerJam = 1000    // Tarif per jam tambahan


    durasi := waktuKeluar.Sub(waktuMasuk) // Menghitung durasi parkir
    jam := int(durasi.Hours())            // Menghitung jumlah jam
    if durasi.Minutes() > float64(jam*60) {
        jam++ // Menambah satu jam jika ada sisa menit
    }


    var biaya int
    if jenis == "mobil" {
        biaya = tarifAwalMobil // Mengatur biaya awal untuk mobil
        if jam > 1 {
            biaya += (jam - 1) * tarifPerJam // Menambahkan biaya tambahan per jam
        }
    } else if jenis == "motor" {
        biaya = tarifAwalMotor // Mengatur biaya awal untuk motor
        if jam > 1 {
            biaya += (jam - 1) * tarifPerJam // Menambahkan biaya tambahan per jam
        }
    }
    return biaya
}


// Fungsi untuk menampilkan daftar slot kosong
func daftarSlotKosong() {
    // Menampilkan header slot kosong
    fmt.Println("Slot Kosong yang Tersedia:")


    // Mendeklarasikan variabel untuk memeriksa apakah ada slot kosong
    kosong := false


    // Iterasi melalui semua slot parkir
    for i := 0; i < jumlahSlotMaks; i++ {
        // Jika slot parkir kosong (tidak ada plat nomor)
        if slotParkir[i].PlatNomor == "" {
            // Menampilkan nomor slot yang kosong
            fmt.Printf("Slot %d\n", i+1)
            // Menandai bahwa ada slot kosong
            kosong = true
        }
    }


    // Jika tidak ada slot kosong
    if !kosong {
        // Menampilkan pesan bahwa tidak ada slot kosong yang tersedia
        fmt.Println("Tidak ada slot kosong yang tersedia.")
    }
}




// Fungsi untuk menambah data parkir
func tambahDataParkir() {
    // Mencari slot kosong
    slotKosong := cariSlotKosong()
   
    // Jika tidak ada slot kosong
    if slotKosong == -1 {
        // Menampilkan pesan bahwa parkir penuh
        fmt.Println("Parkir penuh! Tidak ada slot kosong.")
        return
    }


    // Mendeklarasikan variabel untuk menyimpan nomor plat dan jenis kendaraan
    var platNomor string
    var jenis string
   
    // Meminta admin memasukkan nomor plat kendaraan
    fmt.Print("Masukkan nomor plat kendaraan: ")
    fmt.Scan(&platNomor)
   
    // Meminta admin memasukkan jenis kendaraan
    fmt.Print("Jenis Kendaraan (mobil/motor): ")
    fmt.Scan(&jenis)
   
    // Membuat objek kendaraan baru dengan data yang dimasukkan
    kendaraan := Kendaraan{
        PlatNomor:  platNomor,
        Jenis:      jenis,
        WaktuMasuk: time.Now(), // Menyimpan waktu masuk kendaraan
    }
   
    // Menyimpan data kendaraan ke slot parkir yang kosong
    slotParkir[slotKosong] = kendaraan
   
    // Menampilkan pesan bahwa kendaraan berhasil diparkir di slot tertentu
    fmt.Printf("Kendaraan dengan nomor plat %s diparkir di slot %d\n", platNomor, slotKosong+1)
}


// Fungsi untuk menghapus data parkir
func hapusDataParkir() {
    // Mendeklarasikan variabel untuk menyimpan nomor plat kendaraan yang ingin dihapus
    var platNomor string
   
    // Meminta admin memasukkan nomor plat kendaraan yang ingin dihapus
    fmt.Print("Masukkan nomor plat kendaraan yang ingin dihapus: ")
    fmt.Scan(&platNomor)


    // Mendeklarasikan variabel untuk memeriksa apakah kendaraan ditemukan
    found := false


    // Iterasi melalui semua slot parkir
    for i := 0; i < jumlahSlotMaks; i++ {
        // Jika slot parkir memiliki kendaraan dengan nomor plat yang dimasukkan
        if slotParkir[i].PlatNomor == platNomor {
            // Menandai bahwa kendaraan ditemukan
            found = true


            // Menampilkan pesan bahwa kendaraan berhasil dihapus dari slot keberapa
            fmt.Printf("Kendaraan dengan nomor plat %s berhasil dihapus dari slot %d.\n", slotParkir[i].PlatNomor, i+1)


            // Mengosongkan slot parkir
            slotParkir[i] = Kendaraan{}
           
            // Mengakhiri loop karena kendaraan sudah ditemukan dan dihapus
            break
        }
    }


    // Jika kendaraan tidak ditemukan, menampilkan pesan bahwa kendaraan tidak ditemukan
    if !found {
        fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.")
    }
}




// Fungsi untuk mengubah data parkir
func ubahDataParkir() {
    // Mendeklarasikan variabel untuk menyimpan nomor plat kendaraan yang ingin diubah
    var platNomor string


    // Meminta admin memasukkan nomor plat kendaraan yang ingin diubah
    fmt.Print("Masukkan nomor plat kendaraan yang ingin diubah: ")
    fmt.Scan(&platNomor)


    // Mendeklarasikan variabel untuk memeriksa apakah kendaraan ditemukan
    found := false


    // Iterasi melalui semua slot parkir
    for i := 0; i < jumlahSlotMaks; i++ {
        // Jika slot parkir memiliki kendaraan dengan nomor plat yang dimasukkan
        if slotParkir[i].PlatNomor == platNomor {
            // Menandai bahwa kendaraan ditemukan
            found = true


            // Menampilkan pesan bahwa kendaraan ditemukan dan di slot keberapa
            fmt.Printf("Kendaraan dengan nomor plat %s ditemukan di slot %d.\n", slotParkir[i].PlatNomor, i+1)


            // Meminta admin memasukkan data baru
            fmt.Println("Masukkan data baru:")


            // Mendeklarasikan variabel untuk menyimpan nomor plat baru dan jenis kendaraan baru
            var platNomorBaru string // Mendeklarasikan variabel platNomorBaru adalah string
            var jenisBaru string     // Mendeklarasikan variabel jenisBaru adalah string


            // Meminta admin memasukkan nomor plat baru
            fmt.Print("Nomor Plat Kendaraan baru: ")
            fmt.Scan(&platNomorBaru)


            // Meminta admin memasukkan jenis kendaraan baru
            fmt.Print("Jenis Kendaraan baru (mobil/motor): ")
            fmt.Scan(&jenisBaru)


            // Mengubah nomor plat kendaraan pada slot parkir
            slotParkir[i].PlatNomor = platNomorBaru


            // Mengubah jenis kendaraan pada slot parkir
            slotParkir[i].Jenis = jenisBaru


            // Menampilkan pesan bahwa data kendaraan berhasil diubah
            fmt.Printf("Data kendaraan dengan nomor plat %s berhasil diubah menjadi nomor plat %s dan jenis %s.\n", platNomor, platNomorBaru, jenisBaru)


            // Mengakhiri loop karena kendaraan sudah ditemukan dan datanya diubah
            break
        }
    }


    // Jika kendaraan tidak ditemukan, menampilkan pesan bahwa kendaraan tidak ditemukan
    if !found {
        fmt.Println("Kendaraan dengan nomor plat tersebut tidak ditemukan.")
    }
}


// Fungsi untuk mencari slot parkir kosong
func cariSlotKosong() int {
    // Iterasi melalui semua slot parkir
    for i := 0; i < jumlahSlotMaks; i++ {
        // Jika slot parkir kosong (tidak ada plat nomor)
        if slotParkir[i].PlatNomor == "" {
            // Mengembalikan indeks slot kosong
            return i
        }
    }
    // Mengembalikan -1 jika tidak ada slot kosong
    return -1
}
