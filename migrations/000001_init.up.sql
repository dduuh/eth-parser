CREATE TABLE addresses (
    id SERIAL UNIQUE NOT NULL,
    address VARCHAR(255) NOT NULL,
    balance FLOAT NOT NULL,
    private_key VARCHAR(255) NOT NULL,
);