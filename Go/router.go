package main

import (
	"net/http"
)

func router() {
	//Root Handler
	http.HandleFunc("/", RootHandler)

	//Static Handlers, using html files

	//Router Adapters

	//Dynamic Handlers using gotmpl files

	//Serving local assets
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css"))))
	http.Handle("/assets/fonts/", http.StripPrefix("/assets/fonts/", http.FileServer(http.Dir("assets/fonts"))))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("assets/img"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("assets/js"))))
}
