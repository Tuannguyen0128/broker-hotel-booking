package models

import "time"

type Account struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Fullname     string    `gorm:"size:20;not null" json:"fullname"`
	Email        string    `gorm:"size:50;not null;unique" json:"email"`
	Password     string    `gorm:"size:60;not null" json:"password"`
	CreatedAt    time.Time `gorm:"autoCreatedTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdatedTime" json:"updated_at"`
	MerchantCode string    `json:"merchantcode"`
}
