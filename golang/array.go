package main
/* array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama , index dimulai dari 0*/
import "fmt"

func main(){
	// array secara manual
	var name [3] string;
	name[0] = "Muhammad";
	name[1] = "Riyadhi";
	name[2] = "Lubis";

	fmt.Println(name[2]);

	// dengan memasukkan nilainya langsung
	var values = [3] int {60,70,80,}
	
	fmt.Println(values)

	fmt.Println(len(values)) // mengakses panjang array
	
}



