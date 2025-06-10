// model/klasemen/klasemenModel.go
package klasemen

import (
	"fmt"
	"sort"
	"strings"

	"esport-app/database"
	"esport-app/model/data"
)

// Untuk memudahkan referensi file match (gunakaan dari data.MatchFile)
var matchFile = data.MatchFile

func UpdateTeamStats(team *data.Team, scoreTeam, scoreOther int, kdaList []string) {
	// Update statistik tim berdasarkan skor
	if scoreTeam > scoreOther {
		team.Win++
		team.Score += 3
	} else if scoreTeam == scoreOther {
		team.Score += 1
	} else {
		team.Lose++
	}

	// Update statistik pemain berdasarkan KDA
	n := min(min(5, len(team.Players)), len(kdaList))

	// Pastikan jumlah pemain tidak melebihi jumlah KDA yang diberikan
	for i := 0; i < n; i++ {
		var kill, death, assist int
		fmt.Sscanf(kdaList[i], "%d/%d/%d", &kill, &death, &assist)

		player := &team.Players[i]
		player.TotalKill += kill
		player.TotalDeath += death
		player.TotalAssist += assist
		player.TotalMatch++
		player.KDA = fmt.Sprintf("%d/%d/%d", player.TotalKill, player.TotalDeath, player.TotalAssist)

		// Sinkron ke PlayerList
		for j := range data.PlayerList {
			if data.PlayerList[j].Nama == player.Nama && data.PlayerList[j].TeamID == player.TeamID {
				data.PlayerList[j] = *player
				break
			}
		}
	}
}

// SaveMatchesToFile menulis seluruh MatchList (di package data) ke disk
func SaveMatchesToFile() {
	// Ambil slice match dari data
	matches := data.MatchList

	// Konversi ke JSON
	dataJSON, err := database.SaveToJSON(matches)
	if err != nil {
		fmt.Println("Gagal mengonversi MatchList ke JSON:", err)
		return
	}
	// Simpan
	if err := database.UpdateFile(matchFile, dataJSON); err != nil {
		fmt.Println("Gagal menyimpan file", matchFile, ":", err)
	}
}

// SaveTeamsToFile menulis seluruh TeamList (di package data) ke disk
func SaveTeamsToFile() {
	teams := data.TeamList

	dataJSON, err := database.SaveToJSON(teams)
	if err != nil {
		fmt.Println("Gagal mengonversi TeamList ke JSON:", err)
		return
	}
	if err := database.UpdateFile(data.TeamFile, dataJSON); err != nil {
		fmt.Println("Gagal menyimpan file", data.TeamFile, ":", err)
	}
}

// SearchMatchByTeamName mencari semua match yang melibatkan nama tim tertentu (case-insensitive)
func SearchMatchByTeamName(nama string) []data.Match {
	var hasil []data.Match

	// Cari ID tim di data.TeamList
	var teamID int
	found := false
	for _, team := range data.TeamList {
		if strings.EqualFold(team.NamaTim, nama) {
			teamID = team.ID
			found = true
			break
		}
	}
	if !found {
		return hasil
	}

	// Loop di data.MatchList, cari semua match yang ada teamID
	for _, m := range data.MatchList {
		if m.Team1ID == teamID || m.Team2ID == teamID {
			hasil = append(hasil, m)
		}
	}
	return hasil
}

// SearchMatchBetweenTeams mencari satu match antara dua tim (binary search dengan key (minID,maxID))
func SearchMatchBetweenTeams(namaA, namaB string) *data.Match {
	var idA, idB int
	foundA, foundB := false, false

	// Cari ID dari namaA dan namaB
	for _, t := range data.TeamList {
		if strings.EqualFold(t.NamaTim, namaA) {
			idA = t.ID
			foundA = true
		}
		if strings.EqualFold(t.NamaTim, namaB) {
			idB = t.ID
			foundB = true
		}
	}
	if !foundA || !foundB {
		return nil
	}

	// Siapkan slice berisi keyedMatches untuk binary search
	type KeyedMatch struct {
		MinID int
		MaxID int
		Match data.Match
	}
	var keyedMatches []KeyedMatch
	for _, m := range data.MatchList {
		minID := min(m.Team1ID, m.Team2ID)
		maxID := max(m.Team1ID, m.Team2ID)
		keyedMatches = append(keyedMatches, KeyedMatch{
			MinID: minID,
			MaxID: maxID,
			Match: m,
		})
	}
	// Sort berdasarkan (MinID, MaxID)
	sort.Slice(keyedMatches, func(i, j int) bool {
		if keyedMatches[i].MinID != keyedMatches[j].MinID {
			return keyedMatches[i].MinID < keyedMatches[j].MinID
		}
		return keyedMatches[i].MaxID < keyedMatches[j].MaxID
	})

	targetMin := min(idA, idB)
	targetMax := max(idA, idB)
	low, high := 0, len(keyedMatches)-1
	for low <= high {
		mid := (low + high) / 2
		curr := keyedMatches[mid]
		if curr.MinID == targetMin && curr.MaxID == targetMax {
			return &curr.Match
		}
		if curr.MinID < targetMin || (curr.MinID == targetMin && curr.MaxID < targetMax) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func StatisticsMatch(match data.Match) {
	t1Name, t2Name := "", ""
	var team1Players, team2Players []data.Player

	// Ambil nama tim dan data pemain dari TeamList
	for _, t := range data.TeamList {
		if t.ID == match.Team1ID {
			t1Name = t.NamaTim
			team1Players = t.Players
		}
		if t.ID == match.Team2ID {
			t2Name = t.NamaTim
			team2Players = t.Players
		}
	}

	fmt.Printf("Statistik Match %d: %s vs %s\n", match.ID, t1Name, t2Name)
	fmt.Printf("Skor: %d - %d\n\n", match.Team1Score, match.Team2Score)

	fmt.Println("KDA Team 1:")
	for i, kda := range match.PlayerKDAByTeam1 {
		if i < len(team1Players) {
			player := team1Players[i]
			fmt.Printf("Player %d: %s (%s) | KDA: %s\n", i+1, player.Nama, player.Role, kda)
		}
	}

	fmt.Println("\nKDA Team 2:")
	for i, kda := range match.PlayerKDAByTeam2 {
		if i < len(team2Players) {
			player := team2Players[i]
			fmt.Printf("Player %d: %s (%s) | KDA: %s\n", i+1, player.Nama, player.Role, kda)
		}
	}
}

// Pembantu: fungsi min dan max untuk dua int
// min mengembalikan nilai terkecil dari dua bilangan bulat
// max mengembalikan nilai terbesar dari dua bilangan bulat
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
