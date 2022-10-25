<<<<<<< HEAD
package main ()
=======
package main
>>>>>>> e021c3d578e4f979077bd895bce4092fa6ed9c56

import (
	"fmt"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to LinsEYe!</h1>")
}
<<<<<<< HEAD

//error training
//err := doStuff()
=======
//error training
err := doStuff()
>>>>>>> e021c3d578e4f979077bd895bce4092fa6ed9c56
func main() {
	http.HandleFunc("/", HandlerFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}

//Debugging example
<<<<<<< HEAD
//./main.go:11:6: no new variables on left side of :=
=======

./main.go:11:6: no new variables on left side of :=
>>>>>>> e021c3d578e4f979077bd895bce4092fa6ed9c56
