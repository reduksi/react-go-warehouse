CREATE TABLE master_supplier (
    SupplierPK SERIAL PRIMARY KEY,
    SupplierName VARCHAR(255) NOT NULL
);

CREATE TABLE master_customer (
    CustomerPK SERIAL PRIMARY KEY,
    CustomerName VARCHAR(255) NOT NULL
);

CREATE TABLE master_product (
    ProductPK SERIAL PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL
);

CREATE TABLE master_warehouse (
    WhsPK SERIAL PRIMARY KEY,
    WhsName VARCHAR(255) NOT NULL
);

CREATE TABLE goods_receipt_header (
    TrxInPK SERIAL PRIMARY KEY,
    TrxInNo VARCHAR(255) NOT NULL,
    WhsIdf INT REFERENCES master_warehouse(WhsPK),
    TrxInDate TIMESTAMP NOT NULL,
    TrxInSuppIdf INT REFERENCES master_supplier(SupplierPK),
    TrxInNotes TEXT
);

CREATE TABLE goods_receipt_detail (
    TrxInDPK SERIAL PRIMARY KEY,
    TrxInIDF INT REFERENCES goods_receipt_header(TrxInPK),
    TrxInDProductIdf INT REFERENCES master_product(ProductPK),
    TrxInDQtyDus INT NOT NULL,
    TrxInDQtyPcs INT NOT NULL
);

CREATE TABLE goods_issue_header (
    TrxOutPK SERIAL PRIMARY KEY,
    TrxOutNo VARCHAR(255) NOT NULL,
    WhsIdf INT REFERENCES master_warehouse(WhsPK),
    TrxOutDate TIMESTAMP NOT NULL,
    TrxOutSuppIdf INT REFERENCES master_supplier(SupplierPK),
    TrxOutNotes TEXT
);

CREATE TABLE goods_issue_detail (
    TrxOutDPK SERIAL PRIMARY KEY,
    TrxOutIDF INT REFERENCES goods_issue_header(TrxOutPK),
    TrxOutDProductIdf INT REFERENCES master_product(ProductPK),
    TrxOutDQtyDus INT NOT NULL,
    TrxOutDQtyPcs INT NOT NULL
);