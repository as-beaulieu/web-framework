package main

import (
	"net/http"
	"os"
)

func main() {
	port := initArgs(os.Args)
	console("Running Program")
	console("Listening on localhost ", port)

	http.ListenAndServe(port, nil)
}
