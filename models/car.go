package models

// Car - model for cars
type Car struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Vin    string `json:"vin"`
	Model  string `json:"model"`
	Make   string `json:"make"`
	Year   uint   `json:"year"`
	Msrp   uint   `json:"msrp" gorm:"default:0"`
	Status string `json:"status" gorm:"default:in_stock"`
	Booked bool   `json:"booked"`
	Listed bool   `json:"listed"`
}
