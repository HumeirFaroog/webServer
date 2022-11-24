package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// my main function where I can   excuate anythings

func main() {
	fmt.Println(" HERE  WIL  SHOW HOW TO  MAKE GET REQUEST ")
	makeGet() //  calling the function

}

func makeGet() {
	const url = "http://localhost:8080/events"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //   our mission to  close it

	fmt.Println("Here we can see the status ", response.StatusCode)
	fmt.Println("Here we can see the content length ", response.ContentLength)

	// need to  read this response

	var responseStr strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCon, _ := responseStr.Write(content)
	fmt.Println("ByetCont is  : ", byteCon)
	fmt.Println(responseStr.String())
	//fmt.Println(string(content))
	// fmt.Println(content)

}
