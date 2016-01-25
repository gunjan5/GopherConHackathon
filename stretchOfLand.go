package main

import (
	"fmt"
)



func main() {
	var count int 
	var tmp int
	var val []int

	fmt.Scanf( "%d", &count)
	for i := 0; i < count; i++ {
	fmt.Scanf("%d", &tmp)
	val=append(val,tmp)	
	}


	fmt.Println(val)
}




