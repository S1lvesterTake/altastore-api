package db

import (
	"altastore-api/application/infrastructure"
	"context"
	"errors"

	domain "altastore-api/domain/entities"

	"github.com/jinzhu/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

//NewCustomerRepository customer db repository
func NewCustomerRepository(DB *gorm.DB) infrastructure.CustomerRepository {
	return &customerRepository{
		DB: DB,
	}
}

// Get User by ID
func (p *customerRepository) GetCustomerByID(ctx context.Context, accountID string) (domain.Customer, error) {
	customerData := domain.Customer{}

	if p.DB.Where("id = ?", accountID).First(&customerData).RecordNotFound() {
		return customerData, errors.New("Customer dengan ID " + accountID + " tidak dapat ditemukan")
	}
	return customerData, nil
}
