package main

import (
	"fmt"
)

//Soal 1
func luasPersegiPanjang(p int, l int) int{
	return p*l
}

func kelilingPersegiPanjang(p int, l int) int{
	return 2*(p+l)
}

func volumeBalok(p int, l int, t int) int{
	return p*l*t
}

// SOal 2
func introduce(name string, gender string, job string, age string) (sentence string) {
	if gender == "laki-laki"{
		sentence = "Pak" + " " + name + " " + "adalah seorang" + " " + job + " " + "yang berusia" + " " + age + " " + "tahun"
	} else if gender == "perempuan"{
		sentence = "Bu" + " " + name + " " + "adalah seorang" + " " + job + " " + "yang berusia" + " " + age + " " + "tahun"
	} else {
		sentence = name + " " + "adalah seorang" + " " + job + " " + "yang berusia" + " " + age + " " + "tahun"
	}
	return
}

// Soal 3
func buahFavorit(name string, buahFav ...string) string {
	var buah string
	for i, data := range buahFav{
		buah += `"` + data + `"`
		if i < len(buahFav)-1 {
			buah += ", "
		}
	}
	return "halo nama saya" + " " + name + " " + "dan buah favorit saya adalah" + " " + buah
}

func main() {
	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// Soal 4
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func (title, duration, genre, year string) []map[string]string {
		data := map[string]string{
			"genre": genre,
			"jam": duration,
			"tahun": year,
			"title": title,
		}
		dataFilm = append(dataFilm, data)
		return dataFilm
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")
	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
