package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {

	w.WriteHeader(status)

	if status == 404 {
		fmt.Fprintf(w, "Custom 404: not found!")
	}
}

func ReplyHello(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() // parse arguments, you have to call this by yourself
	/*fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])*/
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world version 6!") // send data to client side

}

func main() {

	http.HandleFunc("/", ReplyHello)         // set router
	err := http.ListenAndServe(":9090", nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
