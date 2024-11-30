package warehouse

import (
	"database/sql"
	"log"
	"warehouse/models"
)

type WarehouseRepository interface {
	GetAllWarehouses() ([]models.Master_Warehouse, error)
	CreateWarehouse(warehouse models.Master_Warehouse) (models.Master_Warehouse, error)
}

type warehouseRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepository{db: db}
}

func (repo *warehouseRepository) GetAllWarehouses() ([]models.Master_Warehouse, error) {
	rows, err := repo.db.Query("SELECT WhsPK, WhsName FROM master_warehouse")
	if err != nil {
		log.Println("Error fetching warehouses:", err)
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.Master_Warehouse
	for rows.Next() {
		var warehouse models.Master_Warehouse
		if err := rows.Scan(&warehouse.WhsPK, &warehouse.WhsName); err != nil {
			log.Println("Error scanning warehouse:", err)
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (repo *warehouseRepository) CreateWarehouse(warehouse models.Master_Warehouse) (models.Master_Warehouse, error) {
	query := `INSERT INTO master_warehouse (WhsName) 
              VALUES ($1) 
              RETURNING WhsPK, WhsName`

	var newWarehouse models.Master_Warehouse
	err := repo.db.QueryRow(query, warehouse.WhsName).Scan(&newWarehouse.WhsPK, &newWarehouse.WhsName)
	if err != nil {
		log.Println("Error inserting warehouse:", err)
		return models.Master_Warehouse{}, err
	}
	return newWarehouse, nil
}