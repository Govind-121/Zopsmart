package utils_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/govind/golang2/pkg/utils"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type Emp struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Designation string `gorm:"column:designation" json:"designation"`
	Manager     string `gorm:"column:manager" json:"manager"`
}

func TestParseBody(t *testing.T) {
	testPayload := TestData{
		Name: "John Doe",
		Age:  30,
	}

	payload, _ := json.Marshal(testPayload)

	req, _ := http.NewRequest("POST", "/test", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	var data TestData
	utils.ParseBody(req, &data)

	assert.Equal(t, testPayload.Name, data.Name)
	assert.Equal(t, testPayload.Age, data.Age)

}
