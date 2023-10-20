package repository

import (
	"errors"
	"fmt"
	"github.com/bahodurnazarov/wallet/pkg/models"
	"gorm.io/gorm"
	"log"
	"time"
)

const maxAmount = 100000
const minAmount = 10000

func (r *Repository) CheckWallet(wallet string) (user models.User, err error) {
	if err = r.Conn.First(&user, "wallet = ?", wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("walllet not found")
		}
		return user, err
	}

	return user, nil
}

func (r *Repository) Replenishment(user models.User) (newUser models.User, err error) {
	if err = r.Conn.First(&newUser, "wallet = ?", user.Wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return newUser, fmt.Errorf("wallet not found")
		}
		return newUser, err
	}

	newUser.Balance += user.Balance

	if newUser.Identified == true && newUser.Balance > maxAmount {
		return newUser, fmt.Errorf("amount should be less than 100000")
	} else if newUser.Identified == false && newUser.Balance > minAmount {
		return newUser, fmt.Errorf("amount should be less than 10000")
	}

	r.Conn.Save(&newUser)
	var replenishment models.Transactions
	replenishment.WalletID = newUser.ID
	replenishment.Wallet = newUser.Wallet
	replenishment.Amount = newUser.Balance
	replenishment.TypeOfTransaction = "Пополнение"

	if result := r.Conn.Create(&replenishment).Error; result != nil {
		return user, result
	}
	return newUser, err
}

func (r *Repository) GetTransactions(wallet string) (replenishment []models.Transactions, err error) {
	var sum float64
	//var count int64
	result := r.Conn.Select("count(wallet), sum(amount)").Where("wallet = ? AND EXTRACT(MONTH FROM created_at) = ?", wallet, time.Now().Month()).Find(&replenishment)
	//result := r.Conn.Model(&replenishment).
	//	Where("wallet = ? AND TO_DATE(created_at) = ?", wallet, int(time.Now().Month())).
	//	Select("SUM(amount)").
	//	Scan(&sum)
	if result != nil {
		//r.Conn.Model(&replenishment).Where("wallet = ?", wallet).Select("SUM(amount)").Scan(&count).Error; err != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return replenishment, result.Error
		}
		return replenishment, result.Error
	}
	log.Println("1111111111111111111", sum)
	return replenishment, result.Error
}

func (r *Repository) Balance(wallet string) (user models.User, err error) {
	if err = r.Conn.First(&user, "wallet = ?", wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("walllet not found")
		}
		return user, err
	}

	return user, nil
}
