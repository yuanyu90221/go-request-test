package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Status  string    `json:status`
	Message string    `json:message`
	Assets  []ResData `json:assets`
	Total   int       `json:total`
}
type ResData struct {
	IsLock       bool   `json:is_lock`
	Presign_Link string `json:presign_link`
	Name         string `json:name`
}

func main() {
	res, err := http.Get("https://resource.rplab.online/default_assets")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonBody := &Response{}
	err = json.Unmarshal(body, jsonBody)
	if err != nil {
		log.Fatal(err)
		return
	}
	//fmt.Printf("%s\n", body)
	fmt.Println(jsonBody.Assets)
}
