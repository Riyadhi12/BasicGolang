package main 

import "fmt"

func main(){
	 var months = [12] string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	 }
	 /*Tipe data slice memiliki tiga data ,Pointer,lenght,dan capacity
	 Pointer : penunjuk data pertama di array pada slice
	 lenght : panjang slice
	 capacity : Kapasitas dari slice
	 */
	 var slice1 = months[4:7]
	 fmt.Println(slice1)
	 fmt.Println(len(slice1)) // melihat panjang
	 fmt.Println(cap(slice1)) // melihat kapasitas dari slice

	 var slice2 = months[10:]
	 fmt.Println(slice2)

	 var slice3 = append(slice2, "Riyadhi") //menambah data -> jika array sudah full maka akan dibuat array baru
	 fmt.Println(slice3)

	 newSlice := make([]string ,2 ,5)
	 newSlice[0] = "Muhammad"
	 newSlice[1] = "Riyadhi"
	 fmt.Println(newSlice)
	 fmt.Println(len(newSlice))
	 fmt.Println(cap(newSlice))

	 copySlice := make([]string, len(newSlice), cap(newSlice)) // copy slice harus memiliki panjang dan kapasitas yang sama dengan yang di copy
	 copy(copySlice, newSlice)
	 fmt.Println(copySlice)

	 /*perbedaan array dan slice*/

	 iniArray := [5]int{1,2,3,4,5};
	 iniSlice := []int{1,2,3,4,5};

}