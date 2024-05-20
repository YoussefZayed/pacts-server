package main

import (
	"fmt"
	"net/http"

	"pacts/server/db"

	"github.com/gorilla/mux"
)

func main() {
	defer db.DB.Close()

	http.HandleFunc("/")
	r := mux.NewRouter()

	fmt.Println("Server listening on :6080")
	http.ListenAndServe(":6080", r)
}
