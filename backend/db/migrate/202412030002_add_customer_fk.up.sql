ALTER TABLE goods_issue_header
ADD COLUMN TrxOutCustIdf INT,
ADD CONSTRAINT fk_customer FOREIGN KEY (TrxOutCustIdf) REFERENCES master_customer(CustomerPK);
