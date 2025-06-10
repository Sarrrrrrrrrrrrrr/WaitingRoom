package panitia

import (
	"esport-app/model/data"
	"esport-app/model/klasemen"
	"fmt"
)

// TambahMatch menambahkan match baru ke daftar match
// dengan meminta input ID tim, skor, dan KDA pemain
func TambahMatch() {
	var t1ID, t2ID, s1, s2 int
	fmt.Print("Masukkan ID Tim 1: ")
	fmt.Scan(&t1ID)
	fmt.Print("Masukkan ID Tim 2: ")
	fmt.Scan(&t2ID)

	// Validasi ID tim tidak boleh sama
	// Jika ID tim sama, tampilkan pesan error dan kembali ke menu
	if t1ID == t2ID {
		fmt.Println("❌ ID Tim tidak boleh sama. Coba lagi.")
		return
	}

	fmt.Print("Masukkan Skor Tim 1: ")
	fmt.Scan(&s1)
	fmt.Print("Masukkan Skor Tim 2: ")
	fmt.Scan(&s2)

	// Validasi ID tim yang dimasukkan harus ada dalam daftar tim
	// Jika ID tim tidak ditemukan, tampilkan pesan error dan kembali ke menu
	team1 := data.GetTeamByID(t1ID)
	team2 := data.GetTeamByID(t2ID)

	if team1 == nil || team2 == nil {
		fmt.Println("❌ Salah satu ID Tim tidak ditemukan.")
		return
	}

	// membuat slice untuk menyimpan KDA pemain dari masing-masing tim
	var kda1 = make([]string, len(team1.Players))
	var kda2 = make([]string, len(team2.Players))

	fmt.Println("Masukkan KDA untuk Tim 1 (format kill/death/assist):")

	// Meminta input KDA untuk setiap pemain di Tim 1
	for i := 0; i < len(team1.Players); i++ {
		fmt.Printf("Player %d: ", i+1)
		fmt.Scan(&kda1[i])
	}

	fmt.Println("Masukkan KDA untuk Tim 2 (format kill/death/assist):")

	// Meminta input KDA untuk setiap pemain di Tim 2
	for i := 0; i < len(team2.Players); i++ {
		fmt.Printf("Player %d: ", i+1)
		fmt.Scan(&kda2[i])
	}

	// Validasi KDA yang dimasukkan
	// membuat id baru untuk match yang akan ditambahkan
	newID := data.GetLastMatchID()

	newMatch := data.Match{
		ID:               newID,
		Team1ID:          t1ID,
		Team2ID:          t2ID,
		Team1Score:       s1,
		Team2Score:       s2,
		PlayerKDAByTeam1: kda1,
		PlayerKDAByTeam2: kda2,
	}

	// menyimpan match baru ke dalam daftar match
	data.MatchList = append(data.MatchList, newMatch)

	// Update statistik tim
	klasemen.UpdateTeamStats(team1, s1, s2, kda1)
	klasemen.UpdateTeamStats(team2, s2, s1, kda2)

	// Simpan data match dan tim ke file
	klasemen.SaveMatchesToFile()
	klasemen.SaveTeamsToFile()

	fmt.Println("✅ Match berhasil ditambahkan!")
}

func ShowMatch() {
	if len(data.MatchList) == 0 {
		fmt.Println("Belum ada match yang tercatat.")
		return
	}
	fmt.Println("📋 Daftar Semua Match:")
	for _, m := range data.MatchList {
		t1 := data.GetTeamByID(m.Team1ID)
		t2 := data.GetTeamByID(m.Team2ID)

		fmt.Printf("Match %d: %s (%d) vs %s (%d)\n", m.ID, t1.NamaTim, m.Team1Score, t2.NamaTim, m.Team2Score)
	}
}

// AverageKDA returns the average KDA for a player.
func AverageKDA(p *data.Player) string {
	if p.TotalMatch == 0 {
		return "0.0"
	}
	// Menghitung rata-rata KDA
	avg := float64(p.TotalKill+p.TotalAssist) / float64(p.TotalDeath+1)
	return fmt.Sprintf("%.2f", avg)
}

// UpdateMatch mengupdate match yang sudah ada
// dengan meminta input ID match yang ingin diupdate
func UpdateMatch() {
	var id, choice int
	fmt.Print("Masukkan ID Match yang ingin diupdate: ")
	fmt.Scan(&id)

	// Mencari match berdasarkan ID yang diberikan
	match := data.GetMatchByID(id)
	if match == nil {
		fmt.Println("❌ Match tidak ditemukan.")
		return
	}

	fmt.Println("1. Ubah Skor")
	fmt.Println("2. Ubah KDA")

	fmt.Print("Pilih aksi: ")
	fmt.Scan(&choice)

	// memilih aksi yang ingin dilakukan
	if choice == 1 {
		var newS1, newS2 int
		fmt.Print("Masukkan Skor Tim 1 baru: ")
		fmt.Scan(&newS1)
		fmt.Print("Masukkan Skor Tim 2 baru: ")
		fmt.Scan(&newS2)

		// mengupdate skor match
		match.Team1Score = newS1
		match.Team2Score = newS2

		klasemen.SaveMatchesToFile()

		fmt.Println("✅ Skor match berhasil diperbarui!")
	} else if choice == 2 {
		// mengambil tim berdasarkan ID yang ada di match
		team1 := data.GetTeamByID(match.Team1ID)
		team2 := data.GetTeamByID(match.Team2ID)

		// Validasi tim yang ditemukan
		var newKDA1 = make([]string, len(team1.Players))
		var newKDA2 = make([]string, len(team2.Players))

		// Memainta input KDA baru untuk setiap pemain di Tim 1 dan Tim 2
		fmt.Println("Masukkan KDA baru untuk Tim 1:")
		for i := 0; i < len(team1.Players); i++ {
			fmt.Printf("Player %d: ", i+1)
			fmt.Scan(&newKDA1[i])
		}
		fmt.Println("Masukkan KDA baru untuk Tim 2:")
		for i := 0; i < len(team2.Players); i++ {
			fmt.Printf("Player %d: ", i+1)
			fmt.Scan(&newKDA2[i])
		}

		// Validasi KDA yang dimasukkan
		match.PlayerKDAByTeam1 = newKDA1
		match.PlayerKDAByTeam2 = newKDA2

		// Update statistik tim berdasarkan KDA baru
		klasemen.SaveMatchesToFile()

		fmt.Println("✅ KDA match berhasil diperbarui!")
	} else {
		fmt.Println("❌ Pilihan tidak valid.")
	}
}

// DeleteMatch menghapus match berdasarkan ID yang diberikan
// Ini akan menghapus match dari daftar match dan menyimpan perubahan ke file
func DeleteMatch() {
	var id int
	fmt.Print("Masukkan ID Match yang ingin dihapus: ")
	fmt.Scan(&id)

	// Mencari match berdasarkan ID yang diberikan
	for i, m := range data.MatchList {
		if m.ID == id {
			// Menghapus match dari daftar match
			// Ini akan menghapus match dari slice MatchList
			data.MatchList = append(data.MatchList[:i], data.MatchList[i+1:]...)
			klasemen.SaveMatchesToFile()
			fmt.Println("✅ Match berhasil dihapus!")
			return
		}
	}
	fmt.Println("❌ Match tidak ditemukan.")
}
