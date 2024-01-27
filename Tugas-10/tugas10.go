package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Soal 1
func soal1(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

// Soal 2
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("\nERROR : ", message)
	}
}
func kelilingSegitigaSamaSisi(sisi uint, showText bool) (string, error) {
	switch {
	case sisi > 0 && showText:
		return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(int(sisi)) + " cm adalah " + strconv.Itoa(int(sisi*3)) + " cm", nil
	case sisi > 0 && !showText:
		return strconv.Itoa(int(sisi*3)), nil
	case sisi == 0 && showText:
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	default:
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
}

// Soal 3
func tambahAngka(number int, angka *int) {
	*angka += number
}

func cetakAngka(angka *int) {
	fmt.Println(strconv.Itoa(*angka))
}

// Soal 4
func tambahPhone(object string, phones *[]string) {
	time.Sleep(time.Second * 1)
	*phones = append(*phones, object)
	fmt.Println(strconv.Itoa(len(*phones)) + ". " + object)
}

// Soal 5
func luasLingkaran(jariJari float64) {
	result := math.Phi * jariJari * jariJari
	fmt.Println(math.Round(result))
}

func kelilingLingkaran(jariJari float64) {
	result := 2 * math.Phi * jariJari
	fmt.Println(math.Round(result))
}

// Soal 6
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func main() {
	defer endApp()
	angka := 1
	// Soal 1
	defer soal1("Golang Backend Development", 2021)

	// Soal 2
	segitiga1, err := kelilingSegitigaSamaSisi(4, true)
	if err == nil {
		fmt.Println(segitiga1)
	} else {
		panic(err.Error())
	}

	segitiga2, err := kelilingSegitigaSamaSisi(8, false)
	if err == nil {
		fmt.Println(segitiga2)
	} else {
		panic(err.Error())
	}

	segitiga3, err := kelilingSegitigaSamaSisi(0, true)
	if err == nil {
		fmt.Println(segitiga3)
	} else {
		panic(err.Error())
	}

	segitiga4, err := kelilingSegitigaSamaSisi(0, false)
	if err == nil {
		fmt.Println(segitiga4)
	} else {
		panic(err.Error())
	}

	// Soal 3
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 4
	var phones = []string{}
	tambahPhone("Xiaomi", &phones)
	tambahPhone("Asus", &phones)
	tambahPhone("Iphone", &phones)
	tambahPhone("Samsung", &phones)
	tambahPhone("Oppo", &phones)
	tambahPhone("Realme", &phones)
	tambahPhone("Vivo", &phones)

	// Soal 5
	luasLingkaran(7)
	luasLingkaran(10)
	luasLingkaran(15)
	kelilingLingkaran(7)
	kelilingLingkaran(10)
	kelilingLingkaran(15)

	// Soal 6
	var panjang = flag.Int("panjang", 5, "type your panjang")
	var lebar = flag.Int("lebar", 4, "type your lebar")

	flag.Parse()
	fmt.Println(luasPersegiPanjang(*panjang, *lebar))
	fmt.Println(kelilingPersegiPanjang(*panjang, *lebar))
}
