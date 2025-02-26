-- Alterações na tabela drivers
ALTER TABLE drivers ADD COLUMN states VARCHAR(100) DEFAULT '';
ALTER TABLE drivers ADD COLUMN city VARCHAR (100) DEFAULT '';

-- Alterações na tabela school 
ALTER TABLE schools ADD COLUMN states VARCHAR(100) DEFAULT '';
ALTER TABLE Schools ADD COLUMN city VARCHAR (100) DEFAULT '';

-- Alterações na tabela responsible
ALTER TABLE responsibles ADD COLUMN states VARCHAR(100) DEFAULT '';
ALTER TABLE responsibles ADD COLUMN city VARCHAR (100) DEFAULT '';