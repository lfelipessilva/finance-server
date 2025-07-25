CREATE TABLE installments (
    id SERIAL PRIMARY KEY,
    value INT NOT NULL,
    quantity INT NOT NULL,
    timestamp_start TIMESTAMP NOT NULL DEFAULT NOW(),
    timestamp_end TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE expenses
ADD COLUMN installment_id INTEGER;

ALTER TABLE expenses
ADD CONSTRAINT fk_installment
FOREIGN KEY (installment_id)
REFERENCES installments (id)
ON DELETE CASCADE;
