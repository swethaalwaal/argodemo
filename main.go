package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Came to handler")
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Started serviced on port 8001")
	http.ListenAndServe(":8001", nil)
}
