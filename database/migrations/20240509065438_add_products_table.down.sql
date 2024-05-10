DROP TABLE IF EXISTS products;

DROP TYPE IF EXISTS product_categories;

DROP INDEX IF EXISTS idx_products_sku CASCADE;

DROP INDEX IF EXISTS idx_products_price CASCADE;

DROP INDEX IF EXISTS idx_products_created_at CASCADE;