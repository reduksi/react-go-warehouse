package issueTransaction

import (
	"database/sql"
	"log"
	"warehouse/models"
)

type TransactionRepository interface {
	GetAllGoodsIssueHeaders() ([]models.GoodsIssueHeader, error)
	CreateGoodsIssueHeader(header models.GoodsIssueHeader, details []models.GoodsIssueDetail) (models.GoodsIssueHeader, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) CreateGoodsIssueHeader(header models.GoodsIssueHeader, details []models.GoodsIssueDetail) (models.GoodsIssueHeader, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return models.GoodsIssueHeader{}, err
	}
	// Insert goods issue header
	var newHeader models.GoodsIssueHeader
	err = tx.QueryRow(
		`INSERT INTO goods_issue_header (TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes) 
		 VALUES ($1, $2, $3, $4, $5) 
		 RETURNING TrxOutPK, TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes`,
		header.TrxOutNo, header.WhsIdf, header.TrxOutDate, header.TrxOutSuppIdf, header.TrxOutNotes,
	).Scan(&newHeader.TrxOutPK, &newHeader.TrxOutNo, &newHeader.WhsIdf, &newHeader.TrxOutDate, &newHeader.TrxOutSuppIdf, &newHeader.TrxOutNotes)
	if err != nil {
		tx.Rollback()
		log.Println("Error inserting goods issue header:", err)
		return models.GoodsIssueHeader{}, err
	}
	// Insert goods issue detail
	for _, detail := range details {
		_, err := tx.Exec(
			`INSERT INTO goods_issue_detail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs) 
			 VALUES ($1, $2, $3, $4)`,
			newHeader.TrxOutPK, detail.TrxOutDProductIdf, detail.TrxOutDQtyDus, detail.TrxOutDQtyPcs,
		)
		if err != nil {
			tx.Rollback()
			log.Println("Error inserting goods issue detail:", err)
			return models.GoodsIssueHeader{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return models.GoodsIssueHeader{}, err
	}
	return newHeader, nil
}


func (repo *transactionRepository) GetAllGoodsIssueHeaders() ([]models.GoodsIssueHeader, error) {
	headers, err := repo.getGoodsIssueHeaders()
	if err != nil {
		return nil, err
	}

	for i, header := range headers {
		details, err := repo.getGoodsIssueDetails(header.TrxOutPK)
		if err != nil {
			return nil, err
		}
		headers[i].Details = details
	}

	return headers, nil
}

func (repo *transactionRepository) getGoodsIssueHeaders() ([]models.GoodsIssueHeader, error) {
	rows, err := repo.db.Query("SELECT TrxOutPK, TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes FROM goods_issue_header")
	if err != nil {
		log.Println("Error fetching goods issue headers:", err)
		return nil, err
	}
	defer rows.Close()

	var headers []models.GoodsIssueHeader
	for rows.Next() {
		var header models.GoodsIssueHeader
		if err := rows.Scan(&header.TrxOutPK, &header.TrxOutNo, &header.WhsIdf, &header.TrxOutDate, &header.TrxOutSuppIdf, &header.TrxOutNotes); err != nil {
			log.Println("Error scanning goods issue header:", err)
			return nil, err
		}
		headers = append(headers, header)
	}
	return headers, nil
}

func (repo *transactionRepository) getGoodsIssueDetails(headerID int) ([]models.GoodsIssueDetail, error) {
	rows, err := repo.db.Query("SELECT TrxOutDPK, TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs FROM goods_issue_detail WHERE TrxOutIDF = $1", headerID)
	if err != nil {
		log.Println("Error fetching goods issue details:", err)
		return nil, err
	}
	defer rows.Close()

	var details []models.GoodsIssueDetail
	for rows.Next() {
		var detail models.GoodsIssueDetail
		if err := rows.Scan(&detail.TrxOutDPK, &detail.TrxOutIDF, &detail.TrxOutDProductIdf, &detail.TrxOutDQtyDus, &detail.TrxOutDQtyPcs); err != nil {
			log.Println("Error scanning goods issue detail:", err)
			return nil, err
		}
		details = append(details, detail)
	}
	return details, nil
}