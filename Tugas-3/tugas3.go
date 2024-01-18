package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjangNum, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangNum, _ := strconv.Atoi(lebarPersegiPanjang)
	alasSegitigaNum, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitigaNum, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang = panjangPersegiPanjangNum * lebarPersegiPanjangNum
	kelilingPersegiPanjang = 2 * (panjangPersegiPanjangNum + lebarPersegiPanjangNum)
	luasSegitiga = (alasSegitigaNum * tinggiSegitigaNum) / 2

	fmt.Printf("Luas Persegi Panjang: %d \n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang: %d \n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga: %d \n", luasSegitiga)


	// Soal 2
	nilaiJohn := 80
	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	nilaiDoe := 50
	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}


	// Soal 3
	var tanggal = 9;
	var bulan = 10; 
	var tahun = 2006;

	var bulanStr string
	switch bulan {
	case 1:
		bulanStr = "Januari"
	case 2:
		bulanStr = "Februari"
	case 3:
		bulanStr = "Maret"
	case 4:
		bulanStr = "April"
	case 5:
		bulanStr = "Mei"
	case 6:
		bulanStr = "Juni"
	case 7:
		bulanStr = "Juli"
	case 8:
		bulanStr = "Agustus"
	case 9:
		bulanStr = "September"
	case 10:
		bulanStr = "Oktober"
	case 11:
		bulanStr = "November"
	case 12:
		bulanStr = "Desember"
	}

	fmt.Printf("%d %s %d \n", tanggal, bulanStr, tahun)


	// Soal 4
	tahunKelahiran := 2006

	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		fmt.Println("Generasi X")
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		fmt.Println("Generasi Y (millenials)")
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		fmt.Println("Generasi Z")
	} else {
		fmt.Println("Tidak terdaftar")
	}
}
