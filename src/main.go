package main

import "fmt"
import "net/http"

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Intersect")

}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
