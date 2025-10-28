package models

type Driver struct {
	Id            int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Vehicle       string `json:"vehicle"`
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
	Email         string `json:"email" gorm:"uniqueIndex"`
	Password      string `json:"password"`
	Phone         string `json:"phone_number"`
}
