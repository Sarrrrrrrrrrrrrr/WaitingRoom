package controllers

import (
	"esport-app/handlers/panitia"
	"esport-app/handlers/team"
	"esport-app/helpers"
	"esport-app/model/data"
	"esport-app/model/klasemen"
	"fmt"
)

// PlayerMenu menampilkan menu untuk pemain
// yang berisi pilihan untuk melihat pertandingan, mencari tim, mencari pertandingan, menampilkan peringkat tertinggi, peringkat terendah, dan melihat statistik pemain
// Menu ini akan terus ditampilkan sampai user memilih untuk keluar
func PlayerMenu() {
	var choice int

	for isRunning := false; !isRunning; {
		helpers.ClearScreen()
		helpers.PlayerMenu()

		// Meminta input pilihan dari user
		for choice < 1 || choice > 7 {
			fmt.Print("Pilih menu : ")
			fmt.Scan(&choice)
		}

		// Menjalankan fungsi sesuai dengan pilihan user
		// Pilihan 1 sampai 6 akan menjalankan fungsi yang sesuai
		// Pilihan 7 akan keluar dari menu
		// Jika pilihan tidak valid, akan meminta input ulang
		switch choice {
		case 1:
			panitia.ShowMatch()
			helpers.ConfirmationScreen()
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
		// agar menu pemain bisa ditampilkan kembali
		// Ini juga untuk menghindari infinite loop jika input tidak valid
		if choice != 7 {
			choice = 0
		}

	}
}

// TampilkanPeringkatTertinggi menampilkan peringkat tim berdasarkan skor tertinggi menggunakan selection sort
func TampilkanPeringkatTertinggi() {
	fmt.Println("=== Peringkat Tim (Skor Tertinggi - Selection Sort) ===")

	teams := append([]data.Team{}, data.TeamList...) // salin slice dari TeamList yang ada di package data

	// Selection sort
	n := len(teams)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if teams[j].Score > teams[maxIdx].Score {
				maxIdx = j
			}
		}
		teams[i], teams[maxIdx] = teams[maxIdx], teams[i]
	}

	// Menampilkan header tabel
	fmt.Printf("%-3s | %-15s | %4s | %4s | %5s\n", "No", "Nama Tim", "Win", "Lose", "Score")

	// Menampilkan data tim yang sudah diurutkan
	for i, t := range teams {
		fmt.Printf("%-3d | %-15s | %4d | %4d | %5d\n", i+1, t.NamaTim, t.Win, t.Lose, t.Score)
	}
}

// TampilkanPeringkatTerendah menampilkan peringkat tim berdasarkan skor terendah menggunakan insertion sort
func TampilkanPeringkatTerendah() {
	fmt.Println("=== Peringkat Tim (Skor Terendah - Insertion Sort) ===")

	teams := append([]data.Team{}, data.TeamList...) // salin slice dari TeamList yang ada di package data

	// Insertion sort
	for i := 1; i < len(teams); i++ {
		key := teams[i]
		j := i - 1
		for j >= 0 && teams[j].Score > key.Score {
			teams[j+1] = teams[j]
			j--
		}
		teams[j+1] = key
	}

	// Menampilkan header tabel
	fmt.Printf("%-3s | %-15s | %4s | %4s | %5s\n", "No", "Nama Tim", "Win", "Lose", "Score")

	// Menampilkan data tim yang sudah diurutkan
	for i, t := range teams {
		fmt.Printf("%-3d | %-15s | %4d | %4d | %5d\n", i+1, t.NamaTim, t.Win, t.Lose, t.Score)
	}
}

// CariNamaTim mencari tim berdasarkan nama dan menampilkan match yang melibatkan tim tersebut
func CariNamaTim() {
	var nama string
	fmt.Print("Masukkan nama tim: ")
	fmt.Scan(&nama)

	matches := klasemen.SearchMatchByTeamName(nama) // Mencari semua match yang melibatkan tim dengan nama yang diberikan
	// Jika tidak ada match yang ditemukan, tampilkan pesan
	if len(matches) == 0 {
		fmt.Println("Tidak ada match ditemukan.")
		return
	}

	team := data.GetTeamByName(nama) // Mencari tim berdasarkan nama
	// Jika tim tidak ditemukan, tampilkan pesan
	if team == nil {
		fmt.Println("Data tim tidak ditemukan.")
		return
	}

	// Menampilkan informasi tim dan match yang melibatkan tim tersebut
	fmt.Printf("Match yang melibatkan %s (Win: %d, Lose: %d, Score: %d):\n", team.NamaTim, team.Win, team.Lose, team.Score)
	for _, m := range matches {
		// Mencari tim berdasarkan ID yang ada di match
		t1 := data.GetTeamByID(m.Team1ID)
		t2 := data.GetTeamByID(m.Team2ID)

		// menampilkan informasi match
		fmt.Printf("- Match %d: %s (%d) vs %s (%d)\n", m.ID, t1.NamaTim, m.Team1Score, t2.NamaTim, m.Team2Score)
	}
}

// CariPertandingan mencari pertandingan antara dua tim berdasarkan nama tim
func CariPertandingan() {
	var namaA, namaB string
	fmt.Print("Masukkan nama tim A: ")
	fmt.Scan(&namaA)
	fmt.Print("Masukkan nama tim B: ")
	fmt.Scan(&namaB)

	// Mencari match antara dua tim berdasarkan nama yang sudah diinputkan
	match := klasemen.SearchMatchBetweenTeams(namaA, namaB)
	// Jika tidak ada match yang ditemukan, tampilkan pesan
	if match == nil {
		fmt.Println("Match tidak ditemukan.")
		return
	}

	// Mencari tim berdasarkan ID yang ada di match
	t1 := data.GetTeamByID(match.Team1ID)
	t2 := data.GetTeamByID(match.Team2ID)

	// Ini juga untuk menghindari infinite loop jika input tidak valid
	if t1 == nil || t2 == nil {
		fmt.Println("Data tim tidak ditemukan.")
		return
	}

	// Menampilkan informasi match dan tim yang terlibat
	fmt.Printf("Match %d: %s (%d) vs %s (%d)\n", match.ID, t1.NamaTim, match.Team1Score, t2.NamaTim, match.Team2Score)
	fmt.Printf("%s - Win: %d, Lose: %d, Score: %d\n", t1.NamaTim, t1.Win, t1.Lose, t1.Score)
	fmt.Printf("%s - Win: %d, Lose: %d, Score: %d\n", t2.NamaTim, t2.Win, t2.Lose, t2.Score)
}
