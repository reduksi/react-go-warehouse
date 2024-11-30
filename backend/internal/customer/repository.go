package customer

import (
	"database/sql"
	"log"
	"warehouse/models"
)

type CustomerRepository interface {
	CreateCustomer(customer models.Master_Customer) (models.Master_Customer, error)
	GetAllCustomers() ([]models.Master_Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (repo *customerRepository) GetAllCustomers() ([]models.Master_Customer, error) {
	rows, err := repo.db.Query("SELECT CustomerPK, CustomerName FROM master_customer")
	if err != nil {
		log.Println("Error fetching customers:", err)
		return nil, err
	}
	defer rows.Close()

	var customers []models.Master_Customer
	for rows.Next() {
		var customer models.Master_Customer
		if err := rows.Scan(&customer.CustomerPK, &customer.CustomerName); err != nil {
			log.Println("Error scanning customer:", err)
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repo *customerRepository) CreateCustomer(customer models.Master_Customer) (models.Master_Customer, error) {
	query := `INSERT INTO master_customer (CustomerName) 
              VALUES ($1) 
              RETURNING CustomerPK, CustomerName`

	var newCustomer models.Master_Customer
	err := repo.db.QueryRow(query, customer.CustomerName).Scan(&newCustomer.CustomerPK, &newCustomer.CustomerName)
	if err != nil {
		log.Println("Error inserting customer:", err)
		return models.Master_Customer{}, err
	}
	return newCustomer, nil
}