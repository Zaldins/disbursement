package models

import (
	"gorm.io/gorm"
)

type (
	Disbursement struct {
		gorm.Model
		WalletID uint    `gorm:"wallet_id"`
		Amount   float64 `gorm:"amount"`
		Status   string  `gorm:"status"`
		Wallet   Wallet
	}

	Wallet struct {
		gorm.Model
		ID      int     `gorm:"primaryKey"`
		Balance float64 `gorm:"balance"`
		Name    string  `gorm:"name"`
	}
)