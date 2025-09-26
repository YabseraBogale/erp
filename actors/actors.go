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
	Password                string
}

// type transaction "add","remove","checkout","checkin"

type Item struct {
	gorm.Model
	ItemId              int `gorm:"primaryKey"`
	EmployeeId          int `gorm:"foreignKey"`
	ItemName            string
	ItemDescription     string
	ItemPrice           float32
	ItemQuantity        int
	Categories          string
	SubCategories       string
	Location            string
	Unit                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CreatedByEmployeeId int `gorm:"foreignKey"`
}

type CheckOut struct {
	gorm.Model
	CheckOutId          int `gorm:"primaryKey"`
	ItemId              int `gorm:"foreignKey"`
	InventoryEmployeeId int
	CheckOutEmployeeId  int
	ItemQuantity        int
	ItemSiv             int
	CheckOutDate        time.Time
	Notes               string
}

type CheckIn struct {
	gorm.Model
	CheckInId           int `gorm:"primaryKey"`
	ItemId              int `gorm:"foreignKey"`
	InventoryEmployeeId int
	CheckInEmployeeId   int
	ItemQuantity        int
	ItemPrice           int
	ItemGrr             int
	ReturnDate          time.Time
	Notes               string
}
