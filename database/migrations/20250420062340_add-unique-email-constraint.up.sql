BEGIN;

ALTER TABLE responsibles ADD CONSTRAINT unique_responsibles_email UNIQUE (email);
ALTER TABLE schools ADD CONSTRAINT unique_schools_email UNIQUE (email);
ALTER TABLE drivers ADD CONSTRAINT unique_drivers_email UNIQUE (email);

COMMIT;
