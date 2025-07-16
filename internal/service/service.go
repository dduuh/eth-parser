package service

import (
	"context"
	"eth-parser/internal/domain"
	"eth-parser/internal/repository"
	"fmt"
)

type addresses interface {
	AddAddress(ctx context.Context, addr domain.Addresses) (domain.Addresses, error)
}

type Service struct {
	addrRepo *repository.Addresses
	addrs    addresses
}

func New(addrRepo *repository.Addresses) *Service {
	return &Service{
		addrRepo: addrRepo,
	}
}

func (s *Service) AddAddress(ctx context.Context, addr domain.Addresses) (domain.Addresses, error) {
	createdAddr, err := s.addrs.AddAddress(ctx, addr)
	if err != nil {
		return domain.Addresses{}, fmt.Errorf("failed to add the address: %w", err)
	}

	return createdAddr, nil
}