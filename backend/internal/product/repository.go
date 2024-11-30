package product

import (
	"database/sql"
	"log"
	"warehouse/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Master_Product, error)
	CreateProduct(product models.Master_Product) (models.Master_Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (repo *productRepository) GetAllProducts() ([]models.Master_Product, error) {
	rows, err := repo.db.Query("SELECT ProductPK, ProductName FROM master_product")
	if err != nil {
		log.Println("Error fetching products:", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Master_Product
	for rows.Next() {
		var product models.Master_Product
		if err := rows.Scan(&product.ProductPK, &product.ProductName); err != nil {
			log.Println("Error scanning product:", err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *productRepository) CreateProduct(product models.Master_Product) (models.Master_Product, error) {
	query := `INSERT INTO master_product (ProductName) 
              VALUES ($1) 
              RETURNING ProductPK, ProductName`

	var newProduct models.Master_Product
	err := repo.db.QueryRow(query, product.ProductName).Scan(&newProduct.ProductPK, &newProduct.ProductName)
	if err != nil {
		log.Println("Error inserting product:", err)
		return models.Master_Product{}, err
	}
	return newProduct, nil
}