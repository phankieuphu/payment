package models

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	FromAccountID   uint      `gorm:"not null" json:"from_account_id"`
	ToAccountID     *uint     `json:"to_account_id"` // Nullable for external transfers
	Address         string    `gorm:"type:varchar(50)" json:"address"`
	Amount          float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Currency        string    `gorm:"type:varchar(10);not null" json:"currency"`
	TransactionType string    `gorm:"type:varchar(20);not null" json:"transaction_type"`
	Status          string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Description     string    `gorm:"type:varchar(255)" json:"description"`
	ExternalBankID  *uint     `json:"external_bank_id"` // Nullable
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	FromAccount  Account       `gorm:"foreignKey:FromAccountID" json:"from_account"`
	ToAccount    *Account      `gorm:"foreignKey:ToAccountID" json:"to_account"`
	ExternalBank *ExternalBank `gorm:"foreignKey:ExternalBankID" json:"external_bank"`
}
