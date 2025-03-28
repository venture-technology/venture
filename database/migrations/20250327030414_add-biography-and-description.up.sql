-- Alterações na tabela drivers
ALTER TABLE drivers ADD COLUMN biography VARCHAR(1000) DEFAULT '';
ALTER TABLE drivers ADD COLUMN descriptions VARCHAR (550) DEFAULT '';