package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:3000"
	fmt.Println(url)
	type User struct {
		Name string
	}
	usr := User{
		Name: "Prabal",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(usr)
	if err != nil {
		log.Fatal("Error while encoding the json: ", err)
	}
	request, requestErr := http.NewRequest("POST", url, &buf)
	if requestErr != nil {
		log.Fatal("Error while creating the request: ", requestErr)
	}
	client := &http.Client{}
	resp, respErr := client.Do(request)
	if respErr != nil {
		log.Fatal("Hey you did something wrong")
	}
	bytes, bytesErr := io.ReadAll(resp.Body)
	if bytesErr != nil {
		log.Fatal("Read error: ", err)
	}
	fmt.Println(string(bytes))
}
