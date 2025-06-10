package database

import (
	"fmt"
	"os"
)

const storageFile = "./storage"

// function init digunakan untuk memastikan direktori penyimpanan ada
// jika tidak ada, maka akan dibuatkan direktori baru dengan permission 0755
func init() {
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		os.Mkdir(storageFile, 0755)
	}
}

// FileExists mengecek apakah file dengan nama filename ada di direktori penyimpanan
// mengembalikan true jika file ada, false jika tidak ada
func FileExists(filename string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", storageFile, filename))
	return !os.IsNotExist(err)
}

// CreateFile membuat file baru dengan nama filename dan data yang diberikan
func UpdateFile(filename string, data []byte) error {
	fullPath := fmt.Sprintf("%s/%s", storageFile, filename)
	return os.WriteFile(fullPath, data, 0644)
}

// ReadFile membaca isi file dengan nama filename dari direktori penyimpanan
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", storageFile, filename))
}
