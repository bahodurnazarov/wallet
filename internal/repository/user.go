package repository

import (
	"fmt"
	"github.com/bahodurnazarov/wallet/pkg/models"
)

func (r *Repository) CreateUser(user models.User) error {
	if result := r.Conn.Create(&user).Error; result != nil {
		return fmt.Errorf("wallet already exist")
	}
	return nil
}
