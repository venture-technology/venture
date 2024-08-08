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
    cnpj VARCHAR(14) PRIMARY KEY,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    zip VARCHAR(8) NOT NULL,
    email VARCHAR(100) NOT NULL,
    complement VARCHAR(10)
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
    zip VARCHAR(8) NOT NULL
); 

-- Tabela de Convites
CREATE TABLE IF NOT EXISTS invites (
    invite_id SERIAL PRIMARY KEY,
    requester VARCHAR(14), -- school
    school VARCHAR(100) NOT NULL, --name_school
    email_school VARCHAR(100) NOT NULL,
    guest VARCHAR(14), -- driver
    driver VARCHAR(100) NOT NULL, --name_driver
    email_driver VARCHAR(100) NOT NULL,
    status TEXT NOT NULL
);