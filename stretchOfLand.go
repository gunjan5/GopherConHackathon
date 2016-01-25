package main

import (
	"fmt"
	"os"
)

func main() {
	var count, tmp, max int
	var val []int

	fmt.Scanf("%d", &count)
	if count < 1 || count > 100000 {
		fmt.Printf("Input must be between 1 and 100000, you entered: %d \n", count)
		fmt.Errorf("Input must be between 1 and 100000, you entered: %d \n", count)
		os.Exit(1)

	}

	sum := make([]int, count)

	total := make([]int, count, count*count)

	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &tmp)
		if tmp < (-100000) || tmp > 100000 {
			fmt.Printf("Input must be between -100000 and 100000, you entered: %d \n", tmp)
			fmt.Errorf("Input must be between -100000 and 100000, you entered: %d \n", tmp)
			os.Exit(1)

		}
		val = append(val, tmp)
	}

	for i, v := range val {
		if i == 0 {
			sum[i] = v
			max = v
		} else {
			sum[i] = sum[i-1] + v
			if max < sum[i] {
				max = sum[i]
			}
		}
	}

	copy(total, sum)

	for l := 0; l < len(sum)-1; l++ {
		for i := l + 1; i < len(sum); i++ {
			diff := sum[i] - sum[l]
			total = append(total, diff)
			if max < diff {
				max = diff
			}
		}

	}

	fmt.Println(max)

}
