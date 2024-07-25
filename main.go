package main

import (
	"fmt"
	storage "learn/db"
	routing "learn/routes"
	"net/http"
)

func main() {
	routing.InitializeRoutes()
	mux := routing.Mux
	storage.Db.InitDb()
	fmt.Println(storage.Db)
	http.ListenAndServe(":8080", mux)
}
