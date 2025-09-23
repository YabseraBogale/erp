package actors

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeContact struct {
	gorm.Model
	EmergencyContactFyidaId string `gorm:"primaryKey"`
	FirstName               string
	MiddleName              string
	LastName                string
	PhoneNumber             string
	Email                   string
	Location                string
	Gender                  string
}

type Employee struct {
	gorm.Model
	EmployeeId              int    `gorm:"primaryKey"`
	EmergencyContactFyidaId string `gorm:"foreignKey"`
	FirstName               string
	MiddleName              string
	LastName                string
	PhoneNumber             string
	Email                   string
	FyidaId                 string
	Position                string
	Department              string
	TinNumber               int
	Location                string
	BankAccountNumber       int
	Salary                  float32
	Gender                  string
	JobDescription          string
	DateOfEmployement       time.Time
}

type Item struct {
	gorm.Model
	ItemId              int `gorm:"primaryKey"`
	EmployeeId          int `gorm:"foreignKey"`
	ItemName            string
	ItemDescription     string
	ItemPrice           string
	ItemQuantity        string
	Categories          string
	SubCategories       string
	Location            string
	Unit                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CreatedByEmployeeId int `gorm:"foreignKey"`
}

type TransactionType struct {
	// type "add","remove","checkout","checkin"
	gorm.Model
	TransactionTypeId int
	TypeName          string
}

type ItemLog struct {
	gorm.Model
	LogId           int    `gorm:"primaryKey"`
	ItemId          int    `gorm:"foreignKey"`
	TransactionName string `gorm:"foreignKey"`
	EmployeeId      int    `gorm:"foreignKey"`
	QuantityChanged int
	ItemPrice       float32
	TransactionDate time.Time
	Description     string
}

type CheckOut struct {
	gorm.Model
	CheckOutId   int `gorm:"primaryKey"`
	ItemId       int `gorm:"foreignKey"`
	EmployeeId   int `gorm:"foreignKey"`
	CheckOutDate time.Time
	Notes        string
}

type CheckIn struct {
	gorm.Model
	CheckInId  int `gorm:"primaryKey"`
	ItemId     int `gorm:"foreignKey"`
	EmployeeId int `gorm:"foreignKey"`
	ReturnDate time.Time
	Notes      string
}
