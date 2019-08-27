package main

import (
	"log"
	"net/http"
	"root3/router"
)

func main() {
	log.Println("Listening !")
	router.Routers()
	http.ListenAndServe(":8080", router.Mux)
}
