package main

import (
	"fmt"
	"os"
	"sort"
	//"time"
)

func main() {
	//start := time.Now()

	var tmp, max, fmax int
	var val, in []int

	fmt.Scanf("%d", &tmp)
	if tmp < 1 || tmp > 100000 {
		fmt.Printf("Input must be between 1 and 100000, you entered: %d \n", tmp)
		os.Exit(1)

	}
	val = append(val, tmp)
	tmp = 0
	fmt.Scanf("%d", &tmp)

	if tmp < 1 || tmp > val[0] {
		fmt.Printf("Input must be between 1 and 100000 & less than the first input, you entered: %d \n", tmp)
		os.Exit(1)

	}
	val = append(val, tmp)
	tmp = 0

	for i := 0; i < val[1]; i++ {
		fmt.Scanf("%d", &tmp)
		if tmp < 0 || tmp >= val[0] {
			fmt.Printf("Input must be between 1 and 100000 & less than total number of cities, you entered: %d \n", tmp)
			os.Exit(1)
		}
		in = append(in, tmp)
	}

	sort.IntSlice(in).Sort()

	for i := 0; i < len(in)-1; i++ {
		diff := in[i+1] - in[i]
		if diff > max {
			max = diff
		}
	}

	final := []int{in[0] - 0, val[0] - in[len(in)-1] - 1, max / 2}
	for _, v := range final {
		if v > fmax {
			fmax = v
		}
	}

	fmt.Println(fmax)

	//fmt.Printf("****** Executed in time %s \n ", time.Since(start))

}
