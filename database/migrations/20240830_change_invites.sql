-- Table Responsible
CREATE TABLE IF NOT EXISTS responsible (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(11) PRIMARY KEY NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    number TEXT NOT NULL,
    complement TEXT,
    zip VARCHAR(8) NOT NULL,
    status TEXT NOT NULL,
    card_token TEXT,
    payment_method_id TEXT,
    customer_id TEXT NOT NULL,
    phone TEXT NOT NULL
);

-- Table Children
CREATE TABLE IF NOT EXISTS children (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    rg VARCHAR(9) PRIMARY KEY NOT NULL,
    responsible_id VARCHAR(11) NOT NULL,
    shift TEXT NOT NULL,
    FOREIGN KEY (responsible_id) REFERENCES responsible(cpf) ON DELETE CASCADE
);

-- Table Schools
CREATE TABLE IF NOT EXISTS schools (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) PRIMARY KEY NOT NULL,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    zip VARCHAR(8) NOT NULL,
    email VARCHAR(100) NOT NULL,
    complement VARCHAR(10),
    phone TEXT NOT NULL
);

-- Table Drivers
CREATE TABLE IF NOT EXISTS drivers (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    cnh VARCHAR(20) PRIMARY KEY NOT NULL,
    qrcode VARCHAR(100) NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(10),
    zip VARCHAR(8) NOT NULL,
    phone TEXT NOT NULL,
    bank_name VARCHAR(100),
    agency_number VARCHAR(4),
    account_number VARCHAR(20),
    pix_key VARCHAR(100)
); 

-- Tabela de Convites
CREATE TABLE IF NOT EXISTS invites (
    id UUID PRIMARY KEY,
    requester VARCHAR(14), -- school
    guester VARCHAR(14), -- driver
    status TEXT NOT NULL,
    FOREIGN KEY (requester) REFERENCES schools(cnpj),
    FOREIGN KEY (guester) REFERENCES drivers(cnh)
);

-- Table partners
CREATE TABLE IF NOT EXISTS partners (
    record SERIAL PRIMARY KEY,
    driver_id VARCHAR(20) NOT NULL,
    school_id VARCHAR(14) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (school_Id) REFERENCES schools(cnpj) ON DELETE CASCADE,
    FOREIGN KEY (driver_id) REFERENCES drivers(cnh) ON DELETE CASCADE
);

-- Table contracts
CREATE TABLE IF NOT EXISTS contracts (
    record UUID PRIMARY KEY,
    title_stripe_subscription TEXT NOT NULL,
    description_stripe_subscription TEXT NOT NULL,
    id_stripe_subscription TEXT NOT NULL,
    id_price_subscription TEXT NOT NULL,
    id_product_subscription TEXT NOT NULL,
    school_id VARCHAR(14) NOT NULL,
    driver_id VARCHAR(20) NOT NULL,
    responsible_id VARCHAR(11) NOT NULL,
    child_id VARCHAR(9) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expire_at TIMESTAMP NOT NULL,
    status TEXT NOT NULL,
    FOREIGN KEY (driver_id) REFERENCES drivers(cnh) ON DELETE CASCADE,
    FOREIGN KEY (school_Id) REFERENCES schools(cnpj) ON DELETE CASCADE,
    FOREIGN KEY (responsible_id) REFERENCES responsible(cpf) ON DELETE CASCADE,
    FOREIGN KEY (child_id) REFERENCES children(rg) ON DELETE CASCADE
);

-- Table emails
CREATE TABLE IF NOT EXISTS email_records (
    id SERIAL PRIMARY KEY,
    recipient TEXT,
    subject TEXT, 
    body TEXT
);