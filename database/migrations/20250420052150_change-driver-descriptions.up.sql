BEGIN;

ALTER TABLE drivers RENAME COLUMN descriptions TO description;

COMMIT
