package models

import "time"

type Trip struct {
	Id          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId      int64     `json:"user_id"`
	DriverId    int64     `json:"driver_id"`
	PickupLat   float64   `json:"pick_up_lat"`
	PickupLong  float64   `json:"pick_up_long"`
	DropoffLat  float64   `json:"drop_off_lat"`
	DropoffLong float64   `json:"drop_off_long"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
