package db

import (
	"altastore-api/application/infrastructure"
	"context"
	"errors"

	domain "altastore-api/domain/entities"

	"github.com/jinzhu/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

//NewAccountRepository account db repository
func NewAccountRepository(DB *gorm.DB) infrastructure.AccountRepository {
	return &accountRepository{
		DB: DB,
	}
}

// Get User by email
func (p *accountRepository) GetAccountByEmail(ctx context.Context, email string) (domain.Account, error) {
	accountData := domain.Account{}

	if p.DB.Where("email = ?", email).First(&accountData).RecordNotFound() {
		return accountData, errors.New("Email salah")
	}
	return accountData, nil
}

// Get User by ID
func (p *accountRepository) GetAccountByID(ctx context.Context, accountID string) (domain.Account, error) {
	account := domain.Account{}

	if p.DB.Preload("AccessToken").First(&account, accountID).RecordNotFound() {
		return account, errors.New("Account dengan ID " + accountID + " tidak dapat ditemukan")
	}
	return account, nil
}
