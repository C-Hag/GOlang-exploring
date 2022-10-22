package main

import (
	"fmt"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to LinsEYe!</h1>")
}

//error training
//err := doStuff()
func main() {
	http.HandleFunc("/", HandlerFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}

//Debugging example
//./main.go:11:6: no new variables on left side of :=
