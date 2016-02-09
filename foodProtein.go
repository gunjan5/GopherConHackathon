package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)



func main() {
	//http://api.nal.usda.gov/ndb/reports/?ndbno=09176&type=f&format=json&api_key=DEMO_KEY

	resp, _ := http.Get("http://api.nal.usda.gov/ndb/reports/?ndbno=09176&type=f&format=json&api_key=DEMO_KEY") //https://api.github.com/user/" + strconv.Itoa(id)) //5966968
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var obj1 map[string]interface{}
		var obj2 map[string]interface{}
			var obj3 map[string]interface{}
			//	var obj4 map[string]interface{}

	err := json.Unmarshal(body, &obj1)
	if err != nil {
		panic(err)
	}
	r1 := obj1["report"]
	// i, ok := repos.(float64) //need to assert the interface to float64 first (can't assert it to int directly!!)
	// if !ok {
	// 	fmt.Println("User ID not found: error converting response to float64")
	// 	fmt.Errorf("User ID not found: error converting response to float64")
	// 	os.Exit(1)
	// }
	//fmt.Println(int(i))
	//fmt.Println( r1)



r11, ok := r1.([]byte)
if !ok {
		fmt.Println("User ID not found: error converting response to float64")
		fmt.Errorf("User ID not found: error converting response to float64")
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(r11), &obj2)
	if err != nil {
		panic(err)
	}
	r2 := obj2["food"]



r21, ok := r2.([]byte)

if !ok {
		fmt.Println("User ID not found: error converting response to float64")
		fmt.Errorf("User ID not found: error converting response to float64")
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(r21), &obj3)
	if err != nil {
		panic(err)
	}
	r3 := obj3["nutrients"]


	// err = json.Unmarshal(r3, &obj4)
	// if err != nil {
	// 	panic(err)
	// }
	// r3 := obj3["nutrients"]

	fmt.Println(r3)


}