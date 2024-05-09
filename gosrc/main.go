package main

import (
	"fmt"
	"log"
	"net/http"

	"employeeManager/gosrc/routes"
	"employeeManager/gosrc/service"

	"github.com/gorilla/mux"
)

func main() {
	store := service.NewEmployeeStore()
	router := mux.NewRouter()

	// routes calls
	routes.RouteCalls(router, store)

	fmt.Println("Listening on server port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
