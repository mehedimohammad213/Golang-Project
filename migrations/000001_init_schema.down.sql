-- Drop tables in reverse dependency order (children first)
DROP TABLE IF EXISTS installments;
DROP TABLE IF EXISTS payment_history;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS lc_purchase_documents;
DROP TABLE IF EXISTS purchase_history;
DROP TABLE IF EXISTS lc_cars;
DROP TABLE IF EXISTS lcs;
DROP TABLE IF EXISTS currency_rates;
DROP TABLE IF EXISTS stocks;
DROP TABLE IF EXISTS car_sub_details;
DROP TABLE IF EXISTS car_details;
DROP TABLE IF EXISTS car_grades;
DROP TABLE IF EXISTS documents;
DROP TABLE IF EXISTS car_photos;
DROP TABLE IF EXISTS cars;
DROP TABLE IF EXISTS car_models;
DROP TABLE IF EXISTS car_makes;
DROP TABLE IF EXISTS permission_role;
DROP TABLE IF EXISTS role_user;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;

-- Drop enum types
DROP TYPE IF EXISTS payment_method_enum;
DROP TYPE IF EXISTS order_status_enum;
DROP TYPE IF EXISTS purchase_currency_enum;
DROP TYPE IF EXISTS car_status_enum;
DROP TYPE IF EXISTS steering_enum;
DROP TYPE IF EXISTS drive_enum;
DROP TYPE IF EXISTS transmission_enum;
DROP TYPE IF EXISTS fuel_enum;
DROP TYPE IF EXISTS body_type_enum;
DROP TYPE IF EXISTS detail_grade_enum;
DROP TYPE IF EXISTS overall_grade_enum;
