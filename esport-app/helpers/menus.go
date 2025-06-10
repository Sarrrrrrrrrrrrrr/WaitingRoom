package helpers

import (
	"fmt"
)

// AuthMenu menampilkan menu autentikasi
func AuthMenu() {
	fmt.Println("==== Auth ====")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")
}

// ConfirmationScreen menampilkan pesan konfirmasi untuk kembali ke halaman sebelumnya
func ConfirmationScreen() {
	var back string
	fmt.Println("Kembali ke halaman sebelumnya ? (y)")
	fmt.Scan(&back)

	for back == "y" {
		return
	}
}

// PanitiaMenu menampilkan menu untuk panitia
func PanitiaMenu() {
	fmt.Println("==== Crew Menu ====")
	fmt.Println("1. Klasemen menu")
	fmt.Println("2. Cari Tim berdasarkan nama")
	fmt.Println("3. Cari Tim berdasarkan dua tim")
	fmt.Println("4. Urutkan tim dari skor tertinggi")
	fmt.Println("5. Urutkan tim dari skor terendah")
	fmt.Println("6. Lihat statistik peforma")
	fmt.Println("7. Exit")
}

// PlayerMenu menampilkan menu untuk pemain
func PlayerMenu() {
	fmt.Println("==== Player Menu ====")
	fmt.Println("1. Lihat Klasmen")
	fmt.Println("2. Cari Tim berdasarkan nama")
	fmt.Println("3. Cari Tim berdasarkan dua tim")
	fmt.Println("4. Urutkan tim dari skor tertinggi")
	fmt.Println("5. Urutkan tim dari skor terendah")
	fmt.Println("6. Lihat statistik peforma")
	fmt.Println("7. Exit")
}

// KlasemenMenu menampilkan menu untuk panitia
func KlasemenMenu() {
	fmt.Println("=== MENU Klasemen ===")
	fmt.Println("1. Tampilkan Semua Data")
	fmt.Println("2. Tambah Match")
	fmt.Println("3. Edit Match")
	fmt.Println("4. Hapus Match")
	fmt.Println("5. Keluar")
}

// ClearScreen membersihkan layar terminal
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
