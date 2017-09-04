package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello Github")
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities' : '北京', '上海', '广州', '赤峰'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
	fmt.Println(" '北京', '上海', '广州', '赤峰' ")
}
