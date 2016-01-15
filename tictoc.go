package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var st []string
	var tmp string

	for i := 0; i < 3; i++ {
		fmt.Scanf("%s", &tmp)
		ln := len(tmp)
		if ln != 3 {
			fmt.Printf("Input sting length must be between 1 and 3, you entered the string of length: %d \n", ln)
			fmt.Errorf("Input sting length must be between 1 and 3, you entered the string of length: %d \n", ln)
			os.Exit(1)
		}

		if strings.Count(tmp, "X")+strings.Count(tmp, "O")+strings.Count(tmp, ".") != 3 {
			fmt.Printf("Invalid input: %v. It must be O, X or . \n", tmp)
			fmt.Errorf("Invalid input: %v. It must be O, X or . \n", tmp)
			os.Exit(1)

		}

		st = append(st, tmp)
		tmp = ""

	}

	x, o := check(st)

	switch {
	case o && x:
		fmt.Println("Its a tie")
	case x:
		fmt.Println("X wins")
	case o:
		fmt.Println("O wins")
	default:
		fmt.Println("Nobody won")
		os.Exit(1)
	}

	//fmt.Println(x, o)

}

func check(s []string) (xwin, owin bool) {
	ftmp, btmp := s[0][0], s[0][len(s)-1]
	var fwdgoodsofar, bckgoodsofar, x, o bool

	for k, _ := range s {

		if s[k][0] == s[k][1] && s[k][0] == s[k][2] {
			x, o = whoIsIt(string(s[k][0]))
			xwin = xwin || x
			owin = owin || o

		}

		if s[0][k] == s[1][k] && s[0][k] == s[2][k] {
			x, o = whoIsIt(string(s[0][k]))
			xwin = xwin || x
			owin = owin || o

		}

		if k > 0 {
			if s[k][k] == ftmp && s[k-1][k-1] == ftmp {
				fwdgoodsofar = true
			} else {
				fwdgoodsofar = false
			}

			if s[k][len(s)-k-1] == btmp && s[k-1][len(s)-k] == btmp {
				bckgoodsofar = true
			} else {
				bckgoodsofar = false
			}

		}

	}

	if fwdgoodsofar {
		//fmt.Println("fwd good so far")
		x, o = whoIsIt(string(ftmp))
		xwin = xwin || x
		owin = owin || o
	}
	if bckgoodsofar {
		//fmt.Println("bck good so far")
		x, o = whoIsIt(string(btmp))
		xwin = xwin || x
		owin = owin || o
	}

	return

}

func whoIsIt(s string) (xwin, owin bool) {
	if s == "X" {
		xwin = true
	} else if s == "O" {
		owin = true
	}

	return
}
