package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server started on port 8080 !")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		fmt.Println("server started on port 8080 !")

	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "error 404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "PareForm() err :", err)
		return

	}
	fmt.Fprintln(w, "Post request successfull")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintln(w, "Name = ", name)
	fmt.Fprintln(w, "Address =", address)
}
