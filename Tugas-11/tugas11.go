package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"time"
)

// Soal 1
func listPhone(phones []string, myWg *sync.WaitGroup) {
	for i, item := range phones {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(strconv.Itoa(i+1) + ". " + item)
	}
	myWg.Done()
}

// Soal 2
func getMovies(moviesChannel chan string, movies ...string) {
	for i, movie := range movies {
		moviesChannel <- strconv.Itoa(i+1) + ". " + movie
	}
	close(moviesChannel)
}

// Soal 3
func luasLingkaran(r float64, luas chan<- float64) {
	luas <- math.Pi * r * r

}

func kelilingLingkaran(r float64, keliling chan<- float64) {
	keliling <- math.Pi * 2 * r
}

func volumeTabung(r, t float64, volume chan<- float64) {
	volume <- math.Pi * r * r * t
}

// Soal 4
func countLuasPersegiPanjang(p int, l int, luasPersegiPanjang chan<- int) {
	luasPersegiPanjang <- p * l
}

func countKelilingPersegiPanjang(p int, l int, kelilingPersegiPanjang chan<- int) {
	kelilingPersegiPanjang <- 2 * (p + l)
}

func countVolumeBalok(p int, l int, t int, volumeBalok chan<- int) {
	volumeBalok <- p * l * t
}

func main() {
	// Soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var myWg sync.WaitGroup
	myWg.Add(1)
	go listPhone(phones, &myWg)
	myWg.Wait()

	// Soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	fmt.Println("List Movies:")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	// Soal 3
	luas := make(chan float64)
	keliling := make(chan float64)
	volume := make(chan float64)

	go luasLingkaran(8, luas)
	go kelilingLingkaran(14, keliling)
	go volumeTabung(20, 10, volume)

	fmt.Println("Luas:", <-luas)
	fmt.Println("Keliling:", <-keliling)
	fmt.Println("Volume:", <-volume)

	// Soal 4
	luasPersegiPanjang := make(chan int)
	kelilingPersegiPanjang := make(chan int)
	volumeBalok := make(chan int)
	go countLuasPersegiPanjang(6, 4, luasPersegiPanjang)
	go countKelilingPersegiPanjang(6, 4, kelilingPersegiPanjang)
	go countVolumeBalok(6, 4, 10, volumeBalok)

	for j := 0; j < 3; j++ {
		select {
		case luasPP := <-luasPersegiPanjang:
			fmt.Println("Luas: " + strconv.Itoa(luasPP))
		case kelPP := <-kelilingPersegiPanjang:
			fmt.Println("Keliling: " + strconv.Itoa(kelPP))
		case volPP := <-volumeBalok:
			fmt.Println("Volume: " + strconv.Itoa(volPP))
		}
	}
}
