BEGIN;

ALTER TABLE responsibles
  ADD COLUMN card_token TEXT,
  ADD COLUMN payment_method_id TEXT;
  ADD COLUMN customer_id TEXT;

COMMIT;
