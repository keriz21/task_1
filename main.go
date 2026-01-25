package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Categories struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Description string `json:"description"`
}

var kategori = []Categories{
	{ID: 1, Nama: "Indomie Goreng", Description: "Mie instan goreng rasa original"},
	{ID: 2, Nama: "Vit 1000ml", Description: "Air mineral botol 1 liter"},
	{ID: 3, Nama: "Kecap Manis", Description: "Kecap manis botol 600ml"},
	{ID: 4, Nama: "Susu UHT", Description: "Susu UHT coklat 200ml"},
	{ID: 5, Nama: "Roti Tawar", Description: "Roti tawar tanpa kulit 10 slice"},
	{ID: 6, Nama: "Teh Botol", Description: "Teh manis dalam botol 350ml"},
	{ID: 7, Nama: "Beras 5kg", Description: "Beras premium kemasan 5kg"},
}

func getCategories(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func getCategoriesByID(w http.ResponseWriter, r * http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w,"Invalid Categories ID", http.StatusBadRequest)
		return
	}

	for _,p := range kategori {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Kategori Beluum ada", http.StatusNotFound)
}

func postCategories(w http.ResponseWriter, r *http.Request){
	var kategoriBaru Categories

	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	kategoriBaru.ID = len(kategori) + 1
	kategori = append(kategori, kategoriBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(kategoriBaru)
}

func updateCategoriesByID(w http.ResponseWriter, r * http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w,"Invalid Categories ID", http.StatusBadRequest)
		return
	}

	var updateKategori Categories
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range kategori {
		if kategori[i].ID == id {
			updateKategori.ID = id
			kategori[i] = updateKategori

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateKategori)
			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func deleteCategoriesByID(w http.ResponseWriter, r * http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w,"Invalid Categories ID", http.StatusBadRequest)
		return
	}

	for i,p := range kategori {
		if p.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]... )
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message" : "sukses menghapus kategori",
			})

			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func main(){
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "API Running",
	})
})

http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCategories(w,r)
	} else if r.Method == "POST" {
		postCategories(w,r)
	}
})

http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCategoriesByID(w,r)
	} else if r.Method == "PUT" {
		updateCategoriesByID(w,r)
	} else if r.Method == "DELETE" {
		deleteCategoriesByID(w,r)
	}
})

err:=http.ListenAndServe(":8080",nil)

if err != nil{
	fmt.Print(("gagal running server"))
}
}



