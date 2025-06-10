package controllers

import (
	"esport-app/handlers/team"
	"esport-app/helpers"
	"fmt"
)

// PanitiaMenu menampilkan menu untuk panitia
// yang berisi pilihan untuk melihat klasemen, mencari tim, cari pertandingan, dll
// Menu ini akan terus ditampilkan sampai user memilih untuk keluar
func PanitiaMenu() {
	var choice int

	for isRunning := false; !isRunning; {
		helpers.ClearScreen()
		helpers.PanitiaMenu()

		// Meminta input pilihan dari user
		for choice < 1 || choice > 7 {
			fmt.Print("Pilih menu : ")
			fmt.Scan(&choice)
		}

		// Menjalankan fungsi sesuai dengan pilihan user
		// Pilihan 1 sampai 6 akan menjalankan fungsi yang sesuai
		// Pilihan 7 akan keluar dari menu

		switch choice {
		case 1:
			KlasemenMenu()
		case 2:
			CariNamaTim()
			helpers.ConfirmationScreen()
		case 3:
			CariPertandingan()
			helpers.ConfirmationScreen()
		case 4:
			TampilkanPeringkatTertinggi()
			helpers.ConfirmationScreen()
		case 5:
			TampilkanPeringkatTerendah()
			helpers.ConfirmationScreen()
		case 6:
			team.PlayerStats()
			helpers.ConfirmationScreen()
		case 7:
			isRunning = true
		}

		// Jika pilihan tidak 7 (keluar), reset pilihan ke 0
		// agar menu panitia bisa ditampilkan kembali
		// Ini juga untuk menghindari infinite loop jika input tidak valid
		if choice != 7 {
			choice = 0
		}
	}
}
