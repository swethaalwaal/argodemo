package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Started service!")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/index", IndexPage)
	if errl:=http.ListenAndServe(":8001", nil);errl!=nil{
		log.Fatalf("Unable to start the server due to :%s",errl.Error())
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to, %s!", r.URL.Path[1:])
}
func IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Came to index page")
	w.Header().Add("Content Type", "text/html")
	http.ServeFile(w, r, "index.html")
}
