package auth

import "esport-app/database"

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

var users []User

// func init digunakan untuk menginisialisasi data user
func init() {
	// jika file user.json sudah ada, maka data user akan dimuat dari file tersebut
	if database.FileExists("user.json") {
		content, err := database.ReadFile("user.json")
		if err != nil {
			panic(err)
		}

		database.LoadFromJSON(content, &users)

	} else {
		// jika file user.json tidak ada, maka data user akan diinisialisasi dengan data default
		users = []User{
			{ID: 1, Username: "admin", Password: "admin123", Role: "panitia"},
			{ID: 2, Username: "player", Password: "player123", Role: "player"},
		}
		content, err := database.SaveToJSON(users)

		if err != nil {
			panic(err)
		}

		err = database.UpdateFile("user.json", content)

		if err != nil {
			panic(err)
		}
	}

}

var currentUser *User

// DataUser mengembalikan slice dari User yang berisi data user
func DataUser() []User {
	return users
}

// Authentic digunakan untuk memeriksa apakah username dan password yang diberikan cocok dengan data user
func Authentic(username, password string) bool {
	for i := range users {
		if users[i].Username == username {
			if users[i].Password == password {
				currentUser = &users[i]
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// GetCurrentUser mengembalikan user yang sedang login
func GetRoleUser() string {
	if currentUser != nil {
		return currentUser.Role
	}
	return ""
}

// Register digunakan untuk mendaftarkan user baru
// Jika username sudah terdaftar, akan mengembalikan false dan pesan error
func Register(username, password string) (bool, string) {
	for _, user := range users {
		if user.Username == username {
			return false, "Username sudah terdaftar"
		}
	}
	// Jika username belum terdaftar, buat user baru
	// ID user baru adalah ID terakhir + 1
	newID := users[len(users)-1].ID + 1
	newUser := User{
		ID:       newID,
		Username: username,
		Password: password,
		Role:     "player",
	}
	// Tambahkan user baru ke slice users
	users = append(users, newUser)

	// Simpan ke file JSON
	content, err := database.SaveToJSON(users)
	if err != nil {
		return false, "Gagal menyimpan data user"
	}
	err = database.UpdateFile("user.json", content)
	if err != nil {
		return false, "Gagal memperbarui file user.json"
	}

	return true, "Registrasi berhasil"
}
