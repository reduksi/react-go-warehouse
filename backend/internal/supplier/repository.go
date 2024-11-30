package supplier

import (
	"database/sql"
	"log"
	"warehouse/models"
)

type CustomerRepository interface {
	GetAllSuppliers() ([]models.Master_Supplier, error)
	CreateSupplier(customer models.Master_Supplier) (models.Master_Supplier, error)
}

type Repository struct {
    DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{DB: db}
}

func (r *Repository) GetAllSuppliers() ([]models.Master_Supplier, error) {
    rows, err := r.DB.Query("SELECT SupplierPK, SupplierName FROM master_supplier")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var suppliers []models.Master_Supplier
    for rows.Next() {
        var s models.Master_Supplier
        if err := rows.Scan(&s.SupplierPK, &s.SupplierName); err != nil {
            return nil, err
        }
        suppliers = append(suppliers, s)
    }
    return suppliers, nil
}

func (r *Repository) CreateSupplier(supplier models.Master_Supplier) (models.Master_Supplier, error) {
    query := `INSERT INTO master_supplier (SupplierName) 
              VALUES ($1) 
              RETURNING SupplierPK, SupplierName`

    var newSupplier models.Master_Supplier
    err := r.DB.QueryRow(query, supplier.SupplierName).Scan(&newSupplier.SupplierPK, &newSupplier.SupplierName)
    if err != nil {
        log.Println("Error inserting supplier:", err)
        return models.Master_Supplier{}, err
    }
    return newSupplier, nil
}