package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/govind/golang2/pkg/controllers"
	"github.com/govind/golang2/pkg/models"
	"github.com/jinzhu/gorm"
)

func TestGetEmpById(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/employees/{empId}", controllers.GetEmpById).Methods("GET")

	mockServer := httptest.NewServer(router)
	defer mockServer.Close()

	empID := "1"
	req, err := http.NewRequest("GET", mockServer.URL+"/employees/"+empID, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var emp models.Emp
	if err := json.NewDecoder(resp.Body).Decode(&emp); err != nil {
		t.Errorf("Error decoding JSON response: %s", err)
	}
	expectedEmp := models.Emp{}
	if emp != expectedEmp {
		t.Errorf("Expected %+v, got %+v", expectedEmp, emp)
	}

}

func TestGetEmp(t *testing.T) {
	mockEmps := []models.Emp{
		{Name: "John Doe", Designation: "Developer", Manager: "Alice"},
		{Name: "Jane Smith", Designation: "Designer", Manager: "Bob"},
	}

	models.GetAllEmps = func() []models.Emp {
		return mockEmps
	}

	router := mux.NewRouter()
	router.HandleFunc("/employees", controllers.GetEmp).Methods("GET")

	mockServer := httptest.NewServer(router)
	defer mockServer.Close()

	req, err := http.NewRequest("GET", mockServer.URL+"/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var receivedEmps []models.Emp
	if err := json.NewDecoder(resp.Body).Decode(&receivedEmps); err != nil {
		t.Errorf("Error decoding JSON response: %s", err)
	}

	if !reflect.DeepEqual(receivedEmps, mockEmps) {
		t.Errorf("Expected %+v, got %+v", mockEmps, receivedEmps)
	}

}

func TestCreateEmp(t *testing.T) {

	newEmp := models.Emp{
		Name:        "New Emp",
		Designation: "Intern",
		Manager:     "Manager X",
	}

	models.CreateEmp = func(emp *models.Emp) models.Emp {
		newEmp.ID = 3
		return newEmp
	}

	payload, err := json.Marshal(newEmp)
	if err != nil {
		t.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/employees", controllers.CreateEmp).Methods("POST")

	mockServer := httptest.NewServer(router)
	defer mockServer.Close()

	req, err := http.NewRequest("POST", mockServer.URL+"/employees", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var createdEmp models.Emp
	if err := json.NewDecoder(resp.Body).Decode(&createdEmp); err != nil {
		t.Errorf("Error decoding JSON response: %s", err)
	}

	if !reflect.DeepEqual(createdEmp, newEmp) {
		t.Errorf("Expected %+v, got %+v", newEmp, createdEmp)
	}
}

func TestDeleteEmp(t *testing.T) {
	mockEmp := models.Emp{
		Name:        "John Doe",
		Designation: "Developer",
		Manager:     "Alice",
	}

	models.DeleteEmp = func(id int64) models.Emp {
		if id != mockEmp.ID {
			return models.Emp{}
		}
		return mockEmp
	}

	router := mux.NewRouter()
	router.HandleFunc("/employees/{empId}", controllers.DeleteEmp).Methods("DELETE")

	mockServer := httptest.NewServer(router)
	defer mockServer.Close()

	empID := "1"
	req, err := http.NewRequest("DELETE", mockServer.URL+"/employees/"+empID, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var deletedEmp models.Emp
	if err := json.NewDecoder(resp.Body).Decode(&deletedEmp); err != nil {
		t.Errorf("Error decoding JSON response: %s", err)
	}

	if deletedEmp != mockEmp {
		t.Errorf("Expected %+v, got %+v", mockEmp, deletedEmp)
	}
}

func TestUpdateEmp(t *testing.T) {
	updateData := models.Emp{
		Name:        "Updated Name",
		Designation: "Updated Designation",
		Manager:     "Updated Manager",
	}

	models.GetEmpById = func(id int64) (models.Emp, *gorm.DB) {
		if id != updateData.ID {
			return models.Emp{}, nil
		}
		return models.Emp{
			Name:        "John Doe",
			Designation: "Developer",
			Manager:     "Alice",
		}, &gorm.DB{}
	}
	savedEmp := updateData
	models.Save = func(emp *models.Emp) *gorm.DB {

		savedEmp = *emp
		return &gorm.DB{}
	}

	router := mux.NewRouter()
	router.HandleFunc("/employees/{empId}", controllers.UpdateEmp).Methods("PUT")

	mockServer := httptest.NewServer(router)
	defer mockServer.Close()

	payload, err := json.Marshal(updateData)
	if err != nil {
		t.Fatal(err)
	}

	empID := "1"
	req, err := http.NewRequest("PUT", mockServer.URL+"/employees/"+empID, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var updatedEmp models.Emp
	if err := json.NewDecoder(resp.Body).Decode(&updatedEmp); err != nil {
		t.Errorf("Error decoding JSON response: %s", err)
	}

	if !reflect.DeepEqual(savedEmp, updateData) {
		t.Errorf("Expected %+v, got %+v", updateData, savedEmp)
	}

}
