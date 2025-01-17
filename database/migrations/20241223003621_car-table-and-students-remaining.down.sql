-- Remover a tabela cars
DROP TABLE IF EXISTS cars;

-- Alterações na tabela drivers
ALTER TABLE drivers DROP CONSTRAINT IF EXISTS unique_cnh;

ALTER TABLE drivers DROP COLUMN IF EXISTS created_at;
ALTER TABLE drivers DROP COLUMN IF EXISTS updated_at;
ALTER TABLE drivers DROP COLUMN IF EXISTS students_remaining;

ALTER TABLE drivers ADD COLUMN bank_name VARCHAR(255);
ALTER TABLE drivers ADD COLUMN agency_number VARCHAR(50);
ALTER TABLE drivers ADD COLUMN account_number VARCHAR(50);
ALTER TABLE drivers ADD COLUMN car_model VARCHAR(255);
ALTER TABLE drivers ADD COLUMN car_year INTEGER;

-- Alterações na tabela children
ALTER TABLE children DROP CONSTRAINT IF EXISTS unique_rg;

ALTER TABLE children DROP COLUMN IF EXISTS created_at;
ALTER TABLE children DROP COLUMN IF EXISTS updated_at;

-- Alterações na tabela responsible
ALTER TABLE responsible DROP CONSTRAINT IF EXISTS unique_cpf;

ALTER TABLE responsible DROP COLUMN IF EXISTS created_at;
ALTER TABLE responsible DROP COLUMN IF EXISTS updated_at;
ALTER TABLE responsible ADD COLUMN status VARCHAR(50);

-- Alterações na tabela schools
ALTER TABLE schools DROP CONSTRAINT IF EXISTS unique_cnpj;

ALTER TABLE schools DROP COLUMN IF EXISTS created_at;
ALTER TABLE schools DROP COLUMN IF EXISTS updated_at;

-- Alterações na tabela contracts
ALTER TABLE contracts DROP CONSTRAINT IF EXISTS unique_contract;

ALTER TABLE contracts DROP COLUMN IF EXISTS created_at;
ALTER TABLE contracts DROP COLUMN IF EXISTS updated_at;

-- Alterações na tabela partners
ALTER TABLE partners DROP CONSTRAINT IF EXISTS unique_partner;

ALTER TABLE partners DROP COLUMN IF EXISTS updated_at;

-- Alterações na tabela invites
ALTER TABLE invites DROP CONSTRAINT IF EXISTS unique_invite;

ALTER TABLE invites DROP COLUMN IF EXISTS created_at;
ALTER TABLE invites DROP COLUMN IF EXISTS updated_at;
