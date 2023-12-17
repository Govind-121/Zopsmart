package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govind/golang2/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	//routes.RegisterEmpStoreRoutes(r)
	routes.RegisterEmployeeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:9010", r))

}
