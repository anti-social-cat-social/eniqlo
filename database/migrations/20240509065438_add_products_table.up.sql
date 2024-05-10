CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE product_categories AS ENUM (
    'Clothing',
    'Accessories',
    'Footwear',
    'Beverages'
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    sku VARCHAR NOT NULL,
    category product_categories NOT NULL,
    image_url VARCHAR NOT NULL,
    notes VARCHAR NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    location VARCHAR NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_products_sku ON products(sku);

CREATE INDEX IF NOT EXISTS idx_products_price ON products(price);

CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);