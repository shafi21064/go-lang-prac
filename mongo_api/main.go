package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shafi21064/go_mongo_api/router"
)

func main() {
	fmt.Println("Mongo db api")
	r := router.Router()
	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server started at 8000")
}
