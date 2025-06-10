package database

import "encoding/json"

// LoadFromJSON memuat data dari JSON ke dalam struct yang diberikan
// v harus berupa pointer ke struct yang sesuai dengan format JSON
func LoadFromJSON(content []byte, v interface{}) error {
	return json.Unmarshal(content, v)
}

// SaveToJSON menyimpan data dari struct ke dalam format JSON
// v harus berupa struct yang sesuai dengan format yang diinginkan
func SaveToJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
