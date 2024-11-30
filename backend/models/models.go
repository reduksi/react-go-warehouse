package models

import "time"

type Master_Supplier struct {
    SupplierPK   int    `json:"supplierPK" db:"SupplierPK"`
    SupplierName string `json:"supplierName" db:"SupplierName"`
}

type Master_Customer struct {
    CustomerPK   int    `json:"customerPK" db:"CustomerPK"`
    CustomerName string `json:"customerName" db:"CustomerName"`
}

type Master_Product struct {
    ProductPK   int    `json:"productPK" db:"ProductPK"`
    ProductName string `json:"productName" db:"ProductName"`
}

type Master_Warehouse struct {
    WhsPK   int    `json:"whsPK" db:"WhsPK"`
    WhsName string `json:"whsName" db:"WhsName"`
}

type GoodsReceiptHeader struct {
    TrxInPK      int       `json:"trxInPK" db:"TrxInPK"`
    TrxInNo      string    `json:"trxInNo" db:"TrxInNo"`
    WhsIdf       int       `json:"whsIdf" db:"WhsIdf"`
    TrxInDate    time.Time `json:"trxInDate" db:"TrxInDate"`
    TrxInSuppIdf int       `json:"trxInSuppIdf" db:"TrxInSuppIdf"`
    TrxInNotes   string    `json:"trxInNotes" db:"TrxInNotes"`
    Details      []GoodsReceiptDetail     `json:"details"`
}

type GoodsReceiptDetail struct {
    TrxInDPK        int `json:"trxInDPK" db:"TrxInDPK"`
    TrxInIDF         int `json:"trxInIDF" db:"TrxInIDF"`  // Foreign key to GoodsReceiptHeader
    TrxInDProductIdf int `json:"trxInDProductIdf" db:"TrxInDProductIdf"`
    TrxInDQtyDus     int `json:"trxInDQtyDus" db:"TrxInDQtyDus"`
    TrxInDQtyPcs     int `json:"trxInDQtyPcs" db:"TrxInDQtyPcs"`
}

type GoodsIssueHeader struct {
    TrxOutPK      int       `json:"trxOutPK" db:"TrxOutPK"`
    TrxOutNo      string    `json:"trxOutNo" db:"TrxOutNo"`
    WhsIdf        int       `json:"whsIdf" db:"WhsIdf"`
    TrxOutDate    time.Time `json:"trxOutDate" db:"TrxOutDate"`
    TrxOutSuppIdf int       `json:"trxOutSuppIdf" db:"TrxOutSuppIdf"`
    TrxOutNotes   string    `json:"trxOutNotes" db:"TrxOutNotes"`
    Details      []GoodsIssueDetail     `json:"details"`
}

type GoodsIssueDetail struct {
    TrxOutDPK        int `json:"trxOutDPK" db:"TrxOutDPK"`
    TrxOutIDF         int `json:"trxOutIDF" db:"TrxOutIDF"` // Foreign key to GoodsIssueHeader
    TrxOutDProductIdf int `json:"trxOutDProductIdf" db:"TrxOutDProductIdf"`
    TrxOutDQtyDus     int `json:"trxOutDQtyDus" db:"TrxOutDQtyDus"`
    TrxOutDQtyPcs     int `json:"trxOutDQtyPcs" db:"TrxOutDQtyPcs"`
}

type StockReport struct {
    ProductPK    int     `json:"productPK" db:"ProductPK"`
    ProductName  string  `json:"productName" db:"ProductName"`
    StockInDus   int     `json:"stockInDus" db:"StockInDus"`
    StockInPcs   int     `json:"stockInPcs" db:"StockInPcs"`
}
