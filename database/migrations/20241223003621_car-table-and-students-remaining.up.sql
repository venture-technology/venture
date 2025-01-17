-- Criação da tabela cars
CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    car_name VARCHAR(255) NOT NULL,
    car_year INTEGER NOT NULL,
    car_limit INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (car_name, car_year)
);

-- Alterações na tabela drivers
ALTER TABLE drivers
ADD CONSTRAINT unique_cnh UNIQUE (cnh);

ALTER TABLE drivers DROP COLUMN bank_name CASCADE;
ALTER TABLE drivers DROP COLUMN agency_number CASCADE;
ALTER TABLE drivers DROP COLUMN account_number CASCADE;
ALTER TABLE drivers DROP COLUMN car_model CASCADE;
ALTER TABLE drivers DROP COLUMN car_year CASCADE;

ALTER TABLE drivers ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE drivers ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE drivers ADD COLUMN students_remaining INTEGER NOT NULL DEFAULT 0;

-- Alterações na tabela children
ALTER TABLE children
ADD CONSTRAINT unique_rg UNIQUE (rg);

ALTER TABLE children ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE children ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Alterações na tabela responsible
ALTER TABLE responsible
ADD CONSTRAINT unique_cpf UNIQUE (cpf);

ALTER TABLE responsible DROP COLUMN status CASCADE;
ALTER TABLE responsible ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE responsible ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Alterações na tabela schools
ALTER TABLE schools
ADD CONSTRAINT unique_cnpj UNIQUE (cnpj);

ALTER TABLE schools ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE schools ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Alterações na tabela contracts
ALTER TABLE contracts
ADD CONSTRAINT unique_contract UNIQUE (responsible_id, child_id, driver_id, school_id);

ALTER TABLE contracts ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Alterações na tabela partners
ALTER TABLE partners
ADD CONSTRAINT unique_partner UNIQUE (driver_id, school_id);

ALTER TABLE partners ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Alterações na tabela invites
ALTER TABLE invites
ADD CONSTRAINT unique_invite UNIQUE (requester, guester);

ALTER TABLE invites ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE invites ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
