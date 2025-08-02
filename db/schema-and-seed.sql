-- Create schema (if you need to explicitly create the database, do it outside this script, as CREATE DATABASE cannot run within a transaction block)
-- CREATE DATABASE acmedb;
-- Switch to the "acmedb" database if executing from a SQL client that supports meta-commands
-- \c acmedb;
-- Create Customers Table
CREATE TABLE IF NOT EXISTS
    customers (
        customer_id serial PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Create Orders Table
CREATE TABLE IF NOT EXISTS
    orders (
        order_id serial PRIMARY KEY,
        customer_id INTEGER NOT NULL,
        order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        total_amount DECIMAL(10, 2) NOT NULL,
        CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers (customer_id) ON DELETE CASCADE
    );

-- Create Inventory Table
CREATE TABLE IF NOT EXISTS
    inventory (
        product_id serial PRIMARY KEY,
        NAME VARCHAR(255) NOT NULL,
        quantity INTEGER NOT NULL,
        price DECIMAL(10, 2) NOT NULL
    );

-- Insert some data into Customers
INSERT INTO
    customers (NAME, email)
VALUES
    ('John Doe', 'john.doe@example.com'),
    ('Jane Doe', 'jane.doe@example.com');

-- Insert some data into Inventory
INSERT INTO
    inventory (NAME, quantity, price)
VALUES
    ('Widget A', 100, 9.99),
    ('Widget B', 200, 19.99),
    ('Widget C', 150, 29.99);

-- Insert some orders linking to customers and assuming some arbitrary totals
INSERT INTO
    orders (customer_id, total_amount)
VALUES
    (1, 199.95),
    (2, 299.95);

-- Add comment on Database
COMMENT ON DATABASE acmedb IS 'Dummy comment on acmedb';