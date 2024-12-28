package service

import (
	databases "payment/database"
	"payment/models"
)

func CreateAccountNumber(account models.Account, isRender bool, number string) {
	var numberAddress string
	if isRender {
		numberAddress = RenderNumberAccount()
	} else {
		numberAddress = number
	}
	// check account number is valid and avaliable
	if !CheckAccountNumber(numberAddress) {
		// return error
		panic("Account number is not valid")
	}
	// save account number to database

	db := databases.GetWriteInstance()
	var accountNumber models.AccountAddressInformation
	accountNumber.Account = account
	accountNumber.Address = numberAddress
	if result := db.Create(&accountNumber); result.Error != nil {
		panic("Failed to save account number")
	}
}

func RenderNumberAccount() string { return "" }

func CheckAccountNumber(accountNumber string) bool { return true }
