-- DUMMY DATA 

-- 1. Users & Profiles (Password is 'admin123' hashed with bcrypt)
INSERT INTO users (id, email, password_hash) VALUES 
(1, 'hacktiv8@mail.com', '$2a$10$y58n1/t98.kZJpT5XJ.W..rX.V7wJ8O0R.J3L/v1vP4G6.H6f2W'),
(2, 'dino@mail.com', '$2a$10$y58n1/t98.kZJpT5XJ.W..rX.V7wJ8O0R.J3L/v1vP4G6.H6f2W'),
(3, 'refpi@mail.com', '$2a$10$y58n1/t98.kZJpT5XJ.W..rX.V7wJ8O0R.J3L/v1vP4G6.H6f2W');

INSERT INTO user_profiles (user_id, full_name, phone) VALUES 
(1, 'Refpi Irmansah', '081234567890'),
(2, 'Budi Santoso', '081298765432'),
(3, 'Siti Aminah', '081345678912');

-- 2. Components
INSERT INTO components (id, name, category, stock, price) VALUES 
(1, 'Intel Core i7-13700K', 'CPU', 50, 450.00),
(2, 'NVIDIA RTX 4070 Ti', 'GPU', 8, 800.00), -- Stock low for report
(3, 'Corsair Vengeance 32GB RAM', 'RAM', 100, 120.00),
(4, 'Samsung 980 PRO 1TB NVMe', 'Storage', 150, 95.00),
(5, 'NZXT H510i Case', 'Case', 5, 105.00); -- Stock low for report

-- 3. Products (Final Rigs)
INSERT INTO products (id, name, description, price, stock) VALUES 
(1, 'Beast Gaming Rig V1', 'High-end gaming computer', 1800.00, 20),
(2, 'Office Workstation Pro', 'Reliable multi-tasking PC', 900.00, 35);

-- 4. Product Components (M:N Relationship)
INSERT INTO product_components (product_id, component_id, quantity) VALUES 
(1, 1, 1), (1, 2, 1), (1, 3, 1), (1, 4, 1), (1, 5, 1); -- Beast Rig Needs all

-- 5. Customers & Orders
INSERT INTO customers (id, name, email, phone, address) VALUES 
(1, 'Alex Wijaya', 'alex@gmail.com', '08111222333', 'Jl. Sudirman No 10, Jakarta'),
(2, 'Citra Lestari', 'citra@yahoo.com', '08222333444', 'BSD City, Tangerang');

INSERT INTO orders (id, customer_id, status) VALUES 
(1, 1, 'completed'),
(2, 2, 'pending_assembly'),
(3, 1, 'completed');

INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES 
(1, 1, 1, 1800.00),
(2, 2, 5, 900.00),
(3, 1, 2, 1800.00);
