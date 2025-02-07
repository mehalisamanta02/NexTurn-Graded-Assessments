package repositories

import (
	"database/sql"
	"ecommerce-inventory/models"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) AddProduct(product *models.Product) error {
	_, err := repo.db.Exec(`INSERT INTO products (name, description, price, stock, category_id) 
		VALUES (?, ?, ?, ?, ?)`, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	return err
}

func (repo *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	row := repo.db.QueryRow(`SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?`, id)
	product := &models.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	_, err := repo.db.Exec(`UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ? 
		WHERE id = ?`, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ID)
	return err
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	_, err := repo.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}

func (repo *ProductRepository) GetAllProducts(page, limit int) ([]models.Product, error) {
	rows, err := repo.db.Query(`SELECT id, name, description, price, stock, category_id FROM products LIMIT ? OFFSET ?`,
		limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
