package services

import (
	"ecommerce-inventory/models"
	"ecommerce-inventory/repositories"
	"errors"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// AddProduct adds a new product to the inventory.
func (service *ProductService) AddProduct(product *models.Product) error {
	// Validate product data
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.New("invalid product data")
	}

	// Insert product into the database
	return service.repo.AddProduct(product)
}

// GetProductByID retrieves a product by its ID.
func (service *ProductService) GetProductByID(id int) (*models.Product, error) {
	product, err := service.repo.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// UpdateProduct updates a product's details.
func (service *ProductService) UpdateProduct(product *models.Product) error {
	// Validate product data
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.New("invalid product data")
	}

	// Update product in the database
	return service.repo.UpdateProduct(product)
}

// DeleteProduct deletes a product from the inventory.
func (service *ProductService) DeleteProduct(id int) error {
	return service.repo.DeleteProduct(id)
}

// GetAllProducts retrieves all products with pagination.
func (service *ProductService) GetAllProducts(page, limit int) ([]models.Product, error) {
	return service.repo.GetAllProducts(page, limit)
}
