package controllers

import (
	"esport-app/handlers/panitia"
	"esport-app/handlers/team"
	"esport-app/helpers"
	"fmt"
)

// KlasemenMenu menampilkan menu untuk panitia
// yang berisi pilihan untuk melihat klasemen, menambah, mengupdate, menghapus pertandingan, dan melihat statistik pemain
// Menu ini akan terus ditampilkan sampai user memilih untuk keluar
func KlasemenMenu() {
	var choice int

	for isRunning := false; !isRunning; {
		helpers.ClearScreen()
		helpers.KlasemenMenu()

		// Meminta input pilihan dari user
		for choice < 1 || choice > 6 {
			fmt.Print("Pilih menu : ")
			fmt.Scan(&choice)
		}

		// Menjalankan fungsi sesuai dengan pilihan user
		// Pilihan 1 sampai 5 akan menjalankan fungsi yang sesuai
		// Pilihan 6 akan keluar dari menu
		switch choice {
		case 1:
			panitia.ShowMatch()
			helpers.ConfirmationScreen()
		case 2:
			panitia.TambahMatch()
			helpers.ConfirmationScreen()
		case 3:
			panitia.UpdateMatch()
			helpers.ConfirmationScreen()

		case 4:
			panitia.DeleteMatch()
			helpers.ConfirmationScreen()

		case 5:
			team.PlayerStats()
			helpers.ConfirmationScreen()

		case 6:
			isRunning = true
		}

		// Jika pilihan tidak 6 (keluar), reset pilihan ke 0
		// agar menu klasemen bisa ditampilkan kembali
		// Ini juga untuk menghindari infinite loop jika input tidak valid
		if choice != 6 {
			choice = 0
		}
	}
}
