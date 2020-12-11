package db

import (
	"altastore-api/application/infrastructure"
	domain "altastore-api/domain/entities"
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

//productRepository  db repository
type productRepository struct {
	DB *gorm.DB
}

//NewProductRepository construct
func NewProductRepository(DB *gorm.DB) infrastructure.ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (c *productRepository) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := c.DB.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (c *productRepository) GetCategoryByID(ctx context.Context, categoryID string) (domain.Category, error) {
	category := domain.Category{}
	if c.DB.First(&category, "id = ?", categoryID).RecordNotFound() {
		return category, errors.New("Kategory dengan id " + categoryID + " tidak ditemukan")
	}

	return category, nil
}

func (c *productRepository) GetListProductByCategory(ctx context.Context, pagination domain.PaginationRequest) []domain.Product {
	var products []domain.Product

	return products
}
