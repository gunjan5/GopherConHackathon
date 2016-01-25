// Go program to read an integer from STDIN and output it to STDOUT
package main

import (
	"fmt"
)

func main() {
    // Declare the variable
    var n int
    count:=make([]int, 20)
    
    // Read the variable from STDIN
    fmt.Scanf("%v", &n)

    for i := 0; i < n; i++ {
    	count[i]=(i+1)*10
    	//print count[i]
    	
    }
    
    // Output the variable to STDOUT
    fmt.Println(n)
    fmt.Printf("%T\n", n)
    fmt.Println(count)


    
}
