package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"os"
	"time"
)

const (
	ConsumerKey    = "O4JaKdDfTthJCXPPOyGg926K4"
	ConsumerSecret = "8ou0302gaXHchZhzqn6a1lMtPrRUiysoFlyfL6eOpz0QivYnrz"
)

func main() {

	start := time.Now()


	client := &http.Client{}
	//Step 1: Encode consumer key and secret
	encodedKeySecret := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		url.QueryEscape(ConsumerKey),
		url.QueryEscape(ConsumerSecret))))

	//Step 2: Obtain a bearer token
	//The body of the request must be grant_type=client_credentials
	reqBody := bytes.NewBuffer([]byte(`grant_type=client_credentials`))
	//The request must be a HTTP POST request
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", reqBody)
	if err != nil {
		log.Fatal(err, client, req)
	}
	//The request must include an Authorization header formatted as
	//Basic <base64 encoded value from step 1>.
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedKeySecret))
	//The request must include a Content-Type header with
	//the value of application/x-www-form-urlencoded;charset=UTF-8.
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(reqBody.Len()))

	//Issue the request and get the bearer token from the JSON you get back
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, respBody)
	}

	type BearerToken struct {
		AccessToken string `json:"access_token"`
	}
	var b BearerToken
	json.Unmarshal(respBody, &b)

	//choose your API endpoint that supports application only auth context
	//and create a request object with that


	var id int

	fmt.Scanf("%d", &id)
	if id < 1 {
		fmt.Printf("Input must be > 1, you entered: %d \n", id)
		fmt.Errorf("Input must be > 1, you entered: %d \n", id)
		os.Exit(1)
	}

	//id:=210462857140252672

	twitterEndPoint := "https://api.twitter.com/1.1/statuses/show.json?id="+strconv.Itoa(id)  
	req, err = http.NewRequest("GET", twitterEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Step 3: Authenticate API requests with the bearer token
	//include an Authorization header formatted as
	//Bearer <bearer token value from step 2>
	req.Header.Add("Authorization",
		fmt.Sprintf("Bearer %s", b.AccessToken))

	//Issue the request and get the JSON API response
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}
	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//fmt.Println(string(respBody))


	


	// res, _ := http.Get("https://api.twitter.com/1.1/statuses/show.json?id=210462857140252672") 
	// defer res.Body.Close()

	// body, _ := ioutil.ReadAll(res.Body)

	 var obj map[string]interface{}

	 err1 := json.Unmarshal(respBody, &obj)
	 if err1 != nil {
	 	panic(err)
	 }
	 res := obj["retweet_count"]
	rtCount, ok := res.(float64) //need to assert the interface to float64 first (can't assert it to int directly!!)
	if !ok {
		fmt.Println("Tweet not found: error converting response to float64")
		fmt.Errorf("Tweet not found: error converting response to float64")
		os.Exit(1)
	}
	 fmt.Println(int(rtCount))
	 fmt.Printf("****** Executed in time %s \n ", time.Since(start))
}








// package main  


// import (
// 	"fmt"
// 	"net/http"
// 	"log"
// 	"io/ioutil"
// 	"strings"
// )



// func main() {
// 	body := strings.NewReader(`id=210462857140252672`)
// req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/statuses/show.json", body)
// if err != nil {
// 	// handle err
// }
// req.Header.Set("Authorization", "OAuth oauth_consumer_key=\"O4JaKdDfTthJCXPPOyGg926K4\", oauth_nonce=\"affa19ff1b7bded14c436e1efc3fe24f\", oauth_signature=\"XXbBXHQjUNe%2B34SMF%2Bw6w5v%2BR1Y%3D\", oauth_signature_method=\"HMAC-SHA1\", oauth_timestamp=\"1454990438\", oauth_version=\"1.0\"")

// resp, err := http.DefaultClient.Do(req)
// if err != nil {
// 	// handle err
// }
// defer resp.Body.Close()

// respBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("response Status:", resp.Status)
// 	//fmt.Println("response Headers:", resp.Header)
// 	fmt.Println(string(respBody))

// }