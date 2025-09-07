package main

import (
	"fmt"
	"go-postgress/config"
	"go-postgress/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.Handle("/", r)
	config.Connect()
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
