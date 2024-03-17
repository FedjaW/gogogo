package main

import (
	"fmt"
	"html/template"
	"net/http"

	"leckomio.dev/go/museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.ParseFiles("templates/index.tmpl")
	// template.Must()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // set http status before write to the body, http works like a buffer, can not change status after body
		w.Write([]byte("Internal Server Error"))
		return
	}
	htmlTemplate.Execute(w, data.GetAll()[0])
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello) // passing a func as an argument, not executing handleHello()
	server.HandleFunc("/template", handleTemplate)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs) // Handle takes a type, therefor nor HandleFunc

	// will open a port in your os for incoming connections over tcp
	// domain:port
	// if you do not specify a host or ip address
	// it will open that port on any network connection that you have on your computer
	fmt.Println("start listening on port 3333")
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("error while opening the server")
	}
}
