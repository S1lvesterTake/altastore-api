//+build wireinject

package main

//TODO will fix it
import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"altastore-api/application/use_case/authentication/login"

	repo "altastore-api/infrastructure/persistence/repository/db"
)

//LoginHandler wire
func LoginHandler(db *gorm.DB) login.LoginHandler {
	wire.Build(repo.NewLoginRepository, repo.NewAccountRepository, repo.NewCustomerRepository, login.NewLoginHandler)
	return login.LoginHandler{}
}
