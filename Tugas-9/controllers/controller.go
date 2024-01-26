package controllers

import (
	"fmt"
	"math"
	"strconv"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * s.Alas * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Phi * (t.JariJari * t.JariJari) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Phi * t.JariJari * (t.JariJari * t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi))
}

// Soal 2
type MerekHp interface {
	PanggilHp() string
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (po Phone) PanggilHp() string {

	for _, s := range po.Colors {
		return "name : " + po.Name + "\n" + "brand : " + po.Brand + "\n" + "Year : " + strconv.Itoa(po.Year) + "\n" + "Colors : " + s
	}
	return ""
}

// Soal 3
func LuasPersegi(sisi uint, showText bool) interface{} {
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
func Gabungkan(prefix, angka1, angka2 interface{}) string {
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
