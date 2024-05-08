package main

import (
	"fmt"
	"github.com/eduardoraider/go-crud/internal/api"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", api.NewRouter()))
}
