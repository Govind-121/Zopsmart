package routes

import (
	"github.com/gorilla/mux"
	"github.com/govind/golang2/pkg/controllers"
)

var RegisterEmployeeRoutes = func(router *mux.Router) {
	router.HandleFunc("/emp/", controllers.CreateEmp).Methods("POST")
	router.HandleFunc("/emp/", controllers.GetEmp).Methods("GET")
	router.HandleFunc("/emp/{empId}", controllers.GetEmpById).Methods("GET")
	router.HandleFunc("/emp/{empId}", controllers.UpdateEmp).Methods("PUT")
	router.HandleFunc("/emp/{empId}", controllers.DeleteEmp).Methods("DELETE")
}
