package models

import (
	"github.com/govind/golang2/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Emp struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Designation string `gorm:"column:designation" json:"designation"`
	Manager     string `gorm:"column:manager" json:"manager"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Emp{})
}

func (b *Emp) CreateEmp() *Emp {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllEmps() []Emp {
	var Emps []Emp
	db.Find(&Emps)
	return Emps
}

func GetEmpById(Id int64) (*Emp, *gorm.DB) {
	var getEmp Emp
	db := db.Where("ID=?", Id).Find(&getEmp)
	return &getEmp, db
}

func DeleteEmp(ID int64) Emp {
	var emp Emp
	db.Where("ID=?", ID).Delete(emp)
	return emp
}
