package service

import (
	"fmt"
	"github.com/bahodurnazarov/wallet/pkg/models"
	"strconv"
)

func (s *Service) CheckWallet(wallet string) (user models.User, err error) {
	if len(wallet) < 9 || len(wallet) > 12 {
		return user, fmt.Errorf("Invalid Wallet Number")
	}
	user, err = s.Repository.CheckWallet(wallet)
	return user, err
}

func (s *Service) Replenishment(user models.User) (newUser models.User, err error) {
	if len(strconv.Itoa(user.Wallet)) < 9 || len(strconv.Itoa(user.Wallet)) > 12 {
		return user, fmt.Errorf("Invalid Wallet Number")
	}

	newUser, err = s.Repository.Replenishment(user)
	return newUser, err
}
func (s *Service) GetTransactions(wallet string) (replenishment []models.Transactions, err error) {
	if len(wallet) < 9 || len(wallet) > 12 {
		return replenishment, fmt.Errorf("Invalid Wallet Number")
	}
	replenishment, err = s.Repository.GetTransactions(wallet)
	return replenishment, err
}
func (s *Service) Balance(wallet string) (user models.User, err error) {
	if len(wallet) < 9 || len(wallet) > 12 {
		return user, fmt.Errorf("Invalid Wallet Number")
	}
	user, err = s.Repository.Balance(wallet)
	return user, err
}
