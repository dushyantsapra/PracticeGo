package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	sampleHTTPRequest()
	fmt.Println()

	redirectHttpResponse()
	fmt.Println()

	SamplePostRequest()
	fmt.Println()

	makingCustomizedRequests()
	fmt.Println()

	makingCustomizedRequest2()
	fmt.Println()

	makingCustomizedRequest3()
	fmt.Println()
}

func sampleHTTPRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal("error occurred ", err)
	}

	data, _ := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(res.Status, " ", res.StatusCode)
	fmt.Printf("%s\n", data)
}

func redirectHttpResponse() {
	res, err := http.Get("https://bit.ly/2IRnmVm")

	if err != nil {
		log.Fatal("Error occurred ", err)
		panic("Error while fetching the Response")
	}

	data, _ := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(res.Header.Get("Content-Type"))
	if res.Header.Get("Content-Type") == "image/jpeg" {
		os.WriteFile("keyboard.jpeg", data, 0777)
		fmt.Println("Image Saved")
	} else {
		log.Fatal("Response is not an image")
	}
}

func SamplePostRequest() {
	requestBody := strings.NewReader(`
		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}
	`)

	// post some data
	res, err := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"application/json; charset=UTF-8",
		requestBody,
	)

	// check for response error
	if err != nil {
		log.Fatal(err)
	}

	// read response data
	data, _ := io.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// print request `Content-Type` header
	requestContentType := res.Request.Header.Get("Content-Type")
	fmt.Println("Request content-type:", requestContentType)

	// print response body
	fmt.Printf("%s\n", data)
}

func makingCustomizedRequests() {
	requestyBody := io.NopCloser(strings.NewReader(`{
		"fName":"Dushyant"
		"lName":"Sapra"
	}`))

	url, _ := url.Parse("http://dummy.restapiexample.com/api/v1/create")
	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"Content-Type": {"application/json; charset=UTF-8"},
		},
		Body: requestyBody,
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error occurred ", err)
		panic(err)
	}

	data, _ := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(res.Request.Header.Get("Content-Type"))
	fmt.Println("Status ", res.Status)
	fmt.Printf("%s\n ", data)
}

func makingCustomizedRequest2() {
	// create request body
	reqBody := strings.NewReader(`
		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}
	`)

	// create a request object
	req, _ := http.NewRequest(
		"POST",
		"http://dummy.restapiexample.com/api/v1/create",
		reqBody,
	)

	// add a request header
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	// send an HTTP using `req` object
	res, err := http.DefaultClient.Do(req)

	// check for response error
	if err != nil {
		log.Fatal("Error:", err)
	}

	// read response body
	data, _ := io.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// print response status and body
	fmt.Printf("status: %d\n", res.StatusCode)
	fmt.Printf("body: %s\n", data)
}

func makingCustomizedRequest3() {
	client := http.Client{
		Timeout: 100 * time.Second,
	}

	res, err := client.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {

		// get `url.Error` struct pointer from `err` interface
		urlErr := err.(*url.Error)

		// check if error occurred due to timeout
		if urlErr.Timeout() {
			fmt.Println("Error occurred due to a timeout.")
		}

		// log error and exit
		log.Fatal("Error:", err)
	} else {
		fmt.Println("Success: status-code", res.StatusCode)
	}

}
