package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ishanshre/GO-Stocks-API/pkg/router"
)

const url = "127.0.0.1:8000"

func main() {
	r := router.Router()
	fmt.Println("Starting server at: ", url)
	log.Fatal(http.ListenAndServe(url, r))
}
