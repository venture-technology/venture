BEGIN;

-- Tabela: schools
ALTER TABLE schools DROP COLUMN neighborhood;
ALTER TABLE schools ALTER COLUMN city DROP NOT NULL;
ALTER TABLE schools ALTER COLUMN state DROP NOT NULL;
ALTER TABLE schools RENAME COLUMN state TO states;

-- Tabela: responsibles
ALTER TABLE responsibles DROP COLUMN neighborhood;
ALTER TABLE responsibles ALTER COLUMN city DROP NOT NULL;
ALTER TABLE responsibles ALTER COLUMN state DROP NOT NULL;
ALTER TABLE responsibles RENAME COLUMN state TO states;

-- Tabela: drivers
ALTER TABLE drivers DROP COLUMN neighborhood;
ALTER TABLE drivers ALTER COLUMN city DROP NOT NULL;
ALTER TABLE drivers ALTER COLUMN state DROP NOT NULL;
ALTER TABLE drivers RENAME COLUMN state TO states;

COMMIT;
