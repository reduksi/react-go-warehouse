CREATE TABLE goods_issue_detail (
    TrxOutDPK SERIAL PRIMARY KEY,
    TrxOutIDF INT REFERENCES goods_issue_header(TrxOutPK),
    TrxOutDProductIdf INT NOT NULL,
    TrxOutDQtyDus INT NOT NULL,
    TrxOutDQtyPcs INT NOT NULL
);