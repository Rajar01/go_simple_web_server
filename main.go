package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintf(w, "<h1>Hello From Golang Simple Web Server</h1>")
		fmt.Fprintf(w, "<a href='/'>Go Back To Homepage</a>")
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "ParseForm err : %v", err)
		log.Fatal(err)
		return
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintf(w, "Form Successfully Submitted\n\n")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		age := r.FormValue("age")

		fmt.Fprintf(w, "<h1>My firstname is %s and my lastname is %s. In this year i am %s years old</h1>", firstname, lastname, age)

		fmt.Fprintf(w, "<a href='/'>Go Back To Homepage</a>")
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
