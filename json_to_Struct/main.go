package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Marshal, unmarshal is for string
//Decode, encode is for streams, https response body is a stream, so Decode

var client *http.Client

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type User struct {
	Results []UserResult
}

type UserResult struct {
	Name    UserName
	Email   string
	Picture UserPicture
}

type UserName struct {
	Title string
	First string
	Last  string
}

type UserPicture struct {
	Large     string
	Medium    string
	Thumbnail string
}

func GetJson(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"

	var catFact CatFact

	err := GetJson(url, &catFact)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(catFact)
	}
}

func GetRandomUser() {
	url := "https://randomuser.me/api/?inc=name,email,picture"

	var user User

	err := GetJson(url, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	GetCatFact()

	catFact := CatFact{
		Fact:   "Cat Attribute",
		Length: 13,
	}

	jsonStr, err := json.Marshal(catFact)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Json String is ", string(jsonStr))
	}

	GetRandomUser()
}
