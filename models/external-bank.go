package models

import "time"

type ExternalBank struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	APIURL          string    `gorm:"type:varchar(255);not null" json:"api_url"`
	AuthCredentials string    `gorm:"type:varchar(255)" json:"auth_credentials"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Transactions []Transaction `gorm:"foreignKey:ExternalBankID" json:"transactions"`
}

func (e *ExternalBank) TableName() string {
	return "external_banks"
}
