package models_test

import (
	"testing"

	"github.com/govind/golang2/pkg/config"
	"github.com/govind/golang2/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestEmpModelFunctions(t *testing.T) {
	config.Connect()
	newEmp := &models.Emp{
		Name:        "John Doe",
		Designation: "Developer",
		Manager:     "Alice",
	}

	createdEmp := newEmp.CreateEmp()

	assert.NotNil(t, createdEmp)
	assert.NotZero(t, createdEmp.ID)

	allEmps := models.GetAllEmps()

	assert.NotNil(t, allEmps)
	assert.NotEmpty(t, allEmps)

	retrievedEmp, _ := models.GetEmpById(createdEmp.ID)

	assert.NotNil(t, retrievedEmp)
	assert.Equal(t, createdEmp.ID, retrievedEmp.ID)

	deletedEmp := models.DeleteEmp(createdEmp.ID)

	assert.NotNil(t, deletedEmp)
	assert.Equal(t, createdEmp.ID, deletedEmp.ID)
}
