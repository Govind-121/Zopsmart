package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/govind/golang2/pkg/models"
	"github.com/govind/golang2/pkg/utils"
)

var NewEmp models.Emp

func GetEmp(w http.ResponseWriter, r *http.Request) {
	newEmps := models.GetAllEmps()
	res, _ := json.Marshal(newEmps)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func GetEmpById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empId"]
	ID, err := strconv.ParseInt(empId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	empDetails, _ := models.GetEmpById(ID)
	res, _ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	CreateEmp := &models.Emp{}
	utils.ParseBody(r, CreateEmp)
	b := CreateEmp.CreateEmp()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empId"]
	ID, err := strconv.ParseInt(empId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	emp := models.DeleteEmp(ID)
	res, _ := json.Marshal(emp)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	var updateEmp = &models.Emp{}
	utils.ParseBody(r, updateEmp)
	vars := mux.Vars(r)
	empId := vars["empId"]
	ID, err := strconv.ParseInt(empId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	empDetails, db := models.GetEmpById(ID)
	if updateEmp.Name != "" {
		empDetails.Name = updateEmp.Name
	}
	if updateEmp.Designation != "" {
		empDetails.Designation = updateEmp.Designation
	}
	if updateEmp.Manager != "" {
		empDetails.Manager = updateEmp.Manager
	}
	db.Save(&empDetails)
	res, _ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
