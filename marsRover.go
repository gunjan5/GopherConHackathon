package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {

	sos := "SOS"
	var in string
	var count int

	fmt.Scanf("%s", &in)
	ln := len(in)

	if ln < 1 || ln > 99 {
		fmt.Printf("input length must be between 1 and 99: %d \n", ln)
		fmt.Errorf("input length must be between 1 and 99: %d \n", ln)
		os.Exit(1)
	}
	if ln%3 != 0 {
		fmt.Printf("Must be a modulo of 3: %d \n", ln)
		fmt.Errorf("Must be a modulo of 3: %d \n", ln)
		os.Exit(1)
	}

	for i := 0; i < ln; i = i + 3 {

		if sos != in[i:i+3] {
			for k, v := range in[i : i+3] {
				if unicode.IsUpper(v) {
					if sos[k] != byte(v) {
						count++
					}
				} else {
					fmt.Printf("Must be an UPPERCASE and English letter: %s \n", string(v))
					fmt.Errorf("Must be an UPPERCASE and English letter: %s \n", string(v))
					os.Exit(1)
				}

			}
		}

	}
	fmt.Println(count)
}
