package main

import (
	"log"
	"net/http"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data := `{"cities":"San Francisco, Amsterdam, Berlin, Rotterdam, Tokyo"}`
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write([]byte(data))
}

func main() {
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
