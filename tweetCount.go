//Flow is explained here: https://g.twimg.com/dev/documentation/image/appauth_0.png

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
	"os"
	"strconv"
)

const (
	ConsumerKey    = "O4JaKdDfTthJCXPPOyGg926K4"
	ConsumerSecret = "8ou0302gaXHchZhzqn6a1lMtPrRUiysoFlyfL6eOpz0QivYnrz"
)

func main() {

	client := &http.Client{}

	encodedKeySecret := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		url.QueryEscape(ConsumerKey),
		url.QueryEscape(ConsumerSecret))))

	reqBody := bytes.NewBuffer([]byte(`grant_type=client_credentials`))

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", reqBody)
	if err != nil {
		log.Fatal(err, client, req)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedKeySecret))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(reqBody.Len()))

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

	var id int

	fmt.Scanf("%d", &id)
	if id < 1 { //actually tweet numbers start from 20 apparently
		fmt.Printf("Input must be > 1, you entered: %d \n", id)
		fmt.Errorf("Input must be > 1, you entered: %d \n", id)
		os.Exit(1)
	}

	//id:=210462857140252672

	twitterEndPoint := "https://api.twitter.com/1.1/statuses/show.json?id=" + strconv.Itoa(id)
	req, err = http.NewRequest("GET", twitterEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization",
		fmt.Sprintf("Bearer %s", b.AccessToken))

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}
	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

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
}
