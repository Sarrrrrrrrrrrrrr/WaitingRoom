// model/data/dataModel.go
package data

import (
	"esport-app/database"
	"fmt"
	"strings"
)

const (
	MatchFile = "match.json"
	TeamFile  = "teams.json"
)

// Player menyimpan data pemain
type Player struct {
	Nama        string
	Role        string
	KDA         string
	TotalKill   int
	TotalDeath  int
	TotalAssist int
	TotalMatch  int
	TeamID      int
}

// Team menyimpan data tim dan pemainnya
type Team struct {
	ID      int
	NamaTim string
	Players []Player
	Win     int
	Lose    int
	Score   int
}

// Match menyimpan data hasil pertandingan
type Match struct {
	ID               int
	Team1ID          int
	Team2ID          int
	Team1Score       int
	Team2Score       int
	PlayerKDAByTeam1 []string
	PlayerKDAByTeam2 []string
}

// Global slices menyimpan data di memori
var (
	TeamList   []Team
	MatchList  []Match
	PlayerList []Player
)

func init() {
	// ─── Inisialisasi MatchList ────────────────────────────────────────────────
	if database.FileExists(MatchFile) {
		// Jika file sudah ada, baca dan load ke MatchList
		content, err := database.ReadFile(MatchFile)
		if err != nil {
			panic(err)
		}
		if err := database.LoadFromJSON(content, &MatchList); err != nil {
			panic(err)
		}
	} else {
		// Jika belum ada, buat kosong dan simpan ke JSON
		MatchList = []Match{}
		data, err := database.SaveToJSON(MatchList)
		if err != nil {
			panic(err)
		}
		if err := database.UpdateFile(MatchFile, data); err != nil {
			panic(err)
		}
	}

	// ─── Inisialisasi PlayerList (default) ──────────────────────────────────────
	PlayerList = []Player{
		{"Branz", "Gold Laner", "0/0/0", 0, 0, 0, 0, 1},
		{"DreamS", "Roamer", "0/0/0", 0, 0, 0, 0, 1},
		{"Cr1te", "Jungler", "0/0/0", 0, 0, 0, 0, 1},
		{"Hijumee", "Mid Laner", "0/0/0", 0, 0, 0, 0, 1},
		{"Fluffy", "EXP Laner", "0/0/0", 0, 0, 0, 0, 1},

		{"Kairi", "Jungler", "0/0/0", 0, 0, 0, 0, 2},
		{"Butsss", "EXP Laner", "0/0/0", 0, 0, 0, 0, 2},
		{"CW", "Gold Laner", "0/0/0", 0, 0, 0, 0, 2},
		{"SamoHT", "Roamer", "0/0/0", 0, 0, 0, 0, 2},
		{"Sanz", "Mid Laner", "0/0/0", 0, 0, 0, 0, 2},

		{"Lemon", "EXP Laner", "0/0/0", 0, 0, 0, 0, 3},
		{"Skylar", "Gold Laner", "0/0/0", 0, 0, 0, 0, 3},
		{"Vyn", "Roamer", "0/0/0", 0, 0, 0, 0, 3},
		{"Albert", "Jungler", "0/0/0", 0, 0, 0, 0, 3},
		{"Clay", "Mid Laner", "0/0/0", 0, 0, 0, 0, 3},

		{"Celiboy", "Jungler", "0/0/0", 0, 0, 0, 0, 4},
		{"Pai", "EXP Laner", "0/0/0", 0, 0, 0, 0, 4},
		{"Udil", "Mid Laner", "0/0/0", 0, 0, 0, 0, 4},
		{"Nino", "Gold Laner", "0/0/0", 0, 0, 0, 0, 4},
		{"Rasy", "Roamer", "0/0/0", 0, 0, 0, 0, 4},
	}

	// ─── Inisialisasi TeamList dan sinkronisasi PlayerList ───────────────────────
	if database.FileExists(TeamFile) {
		// Jika file teams.json sudah ada, load ke TeamList
		content, err := database.ReadFile(TeamFile)
		if err != nil {
			panic(err)
		}
		if err := database.LoadFromJSON(content, &TeamList); err != nil {
			panic(err)
		}
	} else {
		// Jika belum ada, buat default dan simpan
		TeamList = []Team{
			{ID: 1, NamaTim: "Evos"},
			{ID: 2, NamaTim: "Onic"},
			{ID: 3, NamaTim: "RRQ"},
			{ID: 4, NamaTim: "Alter"},
		}
		// Gabungkan pemain ke tim
		for i := range TeamList {
			for _, p := range PlayerList {
				if p.TeamID == TeamList[i].ID {
					TeamList[i].Players = append(TeamList[i].Players, p)
				}
			}
		}
		saveTeamsToFile()
	}
}

// Ambil Match berdasarkan ID
func GetMatchByID(id int) *Match {
	for i := range MatchList {
		if MatchList[i].ID == id {
			return &MatchList[i]
		}
	}
	return nil
}

// Ambil Team berdasarkan ID
func GetTeamByID(id int) *Team {
	for i := range TeamList {
		if TeamList[i].ID == id {
			return &TeamList[i]
		}
	}
	return nil
}

// Ambil Team berdasarkan nama (case-insensitive)
func GetTeamByName(nama string) *Team {
	for i := range TeamList {
		if strings.EqualFold(TeamList[i].NamaTim, nama) {
			return &TeamList[i]
		}
	}
	return nil
}

// Dapatkan ID baru untuk match berikutnya (incremental)
func GetLastMatchID() int {
	maxID := 0
	for _, m := range MatchList {
		if m.ID > maxID {
			maxID = m.ID
		}
	}
	return maxID + 1
}

// saveMatchesToFile menyimpan MatchList ke JSON
func SaveMatchesToFile() {
	dataJSON, err := database.SaveToJSON(MatchList)
	if err != nil {
		fmt.Println("Gagal mengonversi MatchList ke JSON:", err)
		return
	}
	if err := database.UpdateFile(MatchFile, dataJSON); err != nil {
		fmt.Println("Gagal menyimpan file", MatchFile, ":", err)
	}
}

// saveTeamsToFile menyimpan TeamList ke JSON
func saveTeamsToFile() {
	dataJSON, err := database.SaveToJSON(TeamList)
	if err != nil {
		fmt.Println("Gagal mengonversi TeamList ke JSON:", err)
		return
	}
	if err := database.UpdateFile(TeamFile, dataJSON); err != nil {
		fmt.Println("Gagal menyimpan file", TeamFile, ":", err)
	}
}
