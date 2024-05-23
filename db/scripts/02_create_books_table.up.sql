
CREATE TABLE IF NOT EXISTS books (
    id CHAR(36) PRIMARY KEY, 
    title VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

