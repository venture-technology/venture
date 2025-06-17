BEGIN;

ALTER TABLE responsibles
  DROP COLUMN IF EXISTS card_token,
  DROP COLUMN IF EXISTS payment_method_id,
  DROP COLUMN IF EXISTS customer_id;

COMMIT;
