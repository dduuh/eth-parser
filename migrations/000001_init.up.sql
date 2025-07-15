CREATE TABLE addresses (
    id SERIAL PRIMARY KEY NOT NULL,
    address VARCHAR(255) NOT NULL,
    balance FLOAT NOT NULL,
    private_key VARCHAR(255) NOT NULL,
);