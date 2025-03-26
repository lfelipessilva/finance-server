CREATE TABLE installments (
    id SERIAL PRIMARY KEY,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    value DECIMAL(10, 2) NOT NULL CHECK (value > 0),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE expenses
ADD COLUMN installment_id INT NULL,
ADD COLUMN invoice_id INT NULL,
ADD CONSTRAINT fk_installment FOREIGN KEY (installment_id) REFERENCES installments (id) ON DELETE CASCADE,
ADD CONSTRAINT fk_invoice FOREIGN KEY (invoice_id) REFERENCES invoices (id) ON DELETE SET NULL;