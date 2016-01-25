package main


import (
	"fmt"
	_"strings"
	_"strconv"
	"sort"
)


func main() {
	
	var count int
	var t1, t2 int
	//st := make([]string, 0)
	in := make(map[int][]int)

	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		fmt.Scanf("%d %d", &t1, &t2)
		//fmt.Println(t1, t2)
		//out:= strings.Fields(tmp)
		//fmt.Println(out[0])

		if t2 < 1 || t2 > 5 {
			fmt.Printf("Input sting length must be between 1 and 100, you entered the string of length: %d \n", t2)
			fmt.Errorf("Input sting length must be between 1 and 100, you entered the string of length: %d \n", t2)
			os.Exit(1) 
		}

		in[t2]=append(in[t2], t1)
		//ln := len(tmp)


		//st = append(st, strings.ToUpper(tmp))
		t1,t2 = 0,0

	}

	for i:=1; i<=5; i++ {
		sort.Sort(sort.IntSlice(in[i]))
		for k:=0; k<len(in[i]); k++{
			fmt.Println(in[k])
		}
	}
//fmt.Println(in)


}