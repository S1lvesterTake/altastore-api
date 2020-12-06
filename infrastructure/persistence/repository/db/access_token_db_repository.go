package db

import (
	"altastore-api/application/infrastructure"
	"context"
	"time"

	domain "altastore-api/domain/entities"

	"github.com/jinzhu/gorm"
)

type loginRepository struct {
	DB *gorm.DB
}

//NewLoginRepository DB repository
func NewLoginRepository(DB *gorm.DB) infrastructure.LoginRepository {
	return &loginRepository{
		DB: DB,
	}
}

func (c *loginRepository) Login(ctx context.Context, accessToken domain.AccessToken, accountID uint64) (domain.AccessToken, error) {
	accToken := domain.AccessToken{}
	account := domain.Account{}
	accountData := domain.Account{}
	account.LastLogin = time.Now()

	errUpdateAccount := c.DB.Model(&accountData).Where("id = ?", accountID).Update(&account).Error
	if errUpdateAccount != nil {
		return accessToken, errUpdateAccount
	}

	checkToken := c.DB.Where("account_id = ?", accountID).First(&accToken).RecordNotFound()
	if checkToken {
		err := c.DB.Create(&accessToken).Error
		if err != nil {
			return accessToken, err
		}
		return accessToken, nil

	} else {
		if err := c.DB.Model(&accToken).Where("account_id = ?", accountID).Update(&accessToken).Error; err != nil {
			return accessToken, err
		}
		return accessToken, nil
	}

}
