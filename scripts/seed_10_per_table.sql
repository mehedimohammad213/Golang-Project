-- Seed: 10 rows per table (matches live DB schema; run after truncate_all.sql)
-- Run: psql -h localhost -U postgres -d car_db -f scripts/seed_10_per_table.sql

-- 1. users (10)
INSERT INTO users (name, username, email, password_hash, is_active) VALUES
('Alice Smith', 'alice1', 'alice1@example.com', '$2a$10$dummyhash1', true),
('Bob Jones', 'bob2', 'bob2@example.com', '$2a$10$dummyhash2', true),
('Carol White', 'carol3', 'carol3@example.com', '$2a$10$dummyhash3', true),
('David Brown', 'david4', 'david4@example.com', '$2a$10$dummyhash4', true),
('Eve Davis', 'eve5', 'eve5@example.com', '$2a$10$dummyhash5', true),
('Frank Miller', 'frank6', 'frank6@example.com', '$2a$10$dummyhash6', true),
('Grace Lee', 'grace7', 'grace7@example.com', '$2a$10$dummyhash7', true),
('Henry Wilson', 'henry8', 'henry8@example.com', '$2a$10$dummyhash8', true),
('Ivy Taylor', 'ivy9', 'ivy9@example.com', '$2a$10$dummyhash9', true),
('Jack Anderson', 'jack10', 'jack10@example.com', '$2a$10$dummyhash10', true);

-- 2. roles (10)
INSERT INTO roles (name, slug, description) VALUES
('Admin', 'admin', 'Full system access'),
('Manager', 'manager', 'Manage cars and orders'),
('Sales', 'sales', 'Sales and customers'),
('Viewer', 'viewer', 'Read-only access'),
('Finance', 'finance', 'Payments and LCs'),
('Warehouse', 'warehouse', 'Stock and inventory'),
('Support', 'support', 'Customer support'),
('Auditor', 'auditor', 'Audit and reports'),
('Driver', 'driver', 'Delivery and logistics'),
('Guest', 'guest', 'Limited guest access');

-- 3. permissions (10)
INSERT INTO permissions (name, slug, module) VALUES
('users.create', 'users-create', 'users'),
('users.read', 'users-read', 'users'),
('users.update', 'users-update', 'users'),
('users.delete', 'users-delete', 'users'),
('cars.create', 'cars-create', 'cars'),
('cars.read', 'cars-read', 'cars'),
('cars.update', 'cars-update', 'cars'),
('orders.create', 'orders-create', 'orders'),
('orders.read', 'orders-read', 'orders'),
('reports.read', 'reports-read', 'reports');

-- 4. role_user (10): user 1-10 -> role 1-10
INSERT INTO role_user (user_id, role_id, assigned_by) VALUES
(1, 1, 1), (2, 2, 1), (3, 3, 1), (4, 4, 1), (5, 5, 1),
(6, 6, 1), (7, 7, 1), (8, 8, 1), (9, 9, 1), (10, 10, 1);

-- 5. permission_role (10): permission 1-10 -> role 1-10
INSERT INTO permission_role (permission_id, role_id) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 5),
(6, 6), (7, 7), (8, 8), (9, 9), (10, 10);

-- 6. car_makes (10)
INSERT INTO car_makes (name, origin_country, status) VALUES
('Toyota', 'Japan', 'active'),
('Honda', 'Japan', 'active'),
('Nissan', 'Japan', 'active'),
('Mazda', 'Japan', 'active'),
('Suzuki', 'Japan', 'active'),
('Mitsubishi', 'Japan', 'active'),
('BMW', 'Germany', 'active'),
('Mercedes-Benz', 'Germany', 'active'),
('Audi', 'Germany', 'active'),
('Volkswagen', 'Germany', 'active');

-- 7. car_models (10): make_id 1-10
INSERT INTO car_models (make_id, name, status) VALUES
(1, 'Camry', 'active'), (2, 'Accord', 'active'), (3, 'X-Trail', 'active'),
(4, 'Mazda3', 'active'), (5, 'Swift', 'active'), (6, 'Pajero', 'active'),
(7, '3 Series', 'active'), (8, 'C-Class', 'active'), (9, 'A4', 'active'),
(10, 'Golf', 'active');

-- 8. cars (10): model_id 1-10
INSERT INTO cars (model_id, ref_no, package, body_type, year, color, mileage_km, fuel, transmission, drive, steering, status) VALUES
(1, 'REF-001', 'Standard', 'Sedan', 2020, 'White', 45000, 'Petrol', 'Automatic', 'FWD', 'Right', 'available'),
(2, 'REF-002', 'Sport', 'Sedan', 2021, 'Black', 32000, 'Petrol', 'CVT', 'FWD', 'Right', 'available'),
(3, 'REF-003', 'Premium', 'SUV', 2019, 'Silver', 58000, 'Diesel', 'Automatic', 'AWD', 'Right', 'available'),
(4, 'REF-004', 'Base', 'Sedan', 2022, 'Red', 15000, 'Petrol', 'Manual', 'FWD', 'Right', 'available'),
(5, 'REF-005', 'GLX', 'Hatchback', 2020, 'Blue', 41000, 'Petrol', 'Automatic', 'FWD', 'Right', 'available'),
(6, 'REF-006', 'GLS', 'SUV', 2018, 'Gray', 72000, 'Diesel', 'Automatic', '4WD', 'Right', 'available'),
(7, 'REF-007', 'M Sport', 'Sedan', 2021, 'Black', 28000, 'Petrol', 'DCT', 'RWD', 'Left', 'available'),
(8, 'REF-008', 'AMG Line', 'Sedan', 2020, 'White', 35000, 'Petrol', 'Automatic', 'RWD', 'Left', 'available'),
(9, 'REF-009', 'S Line', 'Sedan', 2022, 'Gray', 12000, 'Petrol', 'DCT', 'AWD', 'Left', 'available'),
(10, 'REF-010', 'GTI', 'Hatchback', 2021, 'Red', 22000, 'Petrol', 'DCT', 'FWD', 'Left', 'available');

-- 9. car_photos (10): car_id 1-10
INSERT INTO car_photos (car_id, url, is_primary, sort_order) VALUES
(1, 'https://cdn.example.com/cars/1/photo1.jpg', true, 0),
(2, 'https://cdn.example.com/cars/2/photo1.jpg', true, 0),
(3, 'https://cdn.example.com/cars/3/photo1.jpg', true, 0),
(4, 'https://cdn.example.com/cars/4/photo1.jpg', true, 0),
(5, 'https://cdn.example.com/cars/5/photo1.jpg', true, 0),
(6, 'https://cdn.example.com/cars/6/photo1.jpg', true, 0),
(7, 'https://cdn.example.com/cars/7/photo1.jpg', true, 0),
(8, 'https://cdn.example.com/cars/8/photo1.jpg', true, 0),
(9, 'https://cdn.example.com/cars/9/photo1.jpg', true, 0),
(10, 'https://cdn.example.com/cars/10/photo1.jpg', true, 0);

-- 10. documents (10): car_id 1-10, uploaded_by 1
INSERT INTO documents (car_id, document_type, file_name, file_path, file_size, mime_type) VALUES
(1, 'Registration', 'reg1.pdf', '/uploads/docs/reg1.pdf', 102400, 'application/pdf'),
(2, 'Registration', 'reg2.pdf', '/uploads/docs/reg2.pdf', 98500, 'application/pdf'),
(3, 'Registration', 'reg3.pdf', '/uploads/docs/reg3.pdf', 110000, 'application/pdf'),
(4, 'Registration', 'reg4.pdf', '/uploads/docs/reg4.pdf', 88000, 'application/pdf'),
(5, 'Registration', 'reg5.pdf', '/uploads/docs/reg5.pdf', 95000, 'application/pdf'),
(6, 'Registration', 'reg6.pdf', '/uploads/docs/reg6.pdf', 105000, 'application/pdf'),
(7, 'Registration', 'reg7.pdf', '/uploads/docs/reg7.pdf', 92000, 'application/pdf'),
(8, 'Registration', 'reg8.pdf', '/uploads/docs/reg8.pdf', 101000, 'application/pdf'),
(9, 'Registration', 'reg9.pdf', '/uploads/docs/reg9.pdf', 87000, 'application/pdf'),
(10, 'Registration', 'reg10.pdf', '/uploads/docs/reg10.pdf', 99000, 'application/pdf');

-- 11. car_grades (10): car_id 1-10
INSERT INTO car_grades (car_id, grade_overall, grade_exterior, grade_interior) VALUES
(1, '4', 'A', 'B+'), (2, '4.5', 'A', 'A'), (3, '3.5', 'B+', 'B'),
(4, '5', 'A', 'A'), (5, '4', 'B+', 'B+'), (6, '3', 'B', 'B'),
(7, '4.5', 'A', 'A-'), (8, '4', 'A-', 'B+'), (9, '5', 'A', 'A'),
(10, '4.5', 'A', 'A');

-- 12. car_details (10): car_id 1-10
INSERT INTO car_details (car_id, short_title, full_title, description) VALUES
(1, 'Toyota Camry 2020', 'Toyota Camry 2020 White Sedan', 'Well maintained sedan.'),
(2, 'Honda Accord 2021', 'Honda Accord 2021 Black Sedan', 'Low mileage, single owner.'),
(3, 'Nissan X-Trail 2019', 'Nissan X-Trail 2019 Silver SUV', 'Full option SUV.'),
(4, 'Mazda3 2022', 'Mazda3 2022 Red Sedan', 'Like new condition.'),
(5, 'Suzuki Swift 2020', 'Suzuki Swift 2020 Blue Hatchback', 'Economical hatchback.'),
(6, 'Mitsubishi Pajero 2018', 'Mitsubishi Pajero 2018 Gray SUV', 'Strong off-road SUV.'),
(7, 'BMW 3 Series 2021', 'BMW 3 Series 2021 Black Sedan', 'Luxury sports sedan.'),
(8, 'Mercedes C-Class 2020', 'Mercedes C-Class 2020 White Sedan', 'Premium sedan.'),
(9, 'Audi A4 2022', 'Audi A4 2022 Gray Sedan', 'Top condition.'),
(10, 'VW Golf GTI 2021', 'VW Golf GTI 2021 Red Hatchback', 'Sport hatchback.');

-- 13. car_sub_details (10): car_detail_id 1-10
INSERT INTO car_sub_details (car_detail_id, title, description) VALUES
(1, 'Engine', '2.5L 4-cylinder'), (2, 'Engine', '1.5L Turbo'), (3, 'Engine', '2.0L Diesel'),
(4, 'Engine', '2.0L SkyActiv'), (5, 'Engine', '1.2L'), (6, 'Engine', '3.2L V6'),
(7, 'Engine', '2.0L TwinPower'), (8, 'Engine', '2.0L Turbo'), (9, 'Engine', '2.0L TFSI'),
(10, 'Engine', '2.0L TSI');

-- 14. stocks (10): car_id 1-10
INSERT INTO stocks (car_id, quantity, notes) VALUES
(1, 1, 'In showroom'), (2, 1, 'In showroom'), (3, 1, 'In showroom'),
(4, 1, 'In showroom'), (5, 1, 'In showroom'), (6, 1, 'In showroom'),
(7, 1, 'In showroom'), (8, 1, 'In showroom'), (9, 1, 'In showroom'),
(10, 1, 'In showroom');

-- 15. purchase_history (10): car_id 1-10 — matches live DB columns
INSERT INTO purchase_history (car_id, purchase_date, purchase_amount, govt_duty, cnf_amount, lc_date, lc_number, lc_bank_name, total_units_per_lc, foreign_amount, bdt_amount) VALUES
(1, '2024-01-20', 25000.00, 500000.00, 2762500.00, '2024-01-15', 'LC-001', 'Bank A', 5, 25000.00, 2762500.00),
(2, '2024-01-25', 22000.00, 450000.00, 2431000.00, '2024-01-20', 'LC-002', 'Bank B', 3, 22000.00, 2431000.00),
(3, '2024-02-05', 28000.00, 550000.00, 3108000.00, '2024-02-01', 'LC-003', 'Bank A', 4, 28000.00, 3108000.00),
(4, '2024-02-12', 18000.00, 350000.00, 1998000.00, '2024-02-10', 'LC-004', 'Bank C', 2, 18000.00, 1998000.00),
(5, '2024-02-20', 12000.00, 250000.00, 1332000.00, '2024-02-15', 'LC-005', 'Bank B', 6, 12000.00, 1332000.00),
(6, '2024-03-05', 32000.00, 600000.00, 3584000.00, '2024-03-01', 'LC-006', 'Bank A', 3, 32000.00, 3584000.00),
(7, '2024-03-08', 45000.00, 800000.00, 5040000.00, '2024-03-05', 'LC-007', 'Bank C', 4, 45000.00, 5040000.00),
(8, '2024-03-12', 42000.00, 750000.00, 4704000.00, '2024-03-10', 'LC-008', 'Bank B', 2, 42000.00, 4704000.00),
(9, '2024-03-14', 38000.00, 700000.00, 4256000.00, '2024-03-12', 'LC-009', 'Bank A', 5, 38000.00, 4256000.00),
(10, '2024-03-18', 26000.00, 500000.00, 2912000.00, '2024-03-15', 'LC-010', 'Bank C', 3, 26000.00, 2912000.00);

-- 16. orders (10): user_id 1-10
INSERT INTO orders (user_id, total_amount, shipping_address, status) VALUES
(1, 2800000.00, 'Dhaka, Gulshan', 'delivered'),
(2, 2500000.00, 'Dhaka, Dhanmondi', 'delivered'),
(3, 3200000.00, 'Chittagong', 'shipped'),
(4, 2050000.00, 'Dhaka, Banani', 'pending'),
(5, 1380000.00, 'Sylhet', 'approved'),
(6, 3650000.00, 'Dhaka, Baridhara', 'pending'),
(7, 5150000.00, 'Dhaka, Gulshan-2', 'approved'),
(8, 4800000.00, 'Dhaka, Uttara', 'pending'),
(9, 4350000.00, 'Dhaka, Motijheel', 'delivered'),
(10, 2980000.00, 'Rajshahi', 'shipped');

-- 17. order_items (10): order_id 1-10, car_id 1-10
INSERT INTO order_items (order_id, car_id, quantity, price, notes) VALUES
(1, 1, 1, 2800000.00, 'Toyota Camry'),
(2, 2, 1, 2500000.00, 'Honda Accord'),
(3, 3, 1, 3200000.00, 'Nissan X-Trail'),
(4, 4, 1, 2050000.00, 'Mazda3'),
(5, 5, 1, 1380000.00, 'Suzuki Swift'),
(6, 6, 1, 3650000.00, 'Mitsubishi Pajero'),
(7, 7, 1, 5150000.00, 'BMW 3 Series'),
(8, 8, 1, 4800000.00, 'Mercedes C-Class'),
(9, 9, 1, 4350000.00, 'Audi A4'),
(10, 10, 1, 2980000.00, 'VW Golf GTI');

-- 18. payment_history (10): car_id 1-10
INSERT INTO payment_history (car_id, showroom_name, purchase_amount, purchase_date, customer_name, customer_address, contact_number) VALUES
(1, 'Showroom A', 2800000.00, '2024-02-01', 'Customer One', 'Dhaka', '01710000001'),
(2, 'Showroom A', 2500000.00, '2024-02-05', 'Customer Two', 'Dhaka', '01710000002'),
(3, 'Showroom B', 3200000.00, '2024-02-10', 'Customer Three', 'Chittagong', '01710000003'),
(4, 'Showroom A', 2050000.00, '2024-02-15', 'Customer Four', 'Dhaka', '01710000004'),
(5, 'Showroom C', 1380000.00, '2024-02-20', 'Customer Five', 'Sylhet', '01710000005'),
(6, 'Showroom B', 3650000.00, '2024-03-01', 'Customer Six', 'Dhaka', '01710000006'),
(7, 'Showroom A', 5150000.00, '2024-03-05', 'Customer Seven', 'Dhaka', '01710000007'),
(8, 'Showroom B', 4800000.00, '2024-03-08', 'Customer Eight', 'Dhaka', '01710000008'),
(9, 'Showroom A', 4350000.00, '2024-03-10', 'Customer Nine', 'Dhaka', '01710000009'),
(10, 'Showroom C', 2980000.00, '2024-03-12', 'Customer Ten', 'Rajshahi', '01710000010');

-- 19. installments (10): payment_history_id 1-10
INSERT INTO installments (payment_history_id, installment_date, description, amount, payment_method, balance, remarks) VALUES
(1, '2024-02-01', 'Down payment', 1400000.00, 'Bank', 1400000.00, 'First installment'),
(2, '2024-02-05', 'Down payment', 1250000.00, 'Cash', 1250000.00, 'First installment'),
(3, '2024-02-10', 'Down payment', 1600000.00, 'Bank', 1600000.00, 'First installment'),
(4, '2024-02-15', 'Down payment', 1025000.00, 'Bank', 1025000.00, 'First installment'),
(5, '2024-02-20', 'Full payment', 1380000.00, 'Cash', 0.00, 'Full'),
(6, '2024-03-01', 'Down payment', 1825000.00, 'Bank', 1825000.00, 'First installment'),
(7, '2024-03-05', 'Down payment', 2575000.00, 'Bank', 2575000.00, 'First installment'),
(8, '2024-03-08', 'Down payment', 2400000.00, 'Bank', 2400000.00, 'First installment'),
(9, '2024-03-10', 'Down payment', 2175000.00, 'Bank', 2175000.00, 'First installment'),
(10, '2024-03-12', 'Down payment', 1490000.00, 'Cash', 1490000.00, 'First installment');

-- 20. carts (10): user_id 1-10, car_id 1-10
INSERT INTO carts (user_id, car_id, quantity) VALUES
(1, 1, 1), (2, 2, 1), (3, 3, 1), (4, 4, 1), (5, 5, 1),
(6, 6, 1), (7, 7, 1), (8, 8, 1), (9, 9, 1), (10, 10, 1);

-- 21. rag_chunks (10)
INSERT INTO rag_chunks (source_type, source_id, content, metadata) VALUES
('car', '1', 'Toyota Camry 2020 White Sedan. Well maintained. 45,000 km.', '{"make":"Toyota","model":"Camry"}'),
('car', '2', 'Honda Accord 2021 Black Sedan. Low mileage, single owner.', '{"make":"Honda","model":"Accord"}'),
('car', '3', 'Nissan X-Trail 2019 Silver SUV. Full option.', '{"make":"Nissan","model":"X-Trail"}'),
('car', '4', 'Mazda3 2022 Red Sedan. Like new condition.', '{"make":"Mazda","model":"Mazda3"}'),
('car', '5', 'Suzuki Swift 2020 Blue Hatchback. Economical.', '{"make":"Suzuki","model":"Swift"}'),
('car', '6', 'Mitsubishi Pajero 2018 Gray SUV. Strong off-road.', '{"make":"Mitsubishi","model":"Pajero"}'),
('car', '7', 'BMW 3 Series 2021 Black Sedan. Luxury sports.', '{"make":"BMW","model":"3 Series"}'),
('car', '8', 'Mercedes C-Class 2020 White Sedan. Premium.', '{"make":"Mercedes-Benz","model":"C-Class"}'),
('car', '9', 'Audi A4 2022 Gray Sedan. Top condition.', '{"make":"Audi","model":"A4"}'),
('car', '10', 'VW Golf GTI 2021 Red Hatchback. Sport.', '{"make":"Volkswagen","model":"Golf"}');
