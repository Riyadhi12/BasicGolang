package main

import (
	"fmt"
	"math"
)

func main(){
	counter := 1

	for counter <= 10{
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	//for dengan statement
	for counter :=1; counter <= 10; counter++{
		fmt.Println("Perulangan ke", counter)
	}
	//megakses slice
	slice := []string{"Muhammad","riyadhi","Lubis"}
	for i := 0; i< len(slice); i++{
		fmt.Println(slice[i])
	}
	//menggunakan for range
	//Jika variabel tidak dibutuhkan gunakan _
	for i,value := range slice{
		fmt.Println("Index",i,"=",value)
	}

	var i float64
	for i = 0; i <= 5; i++{
		if (math.Mod (i , 2) == 0  ){
			fmt.Println("**");
			continue;
		}
		fmt.Println("angka ",i," Bukan 3")
	}
}