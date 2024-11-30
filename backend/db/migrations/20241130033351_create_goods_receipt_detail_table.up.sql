CREATE TABLE goods_receipt_detail (
    TrxInDPK SERIAL PRIMARY KEY,
    TrxInIDF INT REFERENCES goods_receipt_header(TrxInPK),
    TrxInDProductIdf INT NOT NULL,
    TrxInDQtyDus INT NOT NULL,
    TrxInDQtyPcs INT NOT NULL
);