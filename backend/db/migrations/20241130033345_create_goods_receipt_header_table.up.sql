CREATE TABLE goods_receipt_header (
    TrxInPK SERIAL PRIMARY KEY,
    TrxInNo VARCHAR(255) NOT NULL,
    WhsIdf INT NOT NULL,
    TrxInDate DATE NOT NULL,
    TrxInSuppIdf INT NOT NULL,
    TrxInNotes VARCHAR(255)
);