package config_test

import (
	"testing"

	"github.com/govind/golang2/pkg/config"
)

func TestDBConnection(t *testing.T) {
	config.Connect()

	db := config.GetDB()

	if db == nil {
		t.Error("Expected a non-nil database instance, got nil")
	}
	if err := db.DB().Ping(); err != nil {
		t.Errorf("Failed to ping the database: %v", err)
	}
}
