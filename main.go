package main

import (
	"bookMyTicket/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Starting BookMyTicket server...")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to BookMyTicket"))

	}).Methods("GET")
	r.HandleFunc("/shows/{name}", routes.GetShowsRouteByID()).Methods("GET")

	r.Handle("/shows", routes.ShowsRoute()).Methods("GET")
	fmt.Println("HERE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
	fmt.Println("Server is running")
	log.Println("Server is running on port 8080")
}
