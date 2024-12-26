package models

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"type:varchar(100);not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Currency  string    `gorm:"type:varchar(10);not null" json:"currency"`
	Balance   float64   `gorm:"type:decimal(10,2);default:0.0" json:"balance"`
	Status    string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	AccountInformation []AccountAddressInformation `gorm:"foreignKey:AccountID" json:"account_information"`
	Transactions       []Transaction               `gorm:"foreignKey:FromAccountID" json:"transactions"`
}
