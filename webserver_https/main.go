package main

import (
	"log"
	"net/http"
	"os"
)

/* func Output(calldepth int, x string) error {
	return std.Output(calldepth+1, x) // +1 for this frame.
} */

func main() {

	f, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))

}
