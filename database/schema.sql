
CREATE DATABASE IF NOT EXISTS komputer_manufacture_db;
USE komputer_manufacture_db;


CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE user_profiles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE components (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50) NOT NULL,
    stock INT NOT NULL CHECK (stock >= 0),
    price DECIMAL(12,2) NOT NULL CHECK (price >= 0)
);

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(12,2) NOT NULL CHECK (price >= 0),
    stock INT NOT NULL CHECK (stock >= 0)
);

CREATE TABLE product_components (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    component_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    UNIQUE (product_id, component_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (component_id) REFERENCES components(id) ON DELETE CASCADE
);

CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT
);

CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    order_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(30) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(12,2) NOT NULL CHECK (unit_price >= 0),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE assemblies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_item_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    start_date DATETIME NULL,
    end_date DATETIME NULL,
    status VARCHAR(30) NOT NULL,
    CHECK (end_date IS NULL OR end_date >= start_date),
    FOREIGN KEY (order_item_id) REFERENCES order_items(id)
);

CREATE TABLE assembly_components (
    id INT AUTO_INCREMENT PRIMARY KEY,
    assembly_id INT NOT NULL,
    component_id INT NOT NULL,
    quantity_used INT NOT NULL CHECK (quantity_used > 0),
    UNIQUE (assembly_id, component_id),
    FOREIGN KEY (assembly_id) REFERENCES assemblies(id) ON DELETE CASCADE,
    FOREIGN KEY (component_id) REFERENCES components(id)
);

ALTER TABLE assemblies
    ADD CONSTRAINT uq_assemblies_order_item
        UNIQUE (order_item_id);