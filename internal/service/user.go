package service

import (
	"fmt"
	"github.com/bahodurnazarov/wallet/pkg/models"
	"strconv"
)

func (s *Service) CreateUser(user models.User) error {
	if len(strconv.Itoa(user.Wallet)) < 9 || len(strconv.Itoa(user.Wallet)) > 12 {
		return fmt.Errorf("Invalid Phone Number")
	}
	if user.Balance != 0 {
		return fmt.Errorf("Balance should be 0")
	}
	return s.Repository.CreateUser(user)
}
