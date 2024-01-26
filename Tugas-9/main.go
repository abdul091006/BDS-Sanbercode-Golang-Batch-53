package main

import (
	"BDS-Sanbercode-Golang-Batch-53/Tugas-9/controllers"
	"fmt"
)

func main() {
	// Soal 1
	var segitigaSamaSisi controllers.HitungBangunDatar = controllers.SegitigaSamaSisi{
		Alas:   5,
		Tinggi: 5,
	}
	fmt.Println("Luas      :", segitigaSamaSisi.Luas())
	fmt.Println("Keliling  :", segitigaSamaSisi.Keliling())

	var persegiPanjang controllers.HitungBangunDatar = controllers.PersegiPanjang{
		Panjang: 10,
		Lebar:   5,
	}
	fmt.Println("Luas      :", persegiPanjang.Luas())
	fmt.Println("Keliling  :", persegiPanjang.Keliling())

	var tabung controllers.HitungBangunRuang = controllers.Tabung{
		JariJari: 7,
		Tinggi:   5,
	}
	fmt.Println("Volume      :", tabung.Volume())
	fmt.Println("Luas Permukaan  :", tabung.LuasPermukaan())

	var balok controllers.HitungBangunRuang = controllers.Balok{
		Panjang: 10,
		Lebar:   5,
		Tinggi:  5,
	}
	fmt.Println("Volume      :", balok.Volume())
	fmt.Println("Luas Permukaan  :", balok.LuasPermukaan())

	// Soal 2
	var counterHp controllers.MerekHp

	counterHp = controllers.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze ,Mystic White, Mystic Black"},
	}

	fmt.Println(counterHp.PanggilHp())

	// Soal 3
	fmt.Println(controllers.LuasPersegi(4, true))
	fmt.Println(controllers.LuasPersegi(8, false))
	fmt.Println(controllers.LuasPersegi(0, true))
	fmt.Println(controllers.LuasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	fmt.Println(controllers.Gabungkan(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua))
}
