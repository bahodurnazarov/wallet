package service

import "github.com/bahodurnazarov/wallet/internal/repository"

type Service struct {
	Repository *repository.Repository
}

func New(rep *repository.Repository) *Service {
	return &Service{Repository: rep}
}
