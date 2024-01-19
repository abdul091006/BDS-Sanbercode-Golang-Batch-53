package main

import "fmt"

func main() {
	// Soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 1 && i%3 == 0  {
			fmt.Printf("%d - I Love Coding \n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d - Santai \n", i)
		} else {
			fmt.Printf("%d - Berkualitas \n", i)
		}
	}

	// Soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		println("")
	}

	// Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	kalimatSlice := kalimat[2:]
	fmt.Println(kalimatSlice) 

	// Soal 4
	var sayuran = []string{}
	sayur := []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}

	for i := range sayur{
		sayuran = append(sayuran, sayur[i])
	}

	for i, value := range sayuran{
		fmt.Printf("%d. %s \n", i+1, value)
	}

	// Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	total := 1
	for i, value := range satuan{
		total *= value
		fmt.Printf("%s = %d \n", i, value)
	}
	fmt.Println("Volume Balok =", total)
}