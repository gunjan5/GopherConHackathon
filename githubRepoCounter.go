package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	var id int

	fmt.Scanf("%d", &id)
	if id < 1 {
		fmt.Printf("Input must be > 1, you entered: %d \n", id)
		fmt.Errorf("Input must be > 1, you entered: %d \n", id)
		os.Exit(1)
	}
	start := time.Now()
	
	resp, _ := http.Get("https://api.github.com/user/" + strconv.Itoa(id)) //5966968
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var obj map[string]interface{}

	err := json.Unmarshal(body, &obj)
	if err != nil {
		panic(err)
	}
	repos := obj["public_repos"]
	i, ok := repos.(float64) //need to assert the interface to float64 first (can't assert it to int directly!!)
	if !ok {
		fmt.Println("User ID not found: error converting response to float64")
		fmt.Errorf("User ID not found: error converting response to float64")
		os.Exit(1)
	}
	fmt.Println(int(i))
	fmt.Printf("****** Executed in time %s \n ", time.Since(start))

}
