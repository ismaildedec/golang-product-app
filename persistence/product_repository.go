package persistence

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"golang-product-app.git/domain"
)

type IProductRepository interface {
	GetAllProduct() []domain.Product
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}

}

func (productRepository *ProductRepository) GetAllProduct() []domain.Product {

	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "Select * from products ")

	if err != nil {
		log.Error("Tüm ürünler alınırken hata oluştu %v", err)
		return []domain.Product{}
	}

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {

		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}
