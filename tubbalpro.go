package main

import "fmt"

type RiwayatSurvey struct {
    Tanggal string
    Jawaban [9]int
}

type Akun struct {
    Username string
    Email    string
    Password string
    Riwayat  []RiwayatSurvey
}

var daftarAkun []Akun
var adminPassword = "admin123"

func isDigit(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] < '0' || s[i] > '9' {
            return false
        }
    }
    return true
}

var phd9Questions = [9]string{
    "1. Dalam 1 minggu terakhir seberapa sering Anda merasa sedikit tertarik atau senang melakukan sesuatu?",
    "2. Dalam 1 minggu terakhir seberapa sering Anda merasa murung, tertekan, atau putus asa?",
    "3. Dalam 1 minggu terakhir seberapa sering Anda mengalami kesulitan tidur (terlalu banyak/tidak bisa tidur)?",
    "4. Dalam 1 minggu terakhir seberapa sering Anda merasa lelah atau tidak memiliki energi?",
    "5. Dalam 1 minggu terakhir seberapa sering Anda kehilangan nafsu makan atau makan terlalu banyak?",
    "6. Dalam 1 minggu terakhir seberapa sering Anda merasa buruk tentang diri sendiri (seolah-olah gagal atau mengecewakan keluarga)?",
    "7. Dalam 1 minggu terakhir seberapa sering Anda kesulitan berkonsentrasi dalam hal-hal seperti membaca koran atau menonton TV?",
    "8. Dalam 1 minggu terakhir seberapa sering Anda bergerak atau berbicara sangat lambat atau sebaliknya—sangat gelisah dan tidak bisa diam?",
    "9. Dalam 1 minggu terakhir seberapa sering Anda berpikir bahwa lebih baik mati atau menyakiti diri sendiri?",
}

func analisaHasilSurvey(jawaban [9]int) (int, string) {
    total := 0
    for _, nilai := range jawaban {
        total += nilai
    }
    rekomendasi := ""
    if total >= 9 && total <= 15 {
        rekomendasi = "Sangat Rendah: Tidak menunjukkan tanda-tanda depresi klinis. Tidak diperlukan intervensi, tetapi tetap dimonitor jika ada gejala ringan."
    } else if total >= 16 && total <= 23 {
        rekomendasi = "Rendah: Gejala ringan, pantau secara berkala. Pertimbangkan konseling ringan atau pendekatan psikososial."
    } else if total >= 24 && total <= 31 {
        rekomendasi = "Sedang: Kemungkinan depresi klinis. Evaluasi lebih lanjut disarankan. Bisa dipertimbangkan terapi psikologis (misalnya CBT)."
    } else if total >= 32 && total <= 39 {
        rekomendasi = "Tinggi: Depresi cukup berat. Perlu evaluasi klinis menyeluruh. Terapi psikologis dan/atau farmakologis biasanya disarankan."
    } else if total >= 40 && total <= 45 {
        rekomendasi = "Sangat Tinggi: Depresi berat. Memerlukan penanganan intensif segera, bisa melibatkan psikiater, farmakoterapi, dan terapi psikologis intensif."
    }
    return total, rekomendasi
}

func binarySearchUsername(username string) *Akun {
    left, right := 0, len(daftarAkun)-1
    for left <= right {
        mid := left + (right-left)/2
        if daftarAkun[mid].Username == username {
            return &daftarAkun[mid]
        } else if daftarAkun[mid].Username < username {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return nil
}

func main() {
    for {
        var email, password string
        var isAdmin, isLoggedIn bool
        var loggedInAkun *Akun
        var input string
        var inputValid bool
        fmt.Println("\n=== MENU AWAL ===")
        fmt.Println("Apakah Anda sudah memiliki akun? (Ya/Tidak)")
        inputValid = false
        for !inputValid {
            fmt.Scanln(&input)
            if input == "Tidak" {
                var newUsername, newEmail, newPassword string
                fmt.Print("Masukkan username baru: ")
                fmt.Scanln(&newUsername)
                fmt.Print("Masukkan email baru: ")
                emailValid := false
                for !emailValid {
                    fmt.Scanln(&newEmail)
                    if len(newEmail) > 10 && newEmail[len(newEmail)-10:] == "@gmail.com" {
                        emailValid = true
                    } else {
                        fmt.Println("Email tidak valid. Harus diakhiri dengan @gmail.com")
                        fmt.Print("Masukkan email baru: ")
                    }
                }
                fmt.Print("Masukkan password baru: ")
                fmt.Scanln(&newPassword)
                newAkun := Akun{Username: newUsername, Email: newEmail, Password: newPassword}
                daftarAkun = append(daftarAkun, newAkun)
                fmt.Println("Akun berhasil dibuat! Silakan login.")
                inputValid = true
            } else if input == "Ya" {
                inputValid = true
            } else {
                fmt.Println("Jawaban hanya Ya/Tidak!")
                fmt.Println("Apakah Anda sudah memiliki akun? (Ya/Tidak)")
                inputValid = false
            }
        }
        for !isLoggedIn {
            fmt.Print("Masukkan email: ")
            fmt.Scanln(&email)
            fmt.Print("Masukkan password: ")
            fmt.Scanln(&password)
            if email == "admin@admin.com" && password == adminPassword {
                isAdmin = true
                isLoggedIn = true
            } else {
                for i := 0; i < len(daftarAkun); i++ {
                    if daftarAkun[i].Email == email && daftarAkun[i].Password == password {
                        loggedInAkun = &daftarAkun[i]
                        isLoggedIn = true
                    }
                }
                if !isLoggedIn {
                    fmt.Println("Email atau password salah. Silakan coba lagi.")
                }
            }
        }
        if isAdmin {
            fmt.Println("\n=== MENU ADMIN ===")
            for i := 0; i < len(daftarAkun); i++ {
                fmt.Printf("\nUser  %d:\n", i+1)
                fmt.Println("Username:", daftarAkun[i].Username)
                fmt.Println("Email:", daftarAkun[i].Email)
                if len(daftarAkun[i].Riwayat) > 0 {
                    fmt.Println("Riwayat Survey:")
                    for _, r := range daftarAkun[i].Riwayat {
                        fmt.Println("Tanggal:", r.Tanggal, "Jawaban:", r.Jawaban)
                    }
                } else {
                    fmt.Println("Belum pernah mengisi survey.")
                }
            }
        } else {
            selesai := false
            for !selesai {
                fmt.Println("\n=== MENU USER ===")
                fmt.Println("1. Isi survey PHQ-9")
                fmt.Println("2. Lihat riwayat survey")
                fmt.Println("3. Ubah survey")
                fmt.Println("4. Hapus survey")
                fmt.Println("5. Cari hasil survey berdasarkan Username")
                fmt.Println("6. Logout / Kembali ke menu awal")
                var pilihan int
                fmt.Print("Pilih menu (1-6): ")
                fmt.Scanln(&pilihan)

                if pilihan == 1 {
                    var tanggal string
                    var jawaban [9]int
                    valid := false
                    for !valid {
                        fmt.Print("Masukkan tanggal (format DD-MM-YYYY): ")
                        fmt.Scanln(&tanggal)
                        if len(tanggal) == 10 && tanggal[2] == '-' && tanggal[5] == '-' {
                            dd := tanggal[0:2]
                            mm := tanggal[3:5]
                            yyyy := tanggal[6:10]
                            if isDigit(dd) && isDigit(mm) && isDigit(yyyy) {
                                valid = true
                            } else {
                                fmt.Println("Format salah. Gunakan angka dengan format DD-MM-YYYY.")
                            }
                        } else {
                            fmt.Println("Format salah. Gunakan format DD-MM-YYYY.")
                        }
                    }
                    for i := 0; i < 9; i++ {
                        jawabanValid := false
                        for !jawabanValid {
                            fmt.Println("")
                            fmt.Println("------------------Isi Survei PHQ-9 (Jawaban 1-5)------------------")
                            fmt.Println("")
                            fmt.Println("1. Tidak Sesuai")
                            fmt.Println("2. Sedikit Sesuai")
                            fmt.Println("3. Sesuai")
                            fmt.Println("4. Lumayan Sesuai")
                            fmt.Println("5. Sangat Sesuai")
                            fmt.Println("")
                            fmt.Println("-------------Silahkan Isi Jawaban Sesuai Kondisi Anda-------------")
                            fmt.Println("")
                            fmt.Println(phd9Questions[i])
                            fmt.Println("")
                            fmt.Scanln(&jawaban[i])
                            fmt.Println("")
                            if jawaban[i] >= 1 && jawaban[i] <= 5 {
                                jawabanValid = true
                            } else {
                                fmt.Println("******************************************************************")
                                fmt.Println("-------Input tidak valid. Masukkan angka antara 1 sampai 5.-------")
                                fmt.Println("******************************************************************")
                                fmt.Println("")
                                fmt.Printf("------------------ Pertanyaan ke-%d diulang yahh ------------------", i+1)
                            }
                        }
                    }
                    loggedInAkun.Riwayat = append(loggedInAkun.Riwayat, RiwayatSurvey{
                        Tanggal: tanggal,
                        Jawaban: jawaban,
                    })
                    fmt.Println("Survey berhasil disimpan.")
                }
                if pilihan == 2 {
                    if len(loggedInAkun.Riwayat) == 0 {
                        fmt.Println("Belum ada riwayat survey.")
                    } else {
                        fmt.Println("Riwayat Pengisian Survey:")
                        for _, r := range loggedInAkun.Riwayat {
                            total, rekomendasi := analisaHasilSurvey(r.Jawaban)
                            fmt.Println("Tanggal:", r.Tanggal)
                            fmt.Println("Jawaban:", r.Jawaban)
                            fmt.Println("Skor:", total)
                            fmt.Println("Rekomendasi:", rekomendasi)
                            fmt.Println()
                        }
                    }
                }

                if pilihan == 3 {
                    if len(loggedInAkun.Riwayat) == 0 {
                        fmt.Println("Belum ada data survey untuk diubah.")
                    } else {
                        var index int
                        fmt.Println("Pilih nomor survey yang ingin diubah:")
                        for i, r := range loggedInAkun.Riwayat {
                            fmt.Printf("%d. Tanggal: %s\n", i+1, r.Tanggal)
                        }
                        fmt.Scanln(&index)
                        if index > 0 && index <= len(loggedInAkun.Riwayat) {
                            index--
                            var jawaban [9]int
                            for i := 0; i < 9; i++ {
                                jawabanValid := false
                                for !jawabanValid {
                                    fmt.Println("")
                                    fmt.Println("------------------Isi Survei PHQ-9 (Jawaban 1-5)------------------")
                                    fmt.Println("")
                                    fmt.Println("1. Tidak Sesuai")
                                    fmt.Println("2. Sedikit Sesuai")
                                    fmt.Println("3. Sesuai")
                                    fmt.Println("4. Lumayan Sesuai")
                                    fmt.Println("5. Sangat Sesuai")
                                    fmt.Println("")
                                    fmt.Println("-------------Silahkan Isi Jawaban Sesuai Kondisi Anda-------------")
                                    fmt.Println("")
                                    fmt.Println(phd9Questions[i])
                                    fmt.Println("")
                                    fmt.Scanln(&jawaban[i])
                                    fmt.Println("")
                                    if jawaban[i] >= 1 && jawaban[i] <= 5 {
                                        jawabanValid = true
                                    } else {
                                        fmt.Println("******************************************************************")
                                        fmt.Println("-------Input tidak valid. Masukkan angka antara 1 sampai 5.-------")
                                        fmt.Println("******************************************************************")
                                        fmt.Println("")
                                        fmt.Printf("------------------ Pertanyaan ke-%d diulang yahh ------------------", i+1)
                                    }
                                }
                            }
                            loggedInAkun.Riwayat[index].Jawaban = jawaban
                            fmt.Println("Survey berhasil diubah.")
                        } else {
                            fmt.Println("Pilihan tidak valid.")
                        }
                    }
                }
                if pilihan == 4 {
                    if len(loggedInAkun.Riwayat) == 0 {
                        fmt.Println("Belum ada data survey untuk dihapus.")
                    } else {
                        var index int
                        fmt.Println("Pilih nomor survey yang ingin dihapus:")
                        for i, r := range loggedInAkun.Riwayat {
                            fmt.Printf("%d. Tanggal: %s\n", i+1, r.Tanggal)
                        }
                        fmt.Scanln(&index)
                        if index > 0 && index <= len(loggedInAkun.Riwayat) {
                            index--
                            loggedInAkun.Riwayat = append(loggedInAkun.Riwayat[:index], loggedInAkun.Riwayat[index+1:]...)
                            fmt.Println("Survey berhasil dihapus.")
                        } else {
                            fmt.Println("Pilihan tidak valid.")
                        }
                    }
                }
                if pilihan == 5 {
                    fmt.Print("Masukkan Username yang ingin dicari: ")
                    var cariUsername string
                    fmt.Scanln(&cariUsername)

                    akunDitemukan := binarySearchUsername(cariUsername)

                    if akunDitemukan != nil {
                        fmt.Println("Akun ditemukan:")
                        fmt.Println("Username:", akunDitemukan.Username)
                        fmt.Println("Email:", akunDitemukan.Email)
                        if len(akunDitemukan.Riwayat) == 0 {
                            fmt.Println("Belum ada data survey.")
                        } else {
                            for _, r := range akunDitemukan.Riwayat {
                                fmt.Println("Tanggal:", r.Tanggal)
                                fmt.Println("Jawaban:", r.Jawaban)
                                total, rekomendasi := analisaHasilSurvey(r.Jawaban)
                                fmt.Println("Skor:", total)
                                fmt.Println("Rekomendasi:", rekomendasi)
                                fmt.Println()
                            }
                        }
                    } else {
                        fmt.Println("Akun dengan username tersebut tidak ditemukan.")
                    }
                }
                if pilihan == 6 {
                    selesai = true
                    fmt.Println("Logout berhasil. Kembali ke menu awal.")
                }
            }
        }
    }
}


