package models

import "time"

// Car - schema for create car data
type Car struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Vin       string `json:"vin" gorm:"unique;not null"`
	Model     string `json:"model" gorm:"size:100"`
	Make      string `json:"make" gorm:"size:100"`
	Year      uint   `json:"year" gorm:"not null"`
	Msrp      uint   `json:"msrp"`
	Status    string `json:"status" gorm:"default:'in_stock'"`
	Booked    bool   `json:"booked"`
	Listed    bool   `json:"listed"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InputCarUpdate - schema for create/update car data
type InputCarUpdate struct {
	Vin    string `json:"vin"`
	Model  string `json:"model"`
	Make   string `json:"make"`
	Year   uint   `json:"year"`
	Msrp   uint   `json:"msrp"`
	Status string `json:"status"`
	Booked bool   `json:"booked"`
	Listed bool   `json:"listed"`
}
