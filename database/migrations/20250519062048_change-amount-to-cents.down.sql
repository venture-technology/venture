BEGIN;

-- Alterar o tipo das colunas para numeric(10, 2)
ALTER TABLE contracts
    ALTER COLUMN amount TYPE numeric(10,2) USING (amount::numeric / 100),
    ALTER COLUMN anual_amount TYPE numeric(10,2) USING (anual_amount::numeric / 100);

-- Alterar o tipo da coluna para numeric(10, 2)
ALTER TABLE drivers
    ALTER COLUMN amount TYPE numeric(10,2) USING (amount::numeric / 100);

COMMIT;
