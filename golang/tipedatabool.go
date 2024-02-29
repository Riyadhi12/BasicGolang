package main

import "fmt"

func main(){
	var nilai = 80;
	var absensi = 75;

	var lulusUjian = nilai >= 80;
	var lulusAbsensi = absensi >= 75;

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

}