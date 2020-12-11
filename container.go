//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"altastore-api/application/use_case/authentication/login"
	"altastore-api/application/use_case/product/create_product"

	repo "altastore-api/infrastructure/persistence/repository/db"
	request "altastore-api/infrastructure/transport/http"
)

//LoginHandler wire
func LoginHandler(db *gorm.DB) login.LoginHandler {
	wire.Build(repo.NewLoginRepository, repo.NewAccountRepository, repo.NewCustomerRepository, login.NewLoginHandler)
	return login.LoginHandler{}
}

func CreateProductHandler(db *gorm.DB) create_product.CreateProductHandler {
	wire.Build(request.NewRequest, repo.NewProductRepository, create_product.NewCreateProductHandler)
	return create_product.CreateProductHandler{}
}
