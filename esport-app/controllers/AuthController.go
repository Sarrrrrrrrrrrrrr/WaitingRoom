package controllers

import (
	"esport-app/handlers/auth"
	"esport-app/helpers"
	"fmt"
)

// Function Login digunakan untuk melakukan autentikasi user
// dengan meminta input username dan password dari user
// Jika autentikasi berhasil, akan mengembalikan nilai true
func Login() bool {
	var username, password string
	fmt.Println("=== Login ===")

	fmt.Print("Masukan Username : ")
	fmt.Scan(&username)

	fmt.Print("Masukan Password : ")
	fmt.Scan(&password)

	// Memanggil fungsi Authentic dari package auth untuk melakukan autentikasi
	// mengecek apakah username dan password yang dimasukkan sesuai
	// dan menyimpan status autentikasi ke dalam variabel status
	status := auth.Authentic(username, password)

	return status
}

// Function Register digunakan untuk mendaftarkan user baru
// dengan meminta input username dan password dari user
// Jika pendaftaran berhasil, akan mengembalikan nilai true
// Jika pendaftaran gagal, akan mengembalikan nilai false
func Register() bool {
	var username, password string
	helpers.ClearScreen()

	fmt.Println("=== Register ===")

	fmt.Print("Masukan Username : ")
	fmt.Scan(&username)

	fmt.Print("Masukan Password : ")
	fmt.Scan(&password)

	// Memanggil fungsi Register dari package auth untuk mendaftarkan user baru
	// dan menyimpan hasil pendaftaran ke dalam variabel result
	result, _ := auth.Register(username, password)

	return result
}
