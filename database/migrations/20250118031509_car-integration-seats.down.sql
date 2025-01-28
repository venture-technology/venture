-- Remover as colunas adicionadas na tabela drivers
ALTER TABLE drivers DROP COLUMN car_name;
ALTER TABLE drivers DROP COLUMN car_year;
ALTER TABLE drivers DROP COLUMN car_capacity;
ALTER TABLE drivers DROP COLUMN schedule;
ALTER TABLE drivers DROP COLUMN seats_remaining;
ALTER TABLE drivers DROP COLUMN seats_morning;
ALTER TABLE drivers DROP COLUMN seats_afternoon;
ALTER TABLE drivers DROP COLUMN seats_night;
