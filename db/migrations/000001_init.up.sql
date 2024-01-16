CREATE TABLE IF NOT EXISTS warehouses (
    warehouse_id SERIAL PRIMARY KEY,
    warehouse_name TEXT NOT NULL,
    warehouse_available BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    product_code UUID NOT NULL PRIMARY KEY,
    product_name TEXT NOT NULL,
    product_size TEXT,
    product_value INTEGER NOT NULL,
    warehouse_id INTEGER NOT NULL REFERENCES warehouses (warehouse_id),
    product_reserved_value INTEGER NOT NULL
);