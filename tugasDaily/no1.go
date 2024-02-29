package main

import "fmt"

func main(){
  var phi float32 = 3.14 ;
  var r,t ,hasil float32;
  
 fmt.Println("Menghitung Volume Tabung ")
 fmt.Println("Masukkan nilai r :")
 fmt.Scanln(&r)
 fmt.Println("Masukkan nilai t :")
 fmt.Scanln(&t)

 hasil = phi * r * r * t
 fmt.Println("Volume Tabung : ", hasil)

}
