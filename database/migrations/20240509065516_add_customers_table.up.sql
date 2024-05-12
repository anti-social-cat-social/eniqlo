-- Active: 1715154768988@@127.0.0.1@5432@eniqlo_db@public
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone_number VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    customers (phone_number, name)
VALUES ('+6281122334455', 'John Doe'),
    (
        '+6285223344556',
        'Jane Smith'
    ),
    (
        '+6281223344557',
        'Michael Johnson'
    ),
    (
        '+6283223344558',
        'Emily Davis'
    ),
    (
        '+6284223344559',
        'David Wilson'
    ),
    (
        '+6285223344560',
        'Sarah Brown'
    ),
    (
        '+6286223344561',
        'Chris Taylor'
    ),
    (
        '+6287223344562',
        'Jessica Martinez'
    ),
    (
        '+6288223344563',
        'Ryan Thompson'
    ),
    (
        '+6289223344564',
        'Amanda Garcia'
    );