CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    value NUMERIC(10,2) NOT NULL
);

CREATE INDEX idx_expenses_timestamp ON expenses(timestamp);
CREATE INDEX idx_expenses_category ON expenses(category);