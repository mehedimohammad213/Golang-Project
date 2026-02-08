-- Postgres Schema

-- Users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    username VARCHAR(80) NOT NULL UNIQUE,
    email VARCHAR(150) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    last_login_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    deleted_at TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_last_login ON users(last_login_at);
CREATE INDEX idx_deleted_at ON users(deleted_at);

-- Roles
CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    slug VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Permissions
CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    slug VARCHAR(150) NOT NULL UNIQUE,
    module VARCHAR(80),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Role ↔ User
CREATE TABLE role_user (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    assigned_by BIGINT,
    UNIQUE (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (assigned_by) REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_role_user_user ON role_user(user_id);
CREATE INDEX idx_role_user_role ON role_user(role_id);

-- Permission ↔ Role
CREATE TABLE permission_role (
    id BIGSERIAL PRIMARY KEY,
    permission_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (permission_id, role_id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);
CREATE INDEX idx_permission_role_perm ON permission_role(permission_id);
CREATE INDEX idx_permission_role_role ON permission_role(role_id);

-- Car Makes
CREATE TABLE car_makes (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    origin_country VARCHAR(80),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active','inactive')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_car_makes_status ON car_makes(status);

-- Car Models
CREATE TABLE car_models (
    id BIGSERIAL PRIMARY KEY,
    make_id BIGINT NOT NULL,
    name VARCHAR(150) NOT NULL,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active','inactive')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (make_id, name),
    FOREIGN KEY (make_id) REFERENCES car_makes(id) ON DELETE RESTRICT
);
CREATE INDEX idx_car_models_make ON car_models(make_id);
CREATE INDEX idx_car_models_status ON car_models(status);

-- Cars
CREATE TYPE body_type_enum AS ENUM (
    'Sedan','Hatchback','SUV','Crossover','Coupe','Convertible',
    'Wagon','Van','Minivan','Pickup','Microbus','Roadster',
    'Fastback','Liftback'
);
CREATE TYPE fuel_enum AS ENUM ('Petrol','Diesel','Hybrid','Electric','CNG','LPG');
CREATE TYPE transmission_enum AS ENUM ('Manual','Automatic','CVT','DCT');
CREATE TYPE drive_enum AS ENUM ('FWD','RWD','AWD','4WD');
CREATE TYPE steering_enum AS ENUM ('Left','Right');
CREATE TYPE car_status_enum AS ENUM ('available','sold','reserved','damaged','lost','stolen');

CREATE TABLE cars (
    id BIGSERIAL PRIMARY KEY,
    model_id BIGINT NOT NULL,
    ref_no VARCHAR(32) UNIQUE,
    package VARCHAR(255),
    body_type body_type_enum,
    year SMALLINT,
    color VARCHAR(64),
    reg_year_month VARCHAR(10),
    mileage_km INT,
    chassis_no_full VARCHAR(64) UNIQUE,
    engine_cc INT,
    fuel fuel_enum,
    transmission transmission_enum,
    drive drive_enum,
    engine_number VARCHAR(64),
    seats SMALLINT,
    number_of_keys INT,
    keys_feature VARCHAR(255),
    steering steering_enum,
    location VARCHAR(128),
    country_origin VARCHAR(64),
    status car_status_enum DEFAULT 'available',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (model_id) REFERENCES car_models(id) ON DELETE RESTRICT
);
CREATE INDEX idx_cars_model ON cars(model_id);
CREATE INDEX idx_cars_status ON cars(status);

-- Car Photos
CREATE TABLE car_photos (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL,
    url VARCHAR(512) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    sort_order INT DEFAULT 0,
    is_hidden BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);
CREATE INDEX idx_car_photos_car ON car_photos(car_id);

-- Documents
CREATE TABLE documents (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL,
    document_type VARCHAR(100) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(512) NOT NULL,
    file_size BIGINT,
    mime_type VARCHAR(100),
    is_primary BOOLEAN DEFAULT FALSE,
    is_hidden BOOLEAN DEFAULT FALSE,
    sort_order INT DEFAULT 0,
    uploaded_by BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE,
    FOREIGN KEY (uploaded_by) REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_documents_car ON documents(car_id);

-- Car Grades
CREATE TYPE overall_grade_enum AS ENUM ('S','6','5','4.5','4','3.5','3','2','1','RA','R','0');
CREATE TYPE detail_grade_enum AS ENUM ('A','A-','B+','B','B-','C+','C','C-','D','E');

CREATE TABLE car_grades (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL UNIQUE,
    grade_overall overall_grade_enum,
    grade_exterior detail_grade_enum,
    grade_interior detail_grade_enum,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);

-- Car Details
CREATE TABLE car_details (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL UNIQUE,
    short_title VARCHAR(255),
    full_title VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);

-- Car Sub Details
CREATE TABLE car_sub_details (
    id BIGSERIAL PRIMARY KEY,
    car_detail_id BIGINT NOT NULL,
    title VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_detail_id) REFERENCES car_details(id) ON DELETE CASCADE
);
CREATE INDEX idx_car_sub_details_detail ON car_sub_details(car_detail_id);

-- Stocks
CREATE TABLE stocks (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL UNIQUE,
    quantity INT NOT NULL CHECK (quantity >= 0),
    notes VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);

-- Carts
CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    car_id BIGINT NOT NULL,
    quantity INT NOT NULL DEFAULT 1 CHECK (quantity > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, car_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);

-- Orders
CREATE TYPE order_status_enum AS ENUM ('pending','approved','shipped','delivered','canceled');

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    total_amount DECIMAL(15,2) NOT NULL CHECK (total_amount >= 0),
    shipping_address VARCHAR(512),
    status order_status_enum DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX idx_orders_user ON orders(user_id);
CREATE INDEX idx_orders_status ON orders(status);

-- Order Items
CREATE TABLE order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    car_id BIGINT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price DECIMAL(15,2) NOT NULL CHECK (price >= 0),
    notes VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (order_id, car_id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE RESTRICT
);
CREATE INDEX idx_order_items_order ON order_items(order_id);
CREATE INDEX idx_order_items_car ON order_items(car_id);

-- Purchase History
CREATE TABLE purchase_history (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL,
    purchase_date DATE,
    purchase_amount DECIMAL(15,2) CHECK (purchase_amount >= 0),
    govt_duty DECIMAL(15,2) CHECK (govt_duty >= 0),
    cnf_amount DECIMAL(15,2) CHECK (cnf_amount >= 0),
    lc_date DATE,
    lc_number VARCHAR(64),
    lc_bank_name VARCHAR(128),
    lc_bank_branch_name VARCHAR(128),
    lc_bank_branch_address VARCHAR(256),
    total_units_per_lc INT CHECK (total_units_per_lc >= 0),
    foreign_amount DECIMAL(15,2) CHECK (foreign_amount >= 0),
    bdt_amount DECIMAL(15,2) CHECK (bdt_amount >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);
CREATE INDEX idx_purchase_history_car ON purchase_history(car_id);

-- Payment History
CREATE TABLE payment_history (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT,
    showroom_name VARCHAR(255),
    wholesaler_address VARCHAR(255),
    purchase_amount DECIMAL(15,2) CHECK (purchase_amount >= 0),
    purchase_date DATE,
    customer_name VARCHAR(255),
    nid_number VARCHAR(50),
    tin_certificate VARCHAR(100),
    customer_address VARCHAR(512),
    contact_number VARCHAR(20),
    email VARCHAR(150),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE SET NULL
);
CREATE INDEX idx_payment_history_car ON payment_history(car_id);

-- Installments
CREATE TYPE payment_method_enum AS ENUM ('Bank','Cash');

CREATE TABLE installments (
    id BIGSERIAL PRIMARY KEY,
    payment_history_id BIGINT NOT NULL,
    installment_date DATE,
    description VARCHAR(255),
    amount DECIMAL(15,2) CHECK (amount >= 0),
    payment_method payment_method_enum,
    bank_name VARCHAR(128),
    cheque_number VARCHAR(64),
    balance DECIMAL(15,2) CHECK (balance >= 0),
    remarks VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_history_id) REFERENCES payment_history(id) ON DELETE CASCADE
);
CREATE INDEX idx_installments_payment ON installments(payment_history_id);
