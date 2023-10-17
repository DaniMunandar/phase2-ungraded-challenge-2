package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Menghubungkan ke database MySQL
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/avengers")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat HTTP handler untuk endpoint "pahlawan"
	http.HandleFunc("/pahlawan", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT name, universe, skill, imageURL FROM Heroes")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var heroes []map[string]string

		for rows.Next() {
			var name, universe, skill, imageURL string
			if err := rows.Scan(&name, &universe, &skill, &imageURL); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			hero := map[string]string{
				"name":     name,
				"universe": universe,
				"skill":    skill,
				"imageURL": imageURL,
			}
			heroes = append(heroes, hero)
		}

		// Mengirim daftar pahlawan sebagai JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(heroes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Membuat HTTP handler untuk endpoint "Villains"
	http.HandleFunc("/Villains", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT name, universe, imageURL FROM Villain")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var villains []map[string]string

		for rows.Next() {
			var name, universe, imageURL string
			if err := rows.Scan(&name, &universe, &imageURL); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			villain := map[string]string{
				"name":     name,
				"universe": universe,
				"imageURL": imageURL,
			}
			villains = append(villains, villain)
		}

		// Mengirim daftar musuh sebagai JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(villains); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Menjalankan server pada port 8080
	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http://localhost:8080/pahlawan
// http://localhost:8080/Villains
