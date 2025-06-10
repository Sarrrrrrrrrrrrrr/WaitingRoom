package team

import (
	"esport-app/model/data"
	"esport-app/model/klasemen"
	"fmt"
)

// MenuStatistikPerMatch menampilkan statistik per match
// dengan meminta input ID match dari user
func MenuStatistikPerMatch() {
	if len(data.MatchList) == 0 {
		fmt.Println("Belum ada match yang tercatat.")
		return
	}

	// Tampilkan daftar match yang tersedia
	fmt.Println("Daftar Match:")
	for _, m := range data.MatchList {
		t1 := data.GetTeamByID(m.Team1ID)
		t2 := data.GetTeamByID(m.Team2ID)
		fmt.Printf("%d. %s vs %s\n", m.ID, t1.NamaTim, t2.NamaTim)
	}

	// Minta input ID match dari user
	var id int
	fmt.Print("Pilih ID Match untuk lihat statistik: ")
	fmt.Scan(&id)

	// Cari match berdasarkan ID yang diberikan
	match := data.GetMatchByID(id)
	// Jika match tidak ditemukan, tampilkan pesan
	if match == nil {
		fmt.Println("❌ Match tidak ditemukan.")
		return
	}

	// Tampilkan statistik lengkap (pakai fungsi dari klasemen)
	klasemen.StatisticsMatch(*match)
}

// PlayerStats menampilkan statistik pemain berdasarkan match yang sudah ada
// dengan meminta input ID match dari user
func PlayerStats() {
	var matchId int
	fmt.Print("Masukkan match ke berapa: ")
	fmt.Scan(&matchId)

	// Mencari match berdasarkan ID yang diberikan
	team := data.GetMatchByID(matchId)
	// Jika match tidak ditemukan, tampilkan pesan
	if team == nil {
		fmt.Println("❌ Tim tidak ditemukan.")
		return
	}

	// Tampilkan statistik lengkap (pakai fungsi dari klasemen)
	klasemen.StatisticsMatch(*team)
}
