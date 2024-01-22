package main

import (
	"fmt"
)

// Soal 1
type buah struct{
	nama string
	warna string
	adaBijinya bool
	harga int
}

// Soal 2
type segitiga struct{
	alas, tinggi int
}
  
type persegi struct{
	sisi int
}
  
type persegiPanjang struct{
	panjang, lebar int
}

func (l segitiga) luasSegitiga(){
	fmt.Println(l.alas * l.tinggi /2) 
}

func (l persegi) luasPersegi(){
	fmt.Println(l.sisi * l.sisi ) 
}

func (l persegiPanjang) luasPersegiPanjang(){
	fmt.Println(l.panjang * l.lebar ) 
}

// Soal 3
type phone struct{
	name, brand string
	year int
	colors []string
}

func (p *phone) addColor(colors string) {
	p.colors = append(p.colors, colors)
	fmt.Println(p.colors)
}

// Soal 4
type movie struct{
	title, genre string
	duration, year int
}

func  tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie){
	*dataFilm = append(*dataFilm, movie{title, genre, duration, year})
}

func main() {
	// Soal 1
	buah1 := buah{"Nanas", "Kuning", false, 9000}
	buah2 := buah{"Jeruk", "Oranye", true, 8000}
	buah3 := buah{"Semangka", "Hijau & Merah", true, 10000}
	buah4 := buah{"Pisang", "Kuning", false, 5000}

	fmt.Println("nama", "warna", "ada bijinya", "harga")
	fmt.Printf("%s	%s	%t	%d \n", buah1.nama, buah1.warna, buah1.adaBijinya, buah1.harga)
	fmt.Printf("%s	%s	%t	%d \n", buah2.nama, buah2.warna, buah2.adaBijinya, buah2.harga)
	fmt.Printf("%s  %s	%t	%d \n", buah3.nama, buah3.warna, buah3.adaBijinya, buah3.harga)
	fmt.Printf("%s	%s	%t	%d \n", buah4.nama, buah4.warna, buah4.adaBijinya, buah4.harga)

	// Soal 2
	segitiga := segitiga{10, 5}
	persegi := persegi{5}
	persegiPanjang := persegiPanjang{10, 5}
	segitiga.luasSegitiga()
	persegi.luasPersegi()
	persegiPanjang.luasPersegiPanjang()

	// Soal 3
	phone := phone{"s23", "samsung", 2003, []string{"merah", "hitam"}}
	phone.addColor("putih")

	// Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, data := range dataFilm{
		title := data.title
		duration := data.duration
		genre := data.genre
		year := data.year

		fmt.Printf("%d. title : %s\n   duration : %d jam\n   genre : %s\n   year : %d\n\n", i+1, title, duration/60, genre, year)
	} 
}