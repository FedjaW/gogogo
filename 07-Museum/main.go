package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from go"))
		})

	// will open a port in your os for incoming connections over tcp
	// domain:port
	// if you do not specify a host or ip address
	// it will open that port on any network connection that you have on your computer
	fmt.Println("start listening on port 3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("error while opening the server")
	}
}
