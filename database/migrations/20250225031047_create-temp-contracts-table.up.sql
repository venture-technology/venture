CREATE TABLE temp_contracts (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    signing_url TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    expired_at BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL,
    driver_cnh VARCHAR(20) NOT NULL,
    responsible_cpf VARCHAR(14) NOT NULL,
    school_cnpj VARCHAR(18) NOT NULL,
    kid_rg VARCHAR(12) NOT NULL,
    driver_assigned_at BIGINT,
    responsible_assigned_at BIGINT
);
