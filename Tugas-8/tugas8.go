package main

import (
	"fmt"
	"math"
	"strconv"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * s.alas * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Phi * (t.jariJari * t.jariJari) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Phi * t.jariJari * (t.jariJari * t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

// Soal 2
type merekHp interface {
	panggilHp() string
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (po phone) panggilHp() string {

	for _, s := range po.colors {
		return "name : " + po.name + "\n" + "brand : " + po.brand + "\n" + "Year : " + strconv.Itoa(po.year) + "\n" + "Colors : " + s
	}
	return ""
}

// Soal 3
func luasPersegi(sisi uint, showText bool) interface{} {
	switch {
	case sisi > 0 && showText:
		return "luas persegi dengan sisi " + strconv.Itoa(int(sisi)) + " cm adalah " + strconv.Itoa(int(sisi*sisi)) + " cm"
	case sisi > 0 && !showText:
		return sisi * sisi
	case sisi == 0 && showText:
		return "Maaf anda belum menginput sisi dari persegi"
	default:
		return nil
	}
}

// Soal 4
func gabungkan(prefix, angka1, angka2 interface{}) string {
    total := 0
    if p, ok := prefix.(string); ok {
        result := p
        if a1, ok := angka1.([]int); ok {
            for _, num := range a1 {
                total += num
                result += fmt.Sprintf("%d + ", num)
            }
        }
        if a2, ok := angka2.([]int); ok {
            for _, num := range a2 {
                total += num
                result += fmt.Sprintf("%d + ", num)
            }
        }
        result = result[:len(result)-2] 
        result += fmt.Sprintf("= %d", total)
        return result
    }
    return ""
}

func main() {

	// Soal 1
	var segitigaSamaSisi hitungBangunDatar = segitigaSamaSisi{
		alas:   5,
		tinggi: 5,
	}
	fmt.Println("luas      :", segitigaSamaSisi.luas())
	fmt.Println("keliling  :", segitigaSamaSisi.keliling())

	var persegiPanjang hitungBangunDatar = persegiPanjang{
		panjang: 10,
		lebar:   5,
	}
	fmt.Println("luas      :", persegiPanjang.luas())
	fmt.Println("keliling  :", persegiPanjang.keliling())

	var tabung hitungBangunRuang = tabung{
		jariJari: 7,
		tinggi:   5,
	}
	fmt.Println("luas      :", tabung.volume())
	fmt.Println("keliling  :", tabung.luasPermukaan())
	
	var balok hitungBangunRuang = balok{
		panjang: 10,
		lebar: 5,
		tinggi: 5,
	}
	fmt.Println("luas      :", balok.volume())
	fmt.Println("keliling  :", balok.luasPermukaan())

	// Soal 2
	var counterHp merekHp

	counterHp = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze ,Mystic White, Mystic Black"},
	}

	fmt.Println(counterHp.panggilHp())

	// Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	fmt.Println(gabungkan(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua))
}
