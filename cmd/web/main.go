package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rcliao/non-player-contest/web"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", web.Hello()).Methods("GET")

	log.Println("Running web server at port 9000")
	http.ListenAndServe(":9000", r)
}
