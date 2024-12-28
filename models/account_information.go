package models

import "time"

type AccountAddressInformation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AccountID uint      `gorm:"not null" json:"account_id"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address"`
	Type      string    `gorm:"type:varchar(20);default:'primary'" json:"type"`
	Status    string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Account   Account   `gorm:"foreignKey:AccountID" json:"account"`
}
