package main

import (
	"fmt"
	"go-server-test/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routes.Init()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version: v%s", "1.0.0")
	})

	port := os.Getenv("port")
	if port == "" {
		port = "5000"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
