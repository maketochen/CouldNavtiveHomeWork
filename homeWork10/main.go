package main

import (
	"httpservice/handler"
	"log"
	"net/http"
)

func main() {
	handler.RegistryHttpRoutine()
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
