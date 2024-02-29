package main

import "fmt"

func main(){
	 nilai := "Riyadhi"

	if nilai == "Riyadhi" {
		fmt.Println("Gacor kang")
	}else if nilai == "Adit"{
		fmt.Println("Tidak bisa")
	}else{
		fmt.Println("gaktau")
	}

	// Short Statemant // bisa juga digunakan dalam switch
	if lenght  := len(nilai) ; lenght > 2{
		fmt.Println("kepanjangan")
	}

	//Switch tanpa nilai(velue)
	lenght := len(nilai)
	switch  {
	case lenght > 10:
		fmt.Println("nama sudah benar")
	case lenght < 10 :
		fmt.Println("nama terlalu panjang")
	}
	
	
}