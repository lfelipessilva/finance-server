CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(7) NOT NULL,
    icon VARCHAR(255) NOT NULL
);

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    value INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    original_name VARCHAR(255) NOT NULL,
    description TEXT,
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    card VARCHAR(255),
    bank VARCHAR(255),
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);