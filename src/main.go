package main

import (
	"log"
	"net/http"
	"os"
)

var PORT = os.Getenv("PORT")

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello wolrd\n"))
	})

	log.Print(`Running on port: ` + PORT)
	panic(http.ListenAndServe(":"+PORT, nil))
}
