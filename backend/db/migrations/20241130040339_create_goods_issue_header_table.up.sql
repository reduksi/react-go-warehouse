CREATE TABLE goods_issue_header (
    TrxOutPK SERIAL PRIMARY KEY,
    TrxOutNo VARCHAR(255) NOT NULL,
    WhsIdf INT NOT NULL,
    TrxOutDate DATE NOT NULL,
    TrxOutSuppIdf INT NOT NULL,
    TrxOutNotes VARCHAR(255)
);