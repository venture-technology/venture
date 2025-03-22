ALTER TABLE contracts
ADD CONSTRAINT fk_responsible FOREIGN KEY (responsible_cpf) REFERENCES responsibles (cpf) ON DELETE CASCADE,
ADD CONSTRAINT fk_driver FOREIGN KEY (driver_cnh) REFERENCES drivers (cnh) ON DELETE CASCADE,
ADD CONSTRAINT fk_school FOREIGN KEY (school_cnpj) REFERENCES schools (cnpj) ON DELETE CASCADE,
ADD CONSTRAINT fk_kid FOREIGN KEY (kid_rg) REFERENCES kids (rg) ON DELETE CASCADE;
