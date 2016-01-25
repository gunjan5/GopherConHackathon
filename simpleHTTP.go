package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var st []string
	var tmp string

	for i := 0; i < 3; i++ {
		fmt.Scanf("%s", &tmp)

		st = append(st, tmp)
		tmp = ""

	}

	if !strings.HasPrefix(st[0], "http://") && !strings.HasPrefix(st[0], "https://") {
		st[0] = "http://" + st[0]
	}

	res, err := http.Get(st[0]) //"http://www.example.com")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(strings.Replace(string(body), st[1], st[2], -1)) //"Example", "Test", -1) )

}
