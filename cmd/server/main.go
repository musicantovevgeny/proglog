package main

import (
	"log"

	"github.com/musicantovevgeny/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
