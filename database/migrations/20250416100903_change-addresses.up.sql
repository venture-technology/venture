BEGIN;

-- Tabela: schools
ALTER TABLE schools RENAME COLUMN states TO state;
UPDATE schools SET city = 'null' WHERE city IS NULL;
UPDATE schools SET state = 'null' WHERE state IS NULL;
ALTER TABLE schools ALTER COLUMN state SET NOT NULL;
ALTER TABLE schools ALTER COLUMN city SET NOT NULL;
ALTER TABLE schools ADD COLUMN neighborhood VARCHAR(255) NOT NULL DEFAULT '';

-- Tabela: responsibles
ALTER TABLE responsibles RENAME COLUMN states TO state;
UPDATE responsibles SET city = 'null' WHERE city IS NULL;
UPDATE responsibles SET state = 'null' WHERE state IS NULL;
ALTER TABLE responsibles ALTER COLUMN state SET NOT NULL;
ALTER TABLE responsibles ALTER COLUMN city SET NOT NULL;
ALTER TABLE responsibles ADD COLUMN neighborhood VARCHAR(255) NOT NULL DEFAULT '';

-- Tabela: drivers
ALTER TABLE drivers RENAME COLUMN states TO state;
UPDATE drivers SET city = 'null' WHERE city IS NULL;
UPDATE drivers SET state = 'null' WHERE state IS NULL;
ALTER TABLE drivers ALTER COLUMN state SET NOT NULL;
ALTER TABLE drivers ALTER COLUMN city SET NOT NULL;
ALTER TABLE drivers ADD COLUMN neighborhood VARCHAR(255) NOT NULL DEFAULT '';

COMMIT;
