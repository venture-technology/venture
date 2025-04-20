BEGIN;

ALTER TABLE responsibles DROP CONSTRAINT unique_responsibles_email;
ALTER TABLE schools DROP CONSTRAINT unique_schools_email;
ALTER TABLE drivers DROP CONSTRAINT unique_drivers_email;

COMMIT;
