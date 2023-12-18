package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/govind/golang2/pkg/routes"
)

func TestMainRoutes(t *testing.T) {
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r)
	testServer := httptest.NewServer(r)
	defer testServer.Close()

	req, err := http.NewRequest("GET", testServer.URL+"/your-endpoint-path", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
