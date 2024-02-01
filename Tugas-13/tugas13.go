package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama		string	`json:"nama"`
	MataKuliah	string	`json:"mata_kuliah"`
	IndeksNilai string	`json:"indeks_nilai"`
	Nilai		uint 	`json:"nilai"`
	ID          uint	`json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Nilai NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Nilai); err != nil {
				log.Fatal(err)
			}
			if Nilai.Nilai > 100 {
				response := []byte(`{"message": "Data nilai maksimal 100"}`)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			switch {
			case Nilai.Nilai >= 80:
				Nilai.IndeksNilai = "A"
			case Nilai.Nilai >= 70 && Nilai.Nilai < 80:
				Nilai.IndeksNilai = "B"
			case Nilai.Nilai >= 60 && Nilai.Nilai < 70:
				Nilai.IndeksNilai = "C"
			case Nilai.Nilai >= 50 && Nilai.Nilai < 60:
				Nilai.IndeksNilai = "D"
			case Nilai.Nilai < 80:
				Nilai.IndeksNilai = "E"
			}
			Nilai.ID = uint(rand.Intn(1000))
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Nilai)
			
		} else {
			// parse dari form
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mata_kuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			Nilai = NilaiMahasiswa{
				Nama:       nama,
				MataKuliah: mataKuliah,
				Nilai:      uint(nilai),
			}
			if Nilai.Nilai > 100 {
				response := []byte(`{"message": "Data nilai maksimal 100"}`)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			switch {
			case Nilai.Nilai >= 80:
				Nilai.IndeksNilai = "A"
			case Nilai.Nilai >= 70 && Nilai.Nilai < 80:
				Nilai.IndeksNilai = "B"
			case Nilai.Nilai >= 60 && Nilai.Nilai < 70:
				Nilai.IndeksNilai = "C"
			case Nilai.Nilai >= 50 && Nilai.Nilai < 60:
				Nilai.IndeksNilai = "D"
			case Nilai.Nilai < 80:
				Nilai.IndeksNilai = "E"
			}
			Nilai.ID = uint(rand.Intn(1000))
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Nilai)
		}

		dataNilai, _ := json.Marshal(Nilai)
		w.Write(dataNilai)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func GetNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.Handle("/", Auth(http.HandlerFunc(GetNilai)))
	http.HandleFunc("/post", PostNilai)

	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
