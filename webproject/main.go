package main

import (
	"fmt"
	"net/http"
)

// Write to the response body
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Welcome to this cool site!</h1>")
}

// Write contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1>p>To Get in contact, e-mail me at <a
	href=\"mailto:c_hagglund@bahnhof.se\"<c_hagglund@bahnhof.se</a>.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
