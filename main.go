package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamtonmoy0/movie-crud-api/routes"
)

func main() {
	fmt.Println("Server is started .....")
	log.Fatal(http.ListenAndServe(":4000", routes.Router()))
	fmt.Println("listening at port 4000")
}
