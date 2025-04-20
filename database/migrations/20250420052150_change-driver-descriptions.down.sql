BEGIN;

ALTER TABLE drivers RENAME COLUMN description TO descriptions;

COMMIT;
