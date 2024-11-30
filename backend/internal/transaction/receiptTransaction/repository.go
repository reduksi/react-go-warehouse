package receiptTransaction

import (
	"database/sql"
	"log"
	"warehouse/models"
)

type TransactionRepository interface {
	GetAllGoodsReceiptHeaders() ([]models.GoodsReceiptHeader, error)
	GetStockReport() ([]models.StockReport, error)
	CreateGoodsReceipt(header models.GoodsReceiptHeader, details []models.GoodsReceiptDetail) (models.GoodsReceiptHeader, error)

}

type transactionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) CreateGoodsReceipt(header models.GoodsReceiptHeader, details []models.GoodsReceiptDetail) (models.GoodsReceiptHeader, error) {
    tx, err := repo.db.Begin()
    if err != nil {
		log.Println("Error starting transaction:", err)
        return models.GoodsReceiptHeader{}, err
    }

    // Insert goods_receipt_header 
    var newHeader models.GoodsReceiptHeader
    err = tx.QueryRow(
        `INSERT INTO goods_receipt_header (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes) 
		 VALUES ($1, $2, $3, $4, $5) 
		 RETURNING TrxInPK, TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes`,
        header.TrxInNo, header.WhsIdf, header.TrxInDate, header.TrxInSuppIdf, header.TrxInNotes,
    ).Scan(&newHeader.TrxInPK, &newHeader.TrxInNo, &newHeader.WhsIdf, &newHeader.TrxInDate, &newHeader.TrxInSuppIdf, &newHeader.TrxInNotes)
    if err != nil {
        return models.GoodsReceiptHeader{}, err
    }
    // Insert goods_receipt_detail
    for _, detail := range details {
        _, err := tx.Exec(
            `INSERT INTO goods_receipt_detail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs) 
			 VALUES ($1, $2, $3, $4)`,
            newHeader.TrxInPK, detail.TrxInDProductIdf, detail.TrxInDQtyDus, detail.TrxInDQtyPcs,
        )
        if err != nil {
			tx.Rollback()
			log.Println("Error inserting goods issue detail:", err)
			return models.GoodsReceiptHeader{}, err
		}
    }

    if err := tx.Commit(); err != nil {
        return models.GoodsReceiptHeader{}, err
    }

    return newHeader, nil
}

func (repo *transactionRepository) GetAllGoodsReceiptHeaders() ([]models.GoodsReceiptHeader, error) {
	headers, err := repo.getGoodsReceiptHeaders()
	if err != nil {
		return nil, err
	}

	for i, header := range headers {
		details, err := repo.getGoodsReceiptDetails(header.TrxInPK)
		if err != nil {
			return nil, err
		}
		headers[i].Details = details
	}

	return headers, nil
}

func (repo *transactionRepository) getGoodsReceiptHeaders() ([]models.GoodsReceiptHeader, error) {
	rows, err := repo.db.Query("SELECT TrxInPK, TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes FROM goods_receipt_header")
	if err != nil {
		log.Println("Error fetching goods receipt headers:", err)
		return nil, err
	}
	defer rows.Close()

	var headers []models.GoodsReceiptHeader
	for rows.Next() {
		var header models.GoodsReceiptHeader
		if err := rows.Scan(&header.TrxInPK, &header.TrxInNo, &header.WhsIdf, &header.TrxInDate, &header.TrxInSuppIdf, &header.TrxInNotes); err != nil {
			log.Println("Error scanning goods receipt header:", err)
			return nil, err
		}
		headers = append(headers, header)
	}
	return headers, nil
}

func (repo *transactionRepository) getGoodsReceiptDetails(headerID int) ([]models.GoodsReceiptDetail, error) {
	rows, err := repo.db.Query("SELECT TrxInDPK, TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs FROM goods_receipt_detail WHERE TrxInIDF = $1", headerID)
	if err != nil {
		log.Println("Error fetching goods receipt details:", err)
		return nil, err
	}
	defer rows.Close()

	var details []models.GoodsReceiptDetail
	for rows.Next() {
		var detail models.GoodsReceiptDetail
		if err := rows.Scan(&detail.TrxInDPK, &detail.TrxInIDF, &detail.TrxInDProductIdf, &detail.TrxInDQtyDus, &detail.TrxInDQtyPcs); err != nil {
			log.Println("Error scanning goods receipt detail:", err)
			return nil, err
		}
		details = append(details, detail)
	}
	return details, nil
}

func (repo *transactionRepository) GetStockReport() ([]models.StockReport, error) {
    query := `
        SELECT 
			p.ProductPK, 
			p.ProductName, 
			COALESCE(SUM(grd.TrxInDQtyDus), 0) - COALESCE(SUM(gid.TrxOutDQtyDus), 0) AS StockInDus,
			COALESCE(SUM(grd.TrxInDQtyPcs), 0) - COALESCE(SUM(gid.TrxOutDQtyPcs), 0) AS StockInPcs
		FROM 
			master_product p
		LEFT JOIN 
			goods_receipt_detail grd ON grd.TrxInDProductIdf = p.ProductPK
		LEFT JOIN 
			goods_issue_detail gid ON gid.TrxOutDProductIdf = p.ProductPK
		GROUP BY 
			p.ProductPK, p.ProductName
		ORDER BY 
			p.ProductPK;
    `

    rows, err := repo.db.Query(query)
    if err != nil {
        log.Println("Error querying stock report:", err)
        return nil, err
    }
    defer rows.Close()

    var report []models.StockReport
    for rows.Next() {
        var r models.StockReport
        if err := rows.Scan(&r.ProductPK, &r.ProductName, &r.StockInDus, &r.StockInPcs); err != nil {
            log.Println("Error scanning row:", err)
            return nil, err
        }
        report = append(report, r)
    }

    if err := rows.Err(); err != nil {
        log.Println("Error in rows:", err)
        return nil, err
    }

    return report, nil
}