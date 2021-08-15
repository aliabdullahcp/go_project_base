package main

import (
	"fmt"
	"net/http"

	"github.com/aliabdullahcp/go_project_base/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := handlers.PrintMain()
		fmt.Fprintf(w, result)

	})

	fmt.Println("Server listening!")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}
