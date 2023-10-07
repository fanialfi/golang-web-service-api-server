package routing

import (
	"encoding/json"
	"net/http"
)

// digunakan sebagai tipe element sample data
type student struct {
	ID    string
	Name  string
	Grade int
}

// sample data untuk keperluan web service
var data = []student{
	{ID: "E001", Name: "fani", Grade: 17},
	{ID: "E002", Name: "ethan", Grade: 21},
	{ID: "E003", Name: "wick", Grade: 22},
	{ID: "E004", Name: "Bourne", Grade: 23},
	{ID: "E005", Name: "Bond", Grade: 24},
}

// users function, digunakan untuk handle endpoint `/users`
func Users(w http.ResponseWriter, r *http.Request) {
	// menentukan tipe response
	w.Header().Set("Content-Type", "application/json")

	// r.Method digunakan untuk mendeteksi jenis request method
	if r.Method == "GET" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// digunakan untuk mengirimkan data ke user sebagai bentuk response
		w.Write(result)
		return
	}
	// jika sebuah request tidak valid, response di set sebagai error
	http.Error(w, "", http.StatusBadRequest)
	return

}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// r.FormValue digunakan untuk mengambil nilai parameter yang dikirim oleh client dalam bentuk url
		id := r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
			http.Error(w, "User Not Found", http.StatusNotFound)
			return // keluar dari function
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
