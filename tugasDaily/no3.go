package main

import (
	"fmt"
)

func main(){
 

  for i := 1 ; i <= 100 ; i ++{
    if i% 4 == 0 {
    output := i * i;
      fmt.Println(output)
   }else if i%7 == 0 {
    pangkat3 := i * i * i;
   fmt.Println(pangkat3)
   }else if i%7 == 0 && i% 4 == 0{
   fmt.Println("Buzz")
   }else{
   fmt.Println(i)
   }}

  
}