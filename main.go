package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Started serviced on port 8001")
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/index", IndexHandler)
	http.ListenAndServe(":8001", nil)
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Health Check Success!"))
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to, %s!", r.URL.Path[1:])
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Came to index page")
	w.Header().Add("Content Type", "text/html")
	http.ServeFile(w, r, "index.html")
}
