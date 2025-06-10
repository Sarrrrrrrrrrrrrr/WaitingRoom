package main

import (
	"esport-app/controllers"
	"esport-app/handlers/auth"
	"esport-app/helpers"
	"fmt"
)

// Func main itu adalah program utama yang akan dijalankan
// yang akan menampilkan menu autentikasi dan kemudian masuk ke menu sesuai dengan role user
func main() {
	var pilihan int

	// berfungsi untuk membersihkan layar terminal
	helpers.ClearScreen()

	// Looping untuk menampilkan menu autentikasi sampai user berhasil login
	for isLogged := false; !isLogged; {
		helpers.ClearScreen()

		// Menampilkan menu autentikasi
		helpers.AuthMenu()

		// Meminta input pilihan dari user
		for pilihan < 1 || pilihan > 3 {
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pilihan)
		}

		helpers.ClearScreen()

		// Memanggil fungsi sesuai dengan input pilihan user
		switch pilihan {
		case 1:
			isLogged = controllers.Login()
		case 2:
			controllers.Register()
		case 3:
			isLogged = true
			return
		}

		// Jika pilihan tidak 3 (keluar), reset pilihan ke 0
		// agar menu autentikasi bisa ditampilkan kembali
		// Ini juga untuk menghindari infinite loop jika input tidak valid
		if pilihan != 3 {
			pilihan = 0
		}
	}

	helpers.ClearScreen()

	// Setelah user berhasil login, program akan masuk ke menu sesuai dengan role user
	// Looping untuk menampilkan menu sesuai dengan role user
	for isRunning := false; !isRunning; {
		// mendapatkan role user yang sudah login
		roleUser := auth.GetRoleUser()

		if roleUser == "panitia" {
			controllers.PanitiaMenu()
			break
		} else {
			controllers.PlayerMenu()
			break
		}
	}

	fmt.Println("Selamat Tinggal")
}
