package db

import (
	"altastore-api/application/infrastructure"
	domain "altastore-api/domain/entities"
	"context"
	"errors"
	"strconv"

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

func (c *productRepository) GetListProduct(ctx context.Context, pagination domain.FilterRequest) ([]domain.Product, int) {
	var products []domain.Product
	var count int
	page, _ := strconv.Atoi(pagination.Page)
	limit, _ := strconv.Atoi(pagination.Limit)

	offset := (page - 1) * limit

	tx := c.DB.Model(&products)

	if pagination.Filter != "" {
		tx = tx.Where("category_id = ?", pagination.Filter)
	}
	tx.Count(&count)

	if pagination.Page != "" || pagination.Limit != "0" {
		tx.Offset(offset).Limit(pagination.Limit).Find(&products)
	}
	return products, count
}
