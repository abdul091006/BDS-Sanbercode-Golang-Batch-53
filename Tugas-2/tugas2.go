package main

import (
  "fmt"
  "strconv"
  "strings"
)

func main() {
  // Tugas 1
  variable1 := "Bootcamp"
  variable2 := "Digital"
  variable3 := "Skill"
  variable4 := "Sanbercode"
  variable5 := "Golang"
  fmt.Printf("%s %s %s %s %s \n", variable1, variable2, variable3, variable4, variable5)
  
  // Tugas 2
  halo := "Halo Dunia"
  find := "Dunia"
  replaceWith := "Golang"
  haloGolang := strings.Replace(halo, find, replaceWith, 1)
  fmt.Println(haloGolang)
  
  // Tugas 3
  var kataPertama = "saya"
  var kataKedua = "senang"
  var kataKetiga = "belajar"
  var kataKeempat = "golang"

  find2 := string(kataKedua[0])
  find3 := string(kataKetiga[len(kataKetiga)-1])

  sUpper := strings.ToUpper(string(kataKedua[0]))
  rUpper := strings.ToUpper(string(kataKetiga[len(kataKetiga)-1]))

  kataKeduaUpper := strings.Replace(kataKedua, find2, sUpper, 1)
  kataKetigaUpper := strings.Replace(kataKetiga, find3, rUpper, 1)
  kataKeempatUpper := strings.ToUpper(kataKeempat)
  fmt.Printf("%s %s %s %s \n", kataPertama, kataKeduaUpper, kataKetigaUpper, kataKeempatUpper)
  
  // Tugas 4
  var angkaPertama = "8"
  var angkaKedua = "5"
  var angkaKetiga = "6"
  var angkaKeempat = "7"

  angkaPertamaNum, _ := strconv.Atoi(angkaPertama)
  angkaKeduaNum, _ := strconv.Atoi(angkaKedua)
  angkaKetigaNum, _ := strconv.Atoi(angkaKetiga)
  angkaKeempatNum, _ := strconv.Atoi(angkaKeempat)

  total := angkaPertamaNum + angkaKeduaNum + angkaKetigaNum + angkaKeempatNum
  fmt.Printf("Total: %d \n", total)

  // Tugas 5
  kalimat := "halo halo bandung"
  angka := 2021
  find = "halo"
  replaceWith = "Hi"
  hiBandung :=  strings.Replace(kalimat, find, replaceWith, -1)
  fmt.Printf("\"%s\" - %d \n", hiBandung, angka)
}