package main

import "fmt"

func main(){
	var nilai float64;
	fmt.Print("Masukkan nilai : ")
	fmt.Scanln(&nilai)

	if nilai >= 85 && nilai <= 100 {
		fmt.Println("Grade : A")
	}else if nilai >= 70 && nilai <= 84{
		fmt.Println("Grade : B")
	}else if nilai >= 55 && nilai <= 69{
		fmt.Println("Grade : C")
	}else if nilai >= 40 && nilai <= 50 {
		fmt.Println("Grade : D")
	}else if nilai >= 0 && nilai <= 39 {
		fmt.Println("Grade : E")
	}else{
		fmt.Println("Nilai Invalid")
	}
	

}