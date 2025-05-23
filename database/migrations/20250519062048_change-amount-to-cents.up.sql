BEGIN;

-- Multiplicar os valores existentes por 100
UPDATE contracts
SET amount = amount * 100,
    anual_amount = anual_amount * 100;

-- Alterar o tipo das colunas para integer
ALTER TABLE contracts
    ALTER COLUMN amount TYPE integer USING amount::integer,
    ALTER COLUMN anual_amount TYPE integer USING anual_amount::integer;

-- Multiplicar os valores existentes por 100
UPDATE drivers
SET amount = amount * 100;

-- Alterar o tipo da coluna para integer
ALTER TABLE drivers
    ALTER COLUMN amount TYPE integer USING amount::integer;

COMMIT;
