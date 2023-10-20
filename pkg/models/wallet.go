package models

import "time"

type User struct {
	ID         uint       `gorm:"type:numeric;primary_key"`
	Wallet     int        `gorm:"type:numeric;uniqueIndex:wallet;not null"`
	Identified bool       `gorm:"type:bool;default:false"`
	Balance    float64    `gorm:"type:decimal(10,2);not null"`
	CreatedAt  *time.Time `gorm:"not null;default:now()"`
}

type Transactions struct {
	ID                uint       `gorm:"type:numeric;primary_key"`
	WalletID          uint       `json:"user_id"`
	Wallet            int        `gorm:"type:numeric;not null"`
	Amount            float64    `gorm:"type:decimal(10,2);not null"`
	TypeOfTransaction string     `gorm:"type:varchar;not null"`
	CreatedAt         *time.Time `gorm:"not null;default:now()"`
}
