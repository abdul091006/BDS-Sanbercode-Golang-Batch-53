package main

import "fmt"

// Soal 1
func luas(phi float64, r float64, luasLingkaran *float64) {
	*luasLingkaran = phi  * r * r
}
func keliling(phi float64, r float64, kelilingLingkaran *float64) {
	*kelilingLingkaran = phi  * (r + r)
}

// Soal 2
func introduce(sentence *string, name string, gender string, job string, age string){
	if gender == "laki-laki"{
		*sentence = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else if gender == "perempuan"{
		*sentence = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else {
		*sentence = name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}
}

// Soal 3
func buahFavorite(buah *[]string, buahFav ...string){
	*buah = append(*buah, buahFav...)
	for i, data := range *buah{
		fmt.Printf("%d. %s\n", i+1, data)
	}
}

// Soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string){
	*dataFilm = append(*dataFilm, map[string]string{
		"title": title,
		"duration": duration,
		"genre": genre,
		"year": year,
	})
}

func main() {
	// Soal 1
	var luasLingkaran float64 
	var kelilingLingkaran float64
	luas(3.14, 7, &luasLingkaran)
	keliling(3.14, 7, &kelilingLingkaran)
	fmt.Println("Luas:", luasLingkaran)
	fmt.Println("Keliling:", kelilingLingkaran)

	// Soal 2
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	var buah = []string{}
	buahFavorite(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	// Soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, data := range dataFilm{
		title := data["title"]
		duration := data["duration"]
		genre := data["genre"]
		year := data["year"]

		fmt.Printf("%d. title : %s\n   duration : %s\n   genre : %s\n   year : %s\n\n", i+1, title, duration, genre, year)
	} 
}