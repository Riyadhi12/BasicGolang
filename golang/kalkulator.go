package main

import (
	"fmt"
	"math"
)

func main(){
	var pilih string
	var num1, num2, result float64 
	for {
	fmt.Println("Pilih Menu")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Perpangkatan")
	fmt.Print("Pilih : ")
	fmt.Scanln(&pilih)
	
	switch pilih{
	case "1" :
		fmt.Println("Penjumlahan")
		fmt.Print("Masukkan nilai pertama: ")
		fmt.Scanln(&num1)
		fmt.Print("Masukkan nilai kedua: ")
		fmt.Scanln(&num2)
		
		result = num1 + num2
		fmt.Println("Hasil:", result)
	case "2" :
		fmt.Println("Pengurangan")
		fmt.Print("Masukkan nilai pertama: ")
		fmt.Scanln(&num1)
		fmt.Print("Masukkan nilai kedua: ")
		fmt.Scanln(&num2)
		
		result = num1 - num2
		fmt.Println("Hasil:", result)
	case "3" :
		fmt.Println("Perkalian")
		fmt.Print("Masukkan nilai pertama: ")
		fmt.Scanln(&num1)
		fmt.Print("Masukkan nilai kedua: ")
		fmt.Scanln(&num2)
		
		result = num1 * num2
		fmt.Println("Hasil:", result)
	case "4" :
		fmt.Println("Pembagian")
		fmt.Print("Masukkan nilai pertama: ")
		fmt.Scanln(&num1)
		fmt.Print("Masukkan nilai kedua: ")
		fmt.Scanln(&num2)
		
		result = num1 / num2
		fmt.Println("Hasil:", result)
	case "5" :
		fmt.Println("Perpangkatan")
		fmt.Print("Masukkan nilai basis: ")
        fmt.Scanln(&num1)

        fmt.Print("Masukkan pangkat: ")
        fmt.Scanln(&num2)
        
        result = math.Pow(num1, num2) 
        fmt.Println("Hasil:", result)

	default:
        fmt.Println("Pilihan salah")
	}
	fmt.Print("Lanjut menghitung? (y/n) : ")
    var answer string
    fmt.Scanln(&answer)

    if answer == "n" {
      break
    } 
	}
}