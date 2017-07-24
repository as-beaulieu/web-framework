package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Args[1]
	listen := "Listening on localhost " + port
	console("Running Program")
	console(listen)

	//Default: program call is Args[0]
	//Test: Checking all arguments passed after program call
	args := os.Args[1:]
	fmt.Println(args)

	http.ListenAndServe(":8080", nil)
}
