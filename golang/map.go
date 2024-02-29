package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "Riyadhi",
		"Address" : "Madina",
	}
	fmt.Println(person)
	person["Tittle"] = "Programmer" // menambah data
	
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Address"])

	
    var book map[string]string = make(map[string]string)
	book["tittle"] = "Belajar Go-Lang"
	book["author"] = "Riyadhi"
	book["Sorry"] = "Salah"
	fmt.Println(book)

	delete(book, "Sorry")
	fmt.Println(book)
}