package main

import (
	"example/web-service-gin/controller"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
