package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SUCCESS -- API Running")
}

func main() {
	http.HandleFunc("/", root)

	http.ListenAndServe(":8090", nil)
}