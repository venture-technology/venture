ALTER TABLE drivers 
ALTER COLUMN seats_remaining 
SET DATA TYPE INTEGER USING seats_remaining::INTEGER;

ALTER TABLE drivers 
ALTER COLUMN seats_morning 
SET DATA TYPE INTEGER USING seats_morning::INTEGER;

ALTER TABLE drivers 
ALTER COLUMN seats_afternoon 
SET DATA TYPE INTEGER USING seats_afternoon::INTEGER;

ALTER TABLE drivers 
ALTER COLUMN seats_night 
SET DATA TYPE INTEGER USING seats_night::INTEGER;
