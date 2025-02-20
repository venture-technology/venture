ALTER TABLE drivers 
ALTER COLUMN seats_remaining 
SET DATA TYPE VARCHAR USING seats_remaining::VARCHAR;

ALTER TABLE drivers 
ALTER COLUMN seats_morning 
SET DATA TYPE VARCHAR USING seats_morning::VARCHAR;

ALTER TABLE drivers 
ALTER COLUMN seats_afternoon 
SET DATA TYPE VARCHAR USING seats_afternoon::VARCHAR;

ALTER TABLE drivers 
ALTER COLUMN seats_night 
SET DATA TYPE VARCHAR USING seats_night::VARCHAR;
