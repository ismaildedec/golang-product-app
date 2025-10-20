package service

import (
	"errors"

	"golang-product-app.git/domain"
	"golang-product-app.git/persistence"
	"golang-product-app.git/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	GetById(productId int64) (domain.Product, error)
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) Add(productCreate model.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}
	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProduct()
}
func (productService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(storeName)
}
func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount can not be greater than 70")
	}
	return nil
}
