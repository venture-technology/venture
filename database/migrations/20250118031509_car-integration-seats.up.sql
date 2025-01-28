ALTER TABLE drivers ADD COLUMN car_name VARCHAR(225);
ALTER TABLE drivers ADD COLUMN car_year VARCHAR(4);
ALTER TABLE drivers ADD COLUMN car_capacity VARCHAR(3);
ALTER TABLE drivers ADD COLUMN schedule VARCHAR(3);
ALTER TABLE drivers ADD COLUMN seats_remaining VARCHAR(3);
ALTER TABLE drivers ADD COLUMN seats_morning VARCHAR(3);
ALTER TABLE drivers ADD COLUMN seats_afternoon VARCHAR(3);
ALTER TABLE drivers ADD COLUMN seats_night VARCHAR(3);

ALTER TABLE invites ADD COLUMN id_temp UUID;
UPDATE invites SET id_temp = id;
ALTER TABLE invites DROP COLUMN id;
ALTER TABLE invites ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE invites DROP COLUMN id_temp;

ALTER TABLE contracts DROP CONSTRAINT contracts_pkey;
ALTER TABLE contracts RENAME COLUMN record TO uuid;
ALTER TABLE contracts ADD COLUMN id SERIAL PRIMARY KEY;

ALTER TABLE invites ADD COLUMN id_temp UUID;
UPDATE invites SET id_temp = id;
ALTER TABLE invites DROP COLUMN id;
ALTER TABLE invites ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE invites DROP COLUMN id_temp;

ALTER TABLE contracts DROP CONSTRAINT contracts_pkey;
ALTER TABLE contracts RENAME COLUMN record TO uuid;
ALTER TABLE contracts ADD COLUMN id SERIAL PRIMARY KEY;

ALTER TABLE responsible RENAME TO responsibles;

ALTER TABLE children rename to kids;

DROP TABLE kids cascade;

CREATE TABLE IF NOT EXISTS kids (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    rg VARCHAR(9) PRIMARY KEY NOT NULL,
    responsible_id VARCHAR(11) NOT NULL,
    shift TEXT NOT NULL,
    profile_image text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (responsible_id) REFERENCES responsibles(cpf) ON DELETE cascade,
    CONSTRAINT unique_kid UNIQUE (rg)
);

alter table drivers rename column qrcode to qr_code;
alter table drivers drop column students_remaining