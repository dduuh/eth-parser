CREATE TABLE addresses (
    id SERIAL PRIMARY KEY NOT NULL,
    address VARCHAR(255) NOT NULL,
    balance FLOAT NOT NULL,
    private_key VARCHAR(255) NOT NULL,
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY NOT NULL,
    from VARCHAR(255) NOT NULL,
    to VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
);