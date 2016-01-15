package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var count int
	var tmp string
	st := make([]string, 0)

	fmt.Scanf("%d", &count)
	if count < 1 || count > 20 {
		fmt.Printf("Input value must be between 1 and 20, you entered: %v \n", count)
		fmt.Errorf("Input value must be between 1 and 20, you entered: %v \n", count)
		os.Exit(1)

	}
	for i := 0; i < count; i++ {
		fmt.Scanf("%s", &tmp)
		ln := len(tmp)
		if ln < 1 || ln > 100 {
			fmt.Printf("Input sting length must be between 1 and 100, you entered the string of length: %d \n", ln)
			fmt.Errorf("Input sting length must be between 1 and 100, you entered the string of length: %d \n", ln)
			os.Exit(1)
		}

		st = append(st, strings.ToUpper(tmp))
		tmp = ""

	}

	for _, v := range st {
		if v == reverse(v) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
