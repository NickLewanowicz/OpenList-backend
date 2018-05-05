package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Init Router
	fmt.Println("Server starting up...")
	r := NewRouter()
	initDb("openlist")
	log.Fatal(http.ListenAndServe(":8000", r))
}
