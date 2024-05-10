package main

import (
	"finance-go/src/config"
	"finance-go/src/router"
	"log"
	"net/http"
)

func main() {
	db := config.Config()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":3000", router.Configure()))
}
